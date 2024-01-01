package funcs

import (
	"fmt"
	// "io/fs"
	"os"
)

func SearchDir(dir string) ([]os.FileInfo, []os.FileInfo) {
	entries, err := os.ReadDir(dir)
	mainEntryArray := []os.FileInfo{}
	for _, ent := range entries {
		if ent.Name()[0] == '.' && !DisplayHidden {
			continue
		}
		info, err := ent.Info()
		if err != nil {
			fmt.Println(RedANSI+BoldANSI+"[search.go] error getting info for entry,", err)
			continue
		}
		mainEntryArray = append(mainEntryArray, info)
	}
	mainEntries := []os.FileInfo{}
	DirEntries := []os.FileInfo{}

	if DisplayHidden {
		file2, err := os.Stat("..")
		if err == nil {
			mainEntries = append(mainEntries, file2)
		}
		file1, err := os.Stat(".")
		if err == nil {
			mainEntries = append(mainEntries, file1)
		}
	}

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