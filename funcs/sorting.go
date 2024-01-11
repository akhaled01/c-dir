package funcs

import (
	"os"
	"strings"
)

func isHidden(name string) bool {
	return strings.HasPrefix(name, ".") && name != "." && name != ".."
}

func Sort(entries []os.FileInfo) {
	for i := 0; i < len(entries)-1; i++ {
		for j := i + 1; j < len(entries); j++ {
			wordAtI := ""
			if isHidden(entries[i].Name()) {
				wordAtI = entries[i].Name()[1:]
			} else {
				wordAtI = entries[i].Name()
			}
			wordAtJ := ""
			if isHidden(entries[j].Name()) {
				wordAtJ = entries[j].Name()[1:]
			} else {
				wordAtJ = entries[j].Name()
			}
			if strings.ToLower(wordAtI) > strings.ToLower(wordAtJ) {
				entries[i], entries[j] = entries[j], entries[i]
			}
		}
	}
}

func sortByReverseTime(mainfs string, entries []os.FileInfo) {
	for i := 0; i < len(entries)-1; i++ {
		for j := i + 1; j < len(entries); j++ {
			info1, _ := os.Stat(mainfs + "/" + entries[i].Name())
			info2, _ := os.Stat(mainfs + "/" + entries[j].Name())
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

func reverseSortAlphabet(entries []os.FileInfo) {
	for i := 0; i < len(entries)-1; i++ {
		for j := i + 1; j < len(entries); j++ {
			wordAtI := ""
			if isHidden(entries[i].Name()) {
				wordAtI = entries[i].Name()[1:]
			} else {
				wordAtI = entries[i].Name()
			}
			wordAtJ := ""
			if isHidden(entries[j].Name()) {
				wordAtJ = entries[j].Name()[1:]
			} else {
				wordAtJ = entries[j].Name()
			}
			if strings.ToLower(wordAtI) < strings.ToLower(wordAtJ) {
				entries[i], entries[j] = entries[j], entries[i]
			}
		}
	}
}

func sortByTime(mainfs string, entries []os.FileInfo) {
	for i := 0; i < len(entries)-1; i++ {
		for j := i + 1; j < len(entries); j++ {
			info1, _ := os.Stat(mainfs + "/" + entries[i].Name())
			info2, _ := os.Stat(mainfs + "/" + entries[j].Name())

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
