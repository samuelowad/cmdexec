package runner

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// RunCommand receives a string of commands and runs them depending on the format and the sudo time
func RunCommand(commandString string) {
	commands := strings.Split(commandString, " &&")

	for _, command := range commands {
		if strings.Contains(command, "sudo") {
			command = strings.Replace(command, "sudo", "sudo -S", 1)
		}
		command = strings.TrimSpace(command)
		fCommand := strings.Fields(command)

		password := "YOUR SUDO PASSWORD"

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
