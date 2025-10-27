package main

import (
	"fmt"
	"os"
	"CloudStorage-client/internal/client/commands"
)

func main(){
	if err := commands.RootCmd.Execute(); err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}