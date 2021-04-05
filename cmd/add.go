package cmd

import (
	"fmt"
	"os"
	"strings"

	"../db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add tasks to the list",
	Run: func(cmd *cobra.Command, args []string) {
		taskName := strings.Join(args, " ")
		_, err := db.CreateTask(taskName)
		if err != nil {
			fmt.Println("Something went wrong", err)
			os.Exit(1)
		}

		fmt.Printf("Added %s  to the list\n", taskName)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
