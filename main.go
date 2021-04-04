package main

import (
	"fmt"
	"os"
	"path/filepath"

	"./cmd"
	"./db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	errorHandler(db.Init(dbPath))
	errorHandler(cmd.RootCmd.Execute())
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
