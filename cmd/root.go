package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "Task Manager",
	Short: "CLI task manager",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
