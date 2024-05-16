package main

import (
	"log"
	"os"

	"gitlab.violinorg.ru/dsidyuk/opsassist/cmd"
)

func main() {
	c := cmd.CreateApp()

	defer recoverPanic()

	if err := c.Run(os.Args); err != nil {
		log.Fatal(err) //nolint:gocritic // we try to recover panics, not regular command errors
	}
}

func recoverPanic() {
	if r := recover(); r != nil {
		switch r.(type) {
		case cmd.CommandNotFoundError:
			log.Error(r)
			log.Exit(127)
		default:
			log.Panic(r)
		}
	}
}
