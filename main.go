package main

import (
	"log"

	"github.com/feo0o/dip/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
