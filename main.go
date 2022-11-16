package main

import (
	"fmt"
	json_actions "github.com/samuelowad/cmdexec/pkg/json-actions"
	"log"
	"os/exec"
	"strings"
)

func main() {

	//str := " sudo ls -l && sudo cat main.go"
	//runCommands(str)

	//json_actions.EncodeJson("name", str, false)
	json_actions.DecodeJson()
}

// run commands
func runCommands(commandString string) {
	commands := strings.Split(commandString, " &&")

	for _, command := range commands {
		command = strings.TrimSpace(command)
		fCommand := strings.Fields(command)
		//
		out, err := exec.Command(fCommand[0], fCommand[1:]...).Output()

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))
	}
}
