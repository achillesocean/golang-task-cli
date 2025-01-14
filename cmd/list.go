package cmd

import (
	"fmt"
	"gophercises/task/db"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Println("There was an error")
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete!")
			return
		}
		fmt.Println("You have the following tasks:") 
		for i, task := range tasks {
			fmt.Printf("%d. %s, %d\n", i+1, task.Value, task.Key)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}