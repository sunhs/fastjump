package fastjump

import (
	"encoding/gob"
	"encoding/json"
	"fastjump/utils"
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

func (j *Jumper) updateDB(path string, idx *int) {
	if idx != nil {
		j.db = append(j.db[:*idx], j.db[*idx+1:]...)
	}

	j.db = append([]string{path}, j.db...)
	j.db = j.db[:j.conf.NHistory]

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

func (j *Jumper) Jump(patterns []string) {
}
