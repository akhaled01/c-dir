package main

import (
	"os"
	"searchDir/funcs"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		args = append(args, ".")
	}
	mainEntries := funcs.ParseFlags(args)

	for _, v := range mainEntries {
		funcs.PrintRes(v)
	}
}
