package cmd

import (
	"fmt"
	"gophercises/task/db"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use: "do",
	Short: "Mark a task completed",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int //remember how to create a slice?
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id) // what's the O complexity of this?
			}
		}
		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Println("Something went wrong", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Println("Failed to mark \"%d\" as completed", id)
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", id)
			}
		}
	},//was this what a receiver function looked like?
}

func init() {
	RootCmd.AddCommand(doCmd)//why RootCmd?
}