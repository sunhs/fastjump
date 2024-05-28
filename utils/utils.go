package utils

import (
	"os/user"
	"path/filepath"
	"strings"
)

// ExpandUser expands "~" to the current user home directory.
func ExpandUser(path string) (newPath string) {
	if strings.HasPrefix(path, "~") {
		userPtr, _ := user.Current()
		userHome := userPtr.HomeDir
		newPath = filepath.Join(userHome, path[1:])
		return
	}
	newPath = path
	return
}
