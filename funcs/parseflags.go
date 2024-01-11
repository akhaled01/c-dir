package funcs

import (
	"os"
)

func IsSingleFlag(s string) bool {
	_, err := os.Stat(s)
	return s[0] == '-' && len(s) == 2 && os.IsNotExist(err)
}

func IsMultiFlag(s string) bool {
	_, err := os.Stat(s)
	if err != nil && !os.IsNotExist(err) {
		RedPrintln(err)
		return false
	}
	return true && os.IsNotExist(err) && s[0] == '-'
}

func ParseFlags(args []string) []string {
	for _, argument := range args {
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
		}
	}
	if FlagCounter == len(args) {
		args = append(args, ".")
	}
	return args
}
