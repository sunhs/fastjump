package search

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// LoadDB loads the local DB file.
func (lcsSearch *LCSSearch) LoadDB(fPath string) error {
	lcsSearch.dbPath = fPath
	f, err := os.Open(fPath)
	defer f.Close()
	if err != nil {
		// Remains the DB as empty.
		return nil
	}

	dec := gob.NewDecoder(f)
	err = dec.Decode(&lcsSearch.db)
	if err != nil {
		return fmt.Errorf("problems occurred while loading db %s", fPath)
	}

	return nil
}

// LoadConf loads the local config file.
func (lcsSearch *LCSSearch) LoadConf(fPath string) error {
	data, err := ioutil.ReadFile(fPath)
	if err != nil {
		lcsSearch.conf.NPaths = DefaultNPaths
		return nil
	}
	err = json.Unmarshal(data, &lcsSearch.conf)
	if err != nil {
		return fmt.Errorf("problems occurred while loading config %s", fPath)
	}

	return nil
}

// Search matche a query string with history paths in the DB.
// 1. If the query string is a valid path on the os, simply use string ==
// to match or insert it into the DB.
// 2. If the query string is a search pattern, use LCS to match it
// or return not found.
func (lcsSearch *LCSSearch) Search(query string) (resString string) {
	fInfo, err := os.Stat(query)
	dirExists := err == nil && fInfo.IsDir()

	if dirExists {
		query, _ = filepath.Abs(query)
		resString = query
		for i := 0; i < len(lcsSearch.db); i++ {
			if query == lcsSearch.db[i].Path {
				lcsSearch.updateDB(i, lcsSearch.db[i])
				return
			}
		}
		lcsSearch.updateDB(len(lcsSearch.db), LCSDBItem{Pattern: "", Path: query, Weight: 0})
		return
	}

	for i := 0; i < len(lcsSearch.db); i++ {
		if query == lcsSearch.db[i].Pattern {
			resString = lcsSearch.db[i].Path
			lcsSearch.updateDB(i, lcsSearch.db[i])
			return
		}
		if hit, _ := LCSImpl(query, lcsSearch.db[i].Path); hit {
			resString = lcsSearch.db[i].Path
			lcsSearch.db[i].Pattern = query
			lcsSearch.updateDB(i, lcsSearch.db[i])
			return
		}
	}

	return ""
}

func (lcsSearch *LCSSearch) updateDB(ind int, item LCSDBItem) {
	item.Weight++

	lenDB := len(lcsSearch.db)
	if lenDB > lcsSearch.conf.NPaths {
		lenDB := lcsSearch.conf.NPaths
		if ind > lenDB {
			ind = lenDB
		}
	}

	if ind == lenDB { // is new
		if lenDB == lcsSearch.conf.NPaths { // is full, replace the last one
			ind--
		}
		lcsSearch.db = append(lcsSearch.db[:ind], item)
	} else {
		lcsSearch.db = lcsSearch.db[:lenDB]
	}

	targetInd := ind - 1
	for {
		if targetInd < 0 || lcsSearch.db[targetInd].Weight > item.Weight {
			break
		}
		lcsSearch.db[targetInd+1] = lcsSearch.db[targetInd]
		targetInd--
	}
	lcsSearch.db[targetInd+1] = item

	f, _ := os.Create(lcsSearch.dbPath)
	defer f.Close()

	enc := gob.NewEncoder(f)
	enc.Encode(&lcsSearch.db)
}

// ListDB prints the DB contents.
func (lcsSearch *LCSSearch) ListDB() {
	for i := 0; i < len(lcsSearch.db); i++ {
		item := lcsSearch.db[i]
		fmt.Printf("%s, %s, %d\n", item.Pattern, item.Path, item.Weight)
	}
}
