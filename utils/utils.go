package utils

import (
	"encoding/json"
	"fastjump/search"
	"io/ioutil"
	"os"
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

// CheckConfigFile checks and creates config file.
func CheckConfigFile() (confPath string, dbPath string) {
	fileDir := ExpandUser(search.FileDir)
	os.MkdirAll(fileDir, 0740)
	confPath = filepath.Join(fileDir, search.ConfFileName)
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		config := search.Config{NPaths: search.DefaultNPaths}
		data, _ := json.Marshal(&config)
		ioutil.WriteFile(confPath, data, 0740)
	}
	dbPath = filepath.Join(fileDir, search.DBFileName)
	return
}
