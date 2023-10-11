package main

import (
	"log"
	"os"

	"github.com/hiimnhan/spoclify/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
