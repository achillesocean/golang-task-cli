package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" to your task list.\n", task) // formatted string!?
	},
}

// you will need to link this command into the application
func init() {
	RootCmd.AddCommand(addCmd)
}
// init is a func that runs before the main func; and you can have many of them.
