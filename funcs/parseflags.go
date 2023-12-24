package funcs

import (
	"fmt"
	"os"
)

func removeElementAtIndex(arr []string, index int) []string {
	// Check if the index is valid
	if index < 0 || index >= len(arr) {
		fmt.Println("Index out of range")
		return arr
	}

	// Remove the element at the specified index
	arr = append(arr[:index], arr[index+1:]...)

	return arr
}

func IsSingleFlag(s string) bool {
	_, err := os.Stat(s)
	return s[0] == '-' && len(s) == 2 && os.IsNotExist(err)
}

func IsMultiFlag(s string) bool {
	_, err := os.Stat(s)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println(err)
		return false
	}
	return true && os.IsNotExist(err) && s[0] == '-'
}

func ParseFlags(args []string) []string {
	for i, argument := range args {
		if IsSingleFlag(argument) {
			FlagCounter++
			switch argument {
			case "-a":
				DisplayHidden = true
			case "-R":
				RecursiveSearch = true
			case "-l":
				LongFormat = true
			case "-r":
				ReverseOrder = true
			case "-t":
				Timesort = true
			case "-o":
				DashO = true
			default:
				continue
			}
			args = removeElementAtIndex(args, i)
		} else if IsMultiFlag(argument) {
			FlagCounter++
			runeArray := []rune(argument)
			for _, v := range runeArray[1:] {
				switch v {
				case 'a':
					DisplayHidden = true
				case 'R':
					RecursiveSearch = true
				case 'l':
					LongFormat = true
				case 'r':
					ReverseOrder = true
				case 't':
					Timesort = true
				case 'o':
					DashO = true
				default:
					continue
				}
			}
			args = removeElementAtIndex(args, i)
		}
	}
	if len(args) == 0 {
		args = append(args, ".")
	}
	return args
}
