package cmd

import (
	"fmt"
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
		fmt.Println(ids)
	},//was this what a receiver function looked like?
}

func init() {
	RootCmd.AddCommand(doCmd)//why RootCmd?
}