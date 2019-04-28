package search

// DefaultNPaths : Number of paths to record.
// FileDir : Directory to store config and DB files.
// ConfFileName : Name of the config file.
// DBFileName : Name of the DB file.
const (
	DefaultNPaths = 100
	FileDir       = "~/.fj"
	ConfFileName  = "config.json"
	DBFileName    = "db"
)

// Search interface.
type Search interface {
	LoadDB(string) error
	LoadConf(string) error
	Search(string) string
	ListDB()
	RmRecord(int) error
	UpdateRecord(int, string) error
}

// Config specifies how many history paths and how many history patter to store.
type Config struct {
	NPaths int
}

// LCSDBItem records an item in the DB.
type LCSDBItem struct {
	Pattern string
	Path    string
	Weight  int
}

// LCSSearch implements the Search interface.
type LCSSearch struct {
	conf   Config
	db     []LCSDBItem
	dbPath string
}
