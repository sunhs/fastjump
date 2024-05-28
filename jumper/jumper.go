package jumper

import (
	"encoding/gob"
	"encoding/json"
	"fastjump/utils"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	NHistory int    `json:"n_history"`
	NHint    int    `json:"n_hint"`
	Sep      string `json:"sep"`
}

func (c *Config) LoadConfig(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(c)
	if err != nil {
		panic(err)
	}
}

func (c *Config) SaveConfig(path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(c)
	if err != nil {
		panic(err)
	}
}

type Jumper struct {
	dbPath string
	db     []string
	conf   Config
}

func NewJumper(dbPath string) *Jumper {
	dbPath = utils.ExpandUser(dbPath)
	confDir := filepath.Dir(dbPath)

	if info, err := os.Stat(confDir); os.IsNotExist(err) || !info.IsDir() {
		err := os.MkdirAll(confDir, 0750)
		if err != nil {
			panic(err)
		}
	}

	conf := Config{
		NHistory: 200,
		NHint:    10,
		Sep:      string(os.PathSeparator),
	}
	conf.LoadConfig(filepath.Join(confDir, "config.json"))

	j := &Jumper{
		dbPath: dbPath,
		db:     []string{},
		conf:   conf,
	}
	j.loadDB()

	return j
}

func (j *Jumper) loadDB() {
	_, err := os.Stat(j.dbPath)
	if os.IsNotExist(err) {
		return
	}

	file, err := os.Open(j.dbPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&j.db)
	if err != nil {
		panic(err)
	}
}

func (j *Jumper) updateDB(path string, idx int) {
	if idx != -1 {
		j.db = append(j.db[:idx], j.db[idx+1:]...)
	}

	j.db = append([]string{path}, j.db...)
	if len(j.db) > j.conf.NHistory {
		j.db = j.db[:j.conf.NHistory]
	}

	file, err := os.Create(j.dbPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(j.db)
	if err != nil {
		panic(err)
	}
}

func (j *Jumper) Jump(patterns []string) string {
	if len(patterns) == 0 {
		patterns = []string{utils.ExpandUser("~")}
	}

	if len(patterns) == 1 {
		if info, err := os.Stat(patterns[0]); err == nil && info.IsDir() {
			path, _ := filepath.Abs(utils.ExpandUser(patterns[0]))
			idx := -1
			for i, p := range j.db {
				if p == path {
					idx = i
					break
				}
			}
			j.updateDB(path, idx)
			return path
		}
	}

	rst := MatchDispatcher(patterns, j.db, 1, j.conf.Sep)

	if len(rst) == 0 {
		return ""
	}

	idx, matched := rst[0].index, rst[0].path
	j.updateDB(matched, idx)
	return matched
}

func (j *Jumper) Hint(patterns []string) (ret []string) {
	rst := MatchDispatcher(patterns, j.db, j.conf.NHint, j.conf.Sep)
	for _, rstTuple := range rst {
		ret = append(ret, rstTuple.path)
	}
	return
}

func (j *Jumper) Clean() {
	oldLen := len(j.db)

	newDb := []string{}
	for _, path := range j.db {
		if info, err := os.Stat(path); err == nil && info.IsDir() {
			newDb = append(newDb, path)
		}
	}
	j.db = newDb

	file, err := os.Create(j.dbPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(j.db)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Removed %d non existent paths.\n", oldLen-len(j.db))
}
