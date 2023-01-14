# ShellShockCheck

To test the script for Shellshock vulnerability, you will need to have Go programming language installed on your system.
It's a checker, it does not exploit any vulnerability.

Download the script or copy its code to a new file on your system.
Open a terminal window and navigate to the location of the script.
Run the command go build to build the script.
Run the command ./shellshock-check -u <url> to execute the script. Replace <url> with the target URL.
The script will send a request to the target URL with a specific payload that tests for the Shellshock vulnerability.
The script will then return a message indicating whether the target is vulnerable or not.
