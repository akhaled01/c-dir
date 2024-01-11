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
	args = funcs.ParseFlags(args)
	args = funcs.SortFilesFlags(args)

	for _, v := range args {
		if funcs.IsSingleFlag(v) || funcs.IsMultiFlag(v) {
			continue
		} else {

			funcs.PrintRes(v)
		}
	}
}
