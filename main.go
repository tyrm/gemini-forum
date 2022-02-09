package main

import (
	"github.com/tyrm/gemini-forum/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err.Error())
	}
}
