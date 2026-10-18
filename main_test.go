package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// captureStdout runs fn with os.Stdout redirected to a pipe and returns what
// was written. Used to assert on the human-readable verdicts, which is the
// whole output contract of this tool.
func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe() failed: %v", err)
	}

	orig := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = orig }()

	done := make(chan string, 1)
	go func() {
		b, _ := io.ReadAll(r)
		done <- string(b)
	}()

	fn()
	_ = w.Close()

	return <-done
}

// TestPackageCompiles references every identifier the tool relies on. On the
// pre-fix tree this file does not even build: main.go imported "flags" (not a
// std package), never imported "os" despite calling os.Exit, and imported
// "strings" without using it.
func TestPackageCompiles(t *testing.T) {
	// Referencing the functions forces the compiler to resolve them; if any
	// import block were wrong this test would not compile at all.
	var (
		_ func(string) = check_shellshock
		_ func()       = main
	)
}

// TestCheckShellshockDetectsVulnerableResponse asserts the positive verdict
// against a local test server. No external target is contacted.
func TestCheckShellshockDetectsVulnerableResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Vulnerable", "Vulnerable")
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	out := captureStdout(t, func() { check_shellshock(srv.URL) })

	if !strings.Contains(out, "Target is vulnerable to Shellshock.") {
		t.Fatalf("expected vulnerable verdict, got: %q", out)
	}
}

// TestCheckShellshockReportsSafeResponse asserts the negative verdict so a
// future edit cannot turn the tool into a tool that always says vulnerable.
func TestCheckShellshockReportsSafeResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	out := captureStdout(t, func() { check_shellshock(srv.URL) })

	if !strings.Contains(out, "Target is not vulnerable to Shellshock.") {
		t.Fatalf("expected safe verdict, got: %q", out)
	}
}