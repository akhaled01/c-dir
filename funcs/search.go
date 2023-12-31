package funcs

import (
	"io/fs"
	"os"
	"fmt"
)

func SearchDir(dir string) ([]fs.DirEntry, []fs.DirEntry) {
	entries, err := os.ReadDir(dir)
	mainEntries := []fs.DirEntry{}
	DirEntries := []fs.DirEntry{}
	mainEntries = append(mainEntries, entries...)
	if err != nil {
		fmt.Println(RedANSI+BoldANSI+"[search.go] error searching directory,", err)
	}
	for _, v := range mainEntries {
		if v.IsDir() {
			DirEntries = append(DirEntries, v)
		}
	}
	return mainEntries, DirEntries
}
