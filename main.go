package main

import (
	"fmt"
	json_actions "github.com/samuelowad/cmdexec/pkg/json-actions"
	"log"
	"os/exec"
	"strings"
	"time"
)

func main() {
	//cli.CreateCLi()
	//checkSudoTime()
	//
	str := " sudo ls -l && sudo cat main.go"
	//runCommands(str)

	start := time.Now()
	err := json_actions.EncodeJson("name", str)
	if err != nil {
		fmt.Println(err)
	}
	elapsed := time.Since(start)
	fmt.Println("EncodeJson() execution time:", elapsed)

	//// Test the AddCommand() function
	//start := time.Now()
	//json_actions.AddCommand("name", str)
	//elapsed := time.Since(start)
	//fmt.Println("AddCommand() execution time:", elapsed)

	//json_actions.FindCommand()
}

// run commands
func runCommands(commandString string) {
	commands := strings.Split(commandString, " &&")

	for _, command := range commands {
		if strings.Contains(command, "sudo") {
			command = strings.Replace(command, "sudo", "sudo -S", 1)
		}
		command = strings.TrimSpace(command)
		fCommand := strings.Fields(command)

		password := "Starship1"

		// run the `sudo` command with the `ls` command as an argument
		cmd := exec.Command(fCommand[0], fCommand[1:]...)

		// set the password for the `sudo` command
		cmd.Stdin = strings.NewReader(password + "\n")

		// run the command
		output, err := cmd.CombinedOutput()

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(output))
	}
}

func checkSudoTime() {
	// Run the sudo -v command to update the user's timestamp
	cmd := exec.Command("sudo", "-v")
	err := cmd.Start()
	if err != nil {
		// Handle error
	}

	// Wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		// Handle error
	}

	// Run the sudo -k command to check the user's timestamp
	cmd = exec.Command("sudo", "-k")
	err = cmd.Start()
	if err != nil {
		// Handle error
	}

	// Wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		// Handle error
	}

	// Check the exit status of the sudo -k command
	if cmd.ProcessState.ExitCode() == 0 {
		// The sudo timeout has not been reached
	} else {
		// The sudo timeout has been reached
	}
}
