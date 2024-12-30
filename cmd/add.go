package cmd

import (
	"fmt"
	"gophercises/task/db"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		if len(task) == 0 {
			fmt.Println("Please try again.")
			return
		}
		id, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("There was an error adding your task")
			os.Exit(1)
		}
		
		fmt.Printf("Added task number \"%d\" to your task list.\n", id) // formatted string!?
	},
}

// you will need to link this command into the application
func init() {
	RootCmd.AddCommand(addCmd)
}
// init is a func that runs before the main func; and you can have many of them.
