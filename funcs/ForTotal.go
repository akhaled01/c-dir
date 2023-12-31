package funcs

import (
	"io/fs"
	"os"
	"syscall"
)

func calculateTotal(entries []fs.DirEntry, path string) int64 {
	size := int64(0)
	for i := 0; i < len(entries); i++ {
		if entries[i].Name()[0] != '.' || DisplayHidden {
			subPath := ""
			subPath = ReturnPath(entries[i].Name(), path)
			fileInfo, err := os.Stat(subPath)
			if err == nil {
				stat, ok := fileInfo.Sys().(*syscall.Stat_t)
				if ok {
					size += stat.Blocks
				}
			}
		}
	}
	return size / 2
}

func ReturnPath(fileName, path string) string {
	if path != "./" && rune(path[len(path)-1]) != '/' {
		return path + "/" + fileName
	}
	return fileName
}


