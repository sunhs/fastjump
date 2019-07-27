package main

import (
	"errors"
	"fastjump/search"
	"fastjump/utils"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type myFlags []string

func (x *myFlags) String() string {
	return strings.Join(*x, ",")
}

func (x *myFlags) Set(value string) error {
	values := strings.Split(value, ",")
	if len(values) != 2 {
		return errors.New("should contains exactly 2 values")
	}
	for _, v := range values {
		*x = append(*x, v)
	}
	return nil
}

func main() {
	list := flag.Bool("l", false, "List DB contents.")
	rm := flag.Int("r", -1, "Remove a record in the DB.")
	var mod myFlags
	flag.Var(&mod, "m", "Modify a record in the DB.")
	flag.Usage = func() {
		message := "Usage:\n" +
			"  %s directory/pattern\n" +
			"  %s -l\n" +
			"  %s -r i\n" +
			"  %s -m i,newPattern\n"
		fmt.Fprintf(flag.CommandLine.Output(), message, os.Args[0], os.Args[0], os.Args[0], os.Args[0])
	}
	flag.Parse()

	searchPtr := new(search.LCSSearch)
	confPath, dbPath := utils.CheckConfigFile()
	searchPtr.LoadConf(confPath)
	searchPtr.LoadDB(dbPath)

	if *list {
		searchPtr.ListDB()
		return
	}
	if *rm != -1 {
		if err := searchPtr.RmRecord(*rm); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}
	if len(mod) == 2 {
		ind, err := strconv.Atoi(mod[0])
		if err != nil {
			fmt.Println("argument #1 should be an integer")
			os.Exit(1)
		}
		err = searchPtr.UpdateRecord(ind, mod[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("should contain exactly 1 argument")
		os.Exit(1)
	}
	pattern := args[0]
	pattern = utils.ExpandUser(pattern)
	if len(pattern) > 1 && strings.HasSuffix(pattern, "/") {
		pattern = pattern[:len(pattern)-1]
	}

	resString := searchPtr.Search(pattern)
	if resString == "" {
		fmt.Println("no matches found")
		os.Exit(1)
	}
	fmt.Println(resString)
}
