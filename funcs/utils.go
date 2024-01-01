package funcs

import (
	"log"
	"os"
	"os/user"
	"strconv"
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
