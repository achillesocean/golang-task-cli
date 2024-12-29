// this is where the root command goes. what happens when we type/run 'task'.
package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{ // why using pointers?
  Use:   "task",
  Short: "Task is a task manager cli",
}