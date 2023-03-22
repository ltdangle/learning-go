package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println(os.Args[1:])
	return
	command := strings.Join(os.Args[1:], " ")
	runCommand(command)
}

func runCommand(string_params string) {
	fmt.Println(string_params)
	fmt.Println("/usr/local/bin/notmuch search --output=files --format=text folder:support@contentcuria.com/INBOX | grep /INBOX")
	return
	var cli_args []string

	params := strings.Split(string_params, " ")
	command := params[0]

	if len(params) == 1 {
		fmt.Println("You did not enter the command to run.")
	}
	if len(params) == 2 {
		cli_args = []string{}
	} else {
		cli_args = params[1:]
	}

	// fmt.Println("args: " + strings.Join(os.Args, " "))
	// fmt.Println("command: " + command)
	// fmt.Println("arguments: " + strings.Join(cli_args, ""))

	out, err := exec.Command(command, cli_args...).Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	files := strings.Split(string(out), "\n")
	fmt.Print(strings.Join(files, "\n"))

}
