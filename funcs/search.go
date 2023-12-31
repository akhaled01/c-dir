package funcs

import (
	"fmt"
	"io/fs"
	"os"
)

func SearchDir(dir string) ([]fs.DirEntry, []fs.DirEntry) {
	entries, err := os.ReadDir(dir)
	mainEntryArray := []fs.DirEntry{}
	for _, ent := range entries {
		if ent.Name()[0] == '.' && !DisplayHidden{
			continue
		}
		mainEntryArray = append(mainEntryArray, ent)
	}
	mainEntries := []fs.DirEntry{}
	DirEntries := []fs.DirEntry{}
	mainEntries = append(mainEntries, mainEntryArray...)
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
