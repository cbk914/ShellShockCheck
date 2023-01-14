# ShellShockCheck

# Install

To test the script for Shellshock vulnerability, you will need to have Go programming language installed on your system.

It's a checker, it does not exploit any vulnerability.

- Download the script or copy its code to a new file on your system.
- Open a terminal window and navigate to the location of the script.
- Run the command go build to build the script.
- Run the command ./shellshock-check -u <url> to execute the script. Replace <url> with the target URL.
- The script will send a request to the target URL with a specific payload that tests for the Shellshock vulnerability.
- The script will then return a message indicating whether the target is vulnerable or not.

# CVE-2014-7169
  
Shellshock, also known as Bashdoor, is a family of security vulnerabilities that affect the Bash shell, a command-line interface commonly used on Linux and UNIX-based operating systems. The vulnerability allows an attacker to execute arbitrary commands on a vulnerable system by injecting specially crafted environment variables into the shell. The first and most well-known variant, known as CVE-2014-6271, was discovered in September 2014 and allows for the execution of arbitrary commands by appending a function definition to the end of the function definition of an environment variable. Subsequent variants, such as CVE-2014-7169, have been discovered which exploit different aspects of the vulnerability.

The vulnerability is caused by a flaw in the way that Bash parses function definitions in environment variables. When Bash parses a function definition, it does not properly validate the function name and treats the entire value of the environment variable as the function body. This allows an attacker to include arbitrary commands in the function definition that will be executed when the function is called. The vulnerability can be exploited in a number of ways, including through HTTP headers, CGI scripts, and SSH environment variables.

The Shellshock vulnerability is considered to be highly critical, as it can be exploited remotely and allows for the execution of arbitrary commands with the permissions of the user running the affected software. Many popular web servers, such as Apache, use Bash as the default shell for CGI scripts, making them vulnerable to attack. Additionally, many Internet of Things (IoT) devices use Bash as the default shell, which makes them particularly susceptible to attack.

The Shellshock vulnerability can be mitigated by patching the affected software. This can be done by upgrading to a version of Bash that includes the patch, or by applying the patch manually. Additionally, it is also recommended to limit the attack surface by disabling unnecessary services and features, such as CGI scripts on web servers and SSH access on IoT devices.
