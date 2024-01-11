package funcs

import (
	"fmt"
	"os"
	"strconv"
)

func getTerminalWidth() (int, error) {
	widthStr := os.Getenv("COLUMNS")
	if widthStr == "" {
		return 0, fmt.Errorf("failed to retrieve terminal width")
	}

	width, err := strconv.Atoi(widthStr)
	if err != nil {
		return 0, fmt.Errorf("failed to parse terminal width: %v", err)
	}

	return width, nil
}

func LongestEntry(entries []os.FileInfo) int {
	longestEntry := 0
	for _, entry := range entries {
		if len(entry.Name()) > longestEntry {
			longestEntry = len(entry.Name())
		}
	}
	return longestEntry
}

// func isZero(entries []os.FileInfo, width int) bool {
// 	totalNames := ""
// 	for _, entry := range entries {
// 		totalNames+= entry.Name()
// 		// if len(entry.Name()) > longestEntry {
// 		// 	longestEntry = len(entry.Name())
// 		// }
// 	}
// 	if len(totalNames)-width < 0 {
// 		return false
// 	}
// 	return true
// }