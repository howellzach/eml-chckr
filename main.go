package main

import (
	"fmt"

	"github.com/howellzach/eml-chckr/cmd"
)

//TODO: Add makefile

func main() {
	err := cmd.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}
}
