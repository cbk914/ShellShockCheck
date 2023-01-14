# ShellShockCheck

A python script to check CVE-2014-7169.

1. Install Go on your machine by following the instructions on the official website ([https://golang.org/doc/install](https://golang.org/doc/install))
2. Download the Shellshock-Checker code provided above and save it to your machine
3. Open a terminal window and navigate to the directory where the code is saved
4. Compile the code by running the command "go build"
5. Run the compiled binary by using the command "./[binary_name] -u [target_url]" where [binary_name] is the name of the compiled binary and [target_url] is the URL of the target you want to test for Shellshock vulnerability.
6. The program will check the target and return the result "VULNERABLE: The target is vulnerable to Shellshock." if the target is vulnerable or "SAFE: The target is not vulnerable to Shellshock." if the target is not vulnerable.
7. You can use the -h option to show the help.
