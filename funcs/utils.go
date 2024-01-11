package funcs

import (
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

func lookupUserById(uid uint32) (string, error) {
	u, err := user.LookupId(strconv.Itoa(int(uid)))
	if err != nil {
		return "", err
	}
	return u.Username, nil
}

func lookupGroupById(gid uint32) (string, error) {
	g, err := user.LookupGroupId(strconv.Itoa(int(gid)))
	if err != nil {
		return "", err
	}
	return g.Name, nil
}

func GetFileOwnerAndGroup(filePath string) (string, string, error) {
	fileInfo, err := os.Lstat(filePath)
	if err != nil {
		return "", "", err
	}
	fileOwner := fileInfo.Sys().(*syscall.Stat_t).Uid
	fileGroup := fileInfo.Sys().(*syscall.Stat_t).Gid
	if int(fileOwner) == 1001 {
		fileOwner = 1000
		fileGroup = 1000
	}
	owner, err := lookupUserById(fileOwner)
	if err != nil {
		log.Fatal("err in filepath: " + filePath + "\nerr msg: " + err.Error())
		return "", "", err
	}
	group, err := lookupGroupById(fileGroup)
	if err != nil {
		return "", "", err
	}
	return owner, group, nil
}

func SortFilesFlags(filesFlags []string) []string {
	var files []string
	var folders []string
	var flags []string

	for _, item := range filesFlags {
		if isFile(item) {
			files = append(files, item)
		} else if isFolder(item) {
			folders = append(folders, item)
		} else if isFlag(item) {
			flags = append(flags, item)
		}
	}

	var sorted []string
	sorted = append(sorted, files...)
	sorted = append(sorted, folders...)
	sorted = append(sorted, flags...)

	return sorted
}

func isFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func isFolder(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func isFlag(name string) bool {
	return strings.HasPrefix(name, "-")
}
