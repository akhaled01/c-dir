package funcs

import (
	"fmt"
	"io/fs"
	"os"
)

func PrintNormal(entry fs.DirEntry) {
	if entry.IsDir() {
		fmt.Println(BlueANSI + BoldANSI + entry.Name() + ResetANSI)
	} else {
		fmt.Println(entry.Name())
	}
}

func PrintRes(mainfs string) {
	stat, err := os.Stat(mainfs)
	if err != nil {
		fmt.Println(RedANSI+BoldANSI+"[printresults.go] getting stat,", err)
	}
	if !stat.IsDir() {
		if !LongFormat {
			fmt.Println(mainfs)
		} else {
			if err != nil {
				fmt.Println(RedANSI+BoldANSI+"[printresults.go] error printing res,", err)
			}
			fmt.Println(fs.FormatFileInfo(stat))
		}
		return
	}
	entries, dirs := SearchDir(mainfs)
	for _, entry := range entries {
		if !LongFormat {
			PrintNormal(entry)
		} else {
			info, err := entry.Info()
			if err != nil {
				fmt.Println(RedANSI+BoldANSI+"[printresults.go] error printing res,", err)
			}
			fmt.Println(fs.FormatFileInfo(info))
		}
	}
	if RecursiveSearch {
		for _, subFS := range dirs {
			fmt.Println(GreenANSI + BoldANSI + mainfs + "/" + subFS.Name() + ResetANSI)
			PrintRes(mainfs + "/" + subFS.Name())
		}
	}
}
