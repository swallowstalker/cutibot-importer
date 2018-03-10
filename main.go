package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/swallowstalker/cutibot-importer/app"
)

func main() {
	fmt.Println("Importing cutibot data")

	if len(os.Args) <= 1 {
		return errors.New("please provide filename to import")
	}

}
