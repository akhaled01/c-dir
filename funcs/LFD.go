package funcs

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
	"syscall"
)

// TODO: get Long Format Display Done -- DONE.
func LFD(path string, grouplen int) {
	info, err := os.Stat(path)
	if err != nil {
		// Handle the error
		fmt.Println("Error:", err)
		return
	}

	// File permissions
	permissions := info.Mode().String()
	//fix if the permissions have a D at the beginning
	if permissions[0] == 'D' {
		permissions = permissions[1:]
	}
	//fix if the permission is "rw-rw----" by adding a "b" at the beginning
	if permissions == "rw-rw----" {
		permissions = "b" + permissions
	}

	// Number of links
	sysInfo, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		fmt.Println("Error: Unable to retrieve file information")
		return
	}
	numLinks := sysInfo.Nlink

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
	//fix the spacing in the long format display
	if grouplen != 0 {
		group = group + strings.Repeat(" ", grouplen-len(group))
	}

	// File size
	size := info.Size()

	// Last modified time
	modTime := info.ModTime().Format("Jan _2 15:04")

	// File name
	fileName := info.Name()

	// Print the formatted information
	fmt.Printf("%s %1d %s %s %4d %s %s\n", permissions, numLinks, owner, group, size, modTime, fileName)
}

// find the maximum group length so we can fix the spacing in the long format display
func MaxGroupLength(mainfs string, entries []fs.DirEntry) int {
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
