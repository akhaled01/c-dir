package funcs

import (
	"io/fs"
	"strings"
)

func Sort(entries []fs.DirEntry) {
	for i := 0; i < len(entries)-1; i++ {
		for j := i + 1; j < len(entries); j++ {
			if entries[i].Name() > entries[j].Name() {
				entries[i], entries[j] = entries[j], entries[i]
			}
		}
	}
}

func sortByReverseTime(entries []fs.DirEntry) {
	for i := 0; i < len(entries)-1; i++ {
		for j := i + 1; j < len(entries); j++ {
			info1, _ := entries[i].Info()
			info2, _ := entries[j].Info()

			time1 := info1.ModTime()
			time2 := info2.ModTime()

			if time1.Equal(time2) {
				if strings.ToLower(entries[i].Name()) > strings.ToLower(entries[j].Name()) {
					entries[i], entries[j] = entries[j], entries[i]
				}
			} else if time1.After(time2) {
				entries[i], entries[j] = entries[j], entries[i]
			}
		}
	}
}

func reverseSortAlphabet(entries []fs.DirEntry) {
	for i := 0; i < len(entries)-1; i++ {
		for j := i + 1; j < len(entries); j++ {
			if entries[i].Name() < entries[j].Name() {
				entries[i], entries[j] = entries[j], entries[i]
			}
		}
	}
}

func sortByTime(entries []fs.DirEntry) {
	for i := 0; i < len(entries)-1; i++ {
		for j := i + 1; j < len(entries); j++ {
			info1, _ := entries[i].Info()
			info2, _ := entries[j].Info()

			time1 := info1.ModTime()
			time2 := info2.ModTime()

			if time1.Equal(time2) {
				if strings.ToLower(entries[i].Name()) < strings.ToLower(entries[j].Name()) {
					entries[i], entries[j] = entries[j], entries[i]
				}
			} else if time1.Before(time2) {
				entries[i], entries[j] = entries[j], entries[i]
			}
		}
	}
}