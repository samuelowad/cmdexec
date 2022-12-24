package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

func CreateCLi() {
	var (
		command string
		name    string
	)
	var rootCmd = &cobra.Command{
		Use:   "addage",
		Short: "Add the command to the list",
		Long:  "Add the your commands to a list",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("The user's age is %s\n", command)
		},
	}

	// Add the age flag to the root command
	rootCmd.Flags().StringVarP(&command, "command", "c", "", "the command")
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "the command name")
	rootCmd.MarkFlagRequired("command")
	rootCmd.MarkFlagRequired("name")
	rootCmd.SetUsageTemplate(`Usage:{{if .Runnable}}
	  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
	  {{.CommandPath}} [command]{{end}}
	
	Error: The -n and -c flag is required.
	`)

	// Parse the command line arguments
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
