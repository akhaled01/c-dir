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

// TODO: fix -a flag
// TODO: fix the Symbolic link.
// TODO: try the audit questions.
func PrintRes(mainfs string) {
	grouplen := 0
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
			// fmt.Println(fs.FormatFileInfo(stat)) تبًا لك.
			LFD(mainfs, grouplen)
		}
		return
	}
	entries, dirs := SearchDir(mainfs)
	// Calculate total size if long format is enabled
	if LongFormat {
		total := calculateTotal(entries, mainfs)
		fmt.Println("Total", total)
	}
	// Sort the mainEntries slice alphabetically
	Sort(entries)
	if Timesort && ReverseOrder {
		sortByReverseTime(entries)
	} else if Timesort {
		sortByTime(entries)
	} else if ReverseOrder {
		reverseSortAlphabet(entries)
	}
	grouplen = MaxGroupLength(mainfs, entries)
	for _, entry := range entries {
		if !LongFormat {
			PrintNormal(entry)
		} else {
			_, err := entry.Info()
			if err != nil {
				fmt.Println(RedANSI+BoldANSI+"[printresults.go] error printing res,", err)
			}
			// fmt.Println(fs.FormatFileInfo(info)) تبًا لك.
			LFD(mainfs + "/" + entry.Name(), grouplen)
		}
	}
	if RecursiveSearch {
		for _, subFS := range dirs {
			fmt.Println(GreenANSI + BoldANSI + mainfs + "/" + subFS.Name() + ResetANSI)
			PrintRes(mainfs + "/" + subFS.Name())
		}
	}
}
