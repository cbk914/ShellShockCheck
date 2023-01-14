# ShellShockCheck

A python script to check CVE-2014-7169.

## go version

1. Install Go on your machine by following the instructions on the official website ([https://golang.org/doc/install](https://golang.org/doc/install))
2. Download the Shellshock-Checker code provided above and save it to your machine
3. Open a terminal window and navigate to the directory where the code is saved
4. Compile the code by running the command "go build"
5. Run the compiled binary by using the command "./[binary_name] -u [target_url]" where [binary_name] is the name of the compiled binary and [target_url] is the URL of the target you want to test for Shellshock vulnerability.
6. The program will check the target and return the result "VULNERABLE: The target is vulnerable to Shellshock." if the target is vulnerable or "SAFE: The target is not vulnerable to Shellshock." if the target is not vulnerable.
7. You can use the -h option to show the help.

## python version


1. Make sure you have Python 3 installed on your system.
2. Download the script and save it to a local directory.
3. Open a terminal/command prompt and navigate to the directory where the script is saved.
4. Run the script using the following command:

   ```
   python shellshock_checker.py -u [TARGET_URL]
   ```
5. Replace [TARGET_URL] with the URL of the target you want to check for the Shellshock vulnerability. The URL must be in valid format (i.e. http or https) and have a valid hostname.
6. The script will send a request to the target URL and check for the presence of the Shellshock vulnerability.
7. If the target is found to be vulnerable, the script will output "VULNERABLE: The target is vulnerable to Shellshock.". If the target is not vulnerable, the script will output "SAFE: The target is not vulnerable to Shellshock.".
