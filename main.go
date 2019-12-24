package main

import (
	"os"
	"shri3k/bulkes/cmd"
	"shri3k/bulkes/feeder"
)

func main() {

	args := cmd.Cmd()

	if args.Index == "" {
		panic("No index name given. Quitting.")
	}

	index := args.Index

	feeder.Feed(index, os.Stdin, os.Stdout)
}
