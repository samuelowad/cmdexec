package cmd

import (
	"fmt"
	json_actions "github.com/samuelowad/cmdexec/pkg/json-actions"
	"github.com/samuelowad/cmdexec/pkg/runner"
	"github.com/spf13/cobra"
	"strings"
)

func CreateCLi() {

	var cmdAdd []string
	var cmdRun, cmdDel string
	var cmdLoad, cmdHelp bool

	rootCmd := &cobra.Command{
		Use:   "cmdexc",
		Short: "The cli app for running cmd exec",
	}

	rootCmd.PersistentFlags().StringArrayVarP(&cmdAdd, "add", "a", []string{}, "Add a command")
	rootCmd.PersistentFlags().StringVarP(&cmdRun, "run", "r", "", "Run a command")
	rootCmd.PersistentFlags().StringVarP(&cmdDel, "delete", "d", "", "Delete a command")
	rootCmd.PersistentFlags().BoolVarP(&cmdLoad, "load", "l", false, "Load all command")
	rootCmd.PersistentFlags().BoolVarP(&cmdHelp, "help", "h", false, "help")

	rootCmd.Flag("add").Value.Set("")
	rootCmd.Flag("add").Usage = "Add a command"
	rootCmd.Flag("add").DefValue = ""
	rootCmd.Flag("add").NoOptDefVal = ""
	//rootCmd.Flag("add").StringSliceVarP(&cmdAdd)

	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {

		if len(cmdAdd) == 0 && cmdRun == "" && cmdDel == "" && !cmdLoad {
			fmt.Println("Usage: cmdexec [flags] [arguments]")
			cmd.Usage()
		}
	})

	//if len(cmdAdd) == 0 && !cmdRun && !cmdDel && !cmdLoad {
	//	fmt.Println(len(cmdAdd) == 0, "add")
	//	fmt.Println(!cmdRun, "run")
	//	fmt.Println(!cmdDel, "del")
	//	fmt.Println(!cmdLoad, "load")
	//	fmt.Println("Error: at least one command is required, run app -h")
	//	return
	//}

	//if !rootCmd.Flag("add").Changed && !cmdRun && !cmdDel && !cmdLoad {
	//
	//	fmt.Println("Error: at least one command is required, run app -h")
	//	return
	//}

	rootCmd.Execute()

	if cmdHelp {
		rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
			fmt.Println("Usage: app [flags] [arguments]")
			fmt.Println("\nFlags:")
			err := cmd.Usage()
			if err != nil {
				return
			}
		})
		return

	}

	if len(cmdAdd) > 0 {
		//loop through cmdAdd and call EncodeJson function
		for _, commandst := range cmdAdd {
			if len(commandst) < 1 {
				continue
			}
			json_actions.EncodeJson(getCommandFromString(commandst))
		}

		fmt.Println("Command(s) added:", strings.Join(cmdAdd, ", "))
	}
	if cmdRun != "" {
		cmd := json_actions.FindCommand(cmdRun)
		runner.RunCommand(cmd)
		fmt.Println("Command ran", cmdRun)
	}
	if cmdDel != "" {
		err := json_actions.DeleteCommand(cmdDel)
		if err != nil {
			return
		}
		fmt.Println("Command deleted", cmdDel)
	}
	if cmdLoad {
		commands, err := json_actions.LoadCommand()
		if err != nil {
			return
		}

		for i, com := range commands {
			fmt.Printf("%d. -name: %s -command: %s\n", i+1, com.Name, com.Command)
		}
		fmt.Println("Command loaded")
	}

	//var rootCmd = &cobra.Command{
	//	Use:   "myapp",
	//	Short: "My app is a CLI app that does some things",
	//	Long:  "My app is a CLI app that does some things. You can use it to add, run, delete, and list commands.",
	//	Run: func(cmd *cobra.Command, args []string) {
	//		// If no command is specified, show the help message
	//		cmd.Help()
	//	},
	//}
	//
	//var addCmd = &cobra.Command{
	//	Use:   "add [command]",
	//	Short: "Add a command",
	//	Long:  "Add a command to the list of commands that can be run.",
	//	Args:  cobra.MinimumNArgs(1),
	//	Run: func(cmd *cobra.Command, args []string) {
	//		var commands []string
	//		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
	//			if flag.Name == "a" {
	//				commands = append(commands, flag.Value.String())
	//			}
	//		})
	//		fmt.Println("Commands added:", strings.Join(commands, " "))
	//
	//	},
	//}
	//addCmd.Flags().StringSliceP("a", "a", []string{}, "add a command")
	//
	//rootCmd.AddCommand(addCmd)
	//
	//var runCmd = &cobra.Command{
	//	Use:   "run [command]",
	//	Short: "Run a command",
	//	Long:  "Run a command that has been added to the list of commands.",
	//	Args:  cobra.MinimumNArgs(1),
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("Command ran:", strings.Join(args, " "))
	//	},
	//}
	//rootCmd.AddCommand(runCmd)
	//
	//var deleteCmd = &cobra.Command{
	//	Use:   "delete [command]",
	//	Short: "Delete a command",
	//	Long:  "Delete a command from the list of commands.",
	//	Args:  cobra.MinimumNArgs(1),
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("Command deleted:", strings.Join(args, " "))
	//	},
	//}
	//rootCmd.AddCommand(deleteCmd)
	//
	//var listCmd = &cobra.Command{
	//	Use:   "list",
	//	Short: "List commands",
	//	Long:  "List the commands that have been added.",
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("Commands loaded:")
	//		// Add code to list the commands here
	//	},
	//}
	//rootCmd.AddCommand(listCmd)
	//
	//rootCmd.Execute()
}

func getCommandFromString(commandString string) (string, string) {
	commandString = strings.TrimSpace(commandString)
	// Split on `,`
	fields := strings.Split(commandString, ",")
	var st json_actions.FileStruct

	// Iterate over the fields
	for _, field := range fields {
		// Split on `:`
		parts := strings.Split(field, ":")
		if len(parts) != 2 {
			// Invalid field
			continue
		}

		// Trim leading and trailing whitespace
		name := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Set the field value
		switch name {
		case "name":
			st.Name = value
		case "command":
			st.Command = value
		}
	}

	return st.Name, st.Command
}
