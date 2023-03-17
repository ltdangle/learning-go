package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	runCommand(os.Args)
}

func runCommand(params []string) {
	var cli_args []string

	command := params[1]

	if len(params) == 1 {
		fmt.Println("You did not enter the command to run.")
	}
	if len(params) == 2 {
		cli_args = []string{}
	} else {
		cli_args = params[2:]
	}

	fmt.Println("args: " + strings.Join(os.Args, " "))
	fmt.Println("command: " + command)
	fmt.Println("arguments: " + strings.Join(cli_args, ""))

	out, err := exec.Command(command, cli_args...).Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	files := strings.Split(string(out), "\n")
	fmt.Print(strings.Join(files, "\n"))

}
