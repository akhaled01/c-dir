package funcs

import (
	"fmt"
	// "io/fs"
	"os"
	"strings"
	"syscall"
	"strconv"
)

func LFD(path string, grouplen, susInfolen int) {
	info, err := os.Lstat(path)
	if err != nil {
		// Handle the error
		fmt.Println("Error:", err)
		return
	}

	// File permissions
	permissions := info.Mode().String()
	// fix if the permissions have a D at the beginning
	if permissions[0] == 'D' {
		permissions = permissions[1:]
	}


	if permissions[0] == 'L' {
		permissions = "l" + permissions[1:]
	}

	// parse block devices
	if info.Mode()&os.ModeDevice != 0 && info.Mode()&os.ModeCharDevice == 0 {
		permissions = "b" + permissions
	}

	// Number of links
	sysInfo, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		fmt.Println("Error: Unable to retrieve file information")
		return
	}
	numLinks := sysInfo.Nlink
	// fix the spacing in the long format display
	numLinkss:= strconv.Itoa(int(numLinks))
	if susInfolen != 0 {
		numLinkss = strings.Repeat(" ", susInfolen-len(numLinkss))
	} else {
		numLinkss = ""
	}

	// Owner and group
	owner, err := lookupUserById(uint32(sysInfo.Uid))
	if err != nil {
		// Handle the error
		fmt.Println("Error:", err)
		return
	}
	group, err := lookupGroupById(uint32(sysInfo.Gid))
	if err != nil {
		// Handle the error
		fmt.Println("Error:", err)
		return
	}
	// fix the spacing in the long format display
	if grouplen != 0 {
		group = group + strings.Repeat(" ", grouplen-len(group))
	}

	// File size
	size := info.Size()

	// Last modified time
	modTime := info.ModTime().Format("Jan _2 15:04")

	// File name
	fileName := info.Name()
	ftype := permissions[0]

	switch ftype {
	case 'd':
		fileName = blueANSI + boldANSI + fileName + resetANSI
	case 'l':
		fileName = cyanANSI + boldANSI + fileName + resetANSI
	case 'c', 'b':
		fileName = blackBgANSI + yellowANSI + boldANSI + fileName + resetANSI
	case 'p':
		fileName = blackBgANSI + yellowANSI + fileName + resetANSI
	case 's':
		fileName = magentaANSI + fileName + resetANSI
	}

	// Print the formatted information
	// check if symlink first
	isSymLink, symlinkdest, err := IsSymlink(path)
	if err != nil {
		fmt.Println("[LFD SYMLINK ERR]:", err)
		os.Exit(1)
	}
	if isSymLink {
		fileName += " -> " + symlinkdest
	}
	fmt.Printf("%s%s %1d %s %s %4d %s %s\n", permissions,numLinkss, numLinks, owner, group, size, modTime, fileName)
}

// find the maximum group length so we can fix the spacing in the long format display
func MaxGroupLength(mainfs string, entries []os.FileInfo) int {
	max := 0
	for _, entry := range entries {
		info, err := os.Stat(mainfs + "/" + entry.Name())
		if err != nil {
			// Handle the error
			fmt.Println("Error:", err)
			return 0
		}
		sysInfo, ok := info.Sys().(*syscall.Stat_t)
		if !ok {
			fmt.Println("Error: Unable to retrieve file information")
			return 0
		}
		group, err := lookupGroupById(uint32(sysInfo.Gid))
		if err != nil {
			// Handle the error
			fmt.Println("Error:", err)
			return 0
		}
		if len(group) > max {
			max = len(group)
		}
	}
	return max
}

func MaxSusInfoLength(mainfs string, entries []os.FileInfo) int {
	max := 0
	for _, entry := range entries {
		info, err := os.Stat(mainfs + "/" + entry.Name())
		if err != nil {
			// Handle the error
			fmt.Println("Error:", err)
			return 0
		}
		sysInfo, ok := info.Sys().(*syscall.Stat_t)
		if !ok {
			fmt.Println("Error: Unable to retrieve file information")
			return 0
		}
		numLinks := sysInfo.Nlink
		numLinkss:= strconv.Itoa(int(numLinks))
		if len(numLinkss) > max {
			max = len(numLinkss)
		}
	}
	return max
}

// check if a file is a symbolic link
func IsSymlink(filename string) (bool, string, error) {
	fileInfo, err := os.Lstat(filename)
	if err != nil {
		return false, "", err
	}

	if fileInfo.Mode()&os.ModeSymlink != 0 {
		linkPath, err := os.Readlink(filename)
		if err != nil {
			return true, "", err
		}
		return true, linkPath, nil
	}

	// Not a symlink
	return false, "", nil
}
