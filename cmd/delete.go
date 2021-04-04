package cmd

import (
	"fmt"
	"os"
	"strconv"

	"../db"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete tasks from the list",
	Run: func(cmd *cobra.Command, args []string) {
		var id int
		for _, arg := range args {
			var er error
			id, er = strconv.Atoi(arg)
			if er != nil {
				fmt.Println("Failed to parse the argument", arg)
			} else {
				err := db.Delete(id)
				if err != nil {
					fmt.Println("Something went wrong..", err.Error())
					os.Exit(1)
				}
			}
		}
		fmt.Println("Deleted task", id)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
