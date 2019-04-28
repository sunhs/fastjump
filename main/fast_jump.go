package main

import (
	"fastjump/search"
	"fastjump/utils"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	list := flag.Bool("l", false, "List DB contents.")
	flag.Parse()

	searchPtr := new(search.LCSSearch)
	confPath, dbPath := utils.CheckConfigFile()
	searchPtr.LoadConf(confPath)
	searchPtr.LoadDB(dbPath)

	if *list {
		searchPtr.ListDB()
		return
	}

	pattern := flag.Args()[0]
	pattern = utils.ExpandUser(pattern)
	if strings.HasSuffix(pattern, "/") {
		pattern = pattern[:len(pattern)-1]
	}

	resString := searchPtr.Search(pattern)
	if resString == "" {
		os.Exit(1)
	}
	fmt.Println(resString)
}
