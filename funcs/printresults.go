package funcs

import (
	"fmt"
	"strings"
	// "strings"
	// "io/fs"
	"os"
)

func PrintNormal(entry os.FileInfo) {
	if entry.IsDir() {
		fmt.Print(BlueANSI + BoldANSI + entry.Name() + ResetANSI + " ")
	} else {
		// if len(entry.Name()) == longestEntry {
		fmt.Print(entry.Name() + " ")
		// } else {
		// 	fmt.Print(entry.Name() + strings.Repeat(" ", longestEntry-len(entry.Name())) + " ")
		// }
	}
}

// TODO: try the audit questions.
func PrintRes(mainfs string) {
	width, err := getTerminalWidth()
	if width == 0 {
		width = 59
	}
	grouplen, susInfolen := 0, 0
	stat, err := os.Stat(mainfs)
	if err != nil {
		fmt.Println(RedANSI+BoldANSI+"[printresults.go] getting stat,", err)
		os.Exit(1)
	}

	if !stat.IsDir() {
		if !LongFormat {
			fmt.Println(mainfs)
		} else {
			if err != nil {
				fmt.Println(RedANSI+BoldANSI+"[printresults.go] error printing res,", err)
			}
			// fmt.Println(fs.FormatFileInfo(stat)) تبًا لك.
			LFD(mainfs, grouplen, susInfolen)
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
		sortByReverseTime(mainfs, entries)
	} else if Timesort {
		sortByTime(mainfs, entries)
	} else if ReverseOrder {
		reverseSortAlphabet(entries)
	}
	// Find the longest entry in the slice
	longestEntry := 0
	longestEntry = LongestEntry(entries)
	// Find the longest group name in the slice
	grouplen = MaxGroupLength(mainfs, entries)
	// Find the longest number of links in the slice
	susInfolen = MaxSusInfoLength(mainfs, entries)
	// Calculate the number of columns (need to be fixed)
	columns := width / (longestEntry)
	// Calculate the number of rows
	rows := (len(entries) + columns - 1) / columns
	// Check if the terminal is wide enough to display it in one line to use the normal format func
	check := isZero(entries, width)
	if !LongFormat && check {
		// Create a 2D slice to hold the file names
		grid := make([][]string, rows)
		for i := range grid {
			grid[i] = make([]string, columns)
		}
		// Populate the grid with the file names
		for i, entry := range entries {
			row := (i / columns)
			column := i % columns
			if row < len(grid) && column < len(grid[row]) {
				grid[row][column] = entry.Name()
			}
		}
		// find the longest name in each column (need to be fixed)
		gr := make([]int, len(grid))
		for i := 0; i < len(grid); i++ {
			num := 0
			for j := 0; j < len(grid[i]); j++ {
				if len(grid[i][j]) > num {
					num = len(grid[i][j])
				}
			}
			gr[i] = num
			if i+1 >= len(grid) {
				break
			}
		}
		fmt.Println(grid)
		count := 0
		// Print the grid (we display the grid row by row here)
		for col := 0; col < columns; col++ {
			// fmt.Printf("Column %d: ", col+1)
			for row := 0; row < rows; row++ {
				count++
				fmt.Printf("%s %s", grid[row][col], strings.Repeat(" ", gr[row]-len(grid[row][col])))
				if count%columns == 0 {
					fmt.Println()
				}
			}
			// fmt.Println(count, columns)
			
		}
		// for i := 0; i < len(grid[i]); i++ {
		// 	for j := 0; j < rows; j++ {
		// 		name := grid[j][i]
		// 		if name != "" {
		// 			fmt.Print(name + strings.Repeat(" ", gr[j]-len(name)) + " ")
		// 		}
		// 	}
		// 	fmt.Println()
		// 	if i+1 >= len(grid) {
		// 		break
		// 	}
		// }
	}
	// Print the entries in a list (normal one + if the files can be print in one line)
	for _, entry := range entries {
		if !LongFormat {
			if !check {
				PrintNormal(entry)
			}
		} else {
			_, err := os.Stat(mainfs + "/" + entry.Name())
			if err != nil {
				fmt.Println(RedANSI+BoldANSI+"[printresults.go] error printing res,", err)
			}
			// fmt.Println(fs.FormatFileInfo(info)) تبًا لك.
			LFD(mainfs+"/"+entry.Name(), grouplen, susInfolen)
		}
	}
	if RecursiveSearch {
		for _, subFS := range dirs {
			fmt.Println(GreenANSI + BoldANSI + mainfs + "/" + subFS.Name() + ResetANSI)
			PrintRes(mainfs + "/" + subFS.Name())
		}
	}
	// Print a new line if the entries are displayed in a normal format
	if !LongFormat && !check {
		fmt.Println()
	}
}
