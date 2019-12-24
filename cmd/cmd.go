package cmd

import (
	"flag"
)

type Params struct {
	Index string
}

func Cmd() Params {
	flag.Parse()
	return Params{
		Index: flag.Arg(0),
	}
}
