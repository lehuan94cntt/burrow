package config

import (
	"github.com/hyperledger/burrow/vent/types"
)

const DefaultPostgresDBURL = "postgres://postgres@localhost:5432/postgres?sslmode=disable"

// Flags is a set of configuration parameters
type Flags struct {
	DBAdapter string
	DBURL     string
	DBSchema  string
	GRPCAddr  string
	HTTPAddr  string
	LogLevel  string
	SpecFile  string
	SpecDir   string
	AbiFile   string
	AbiDir    string
	DBBlockTx bool
}

// DefaultFlags returns a configuration with default values
func DefaultFlags() *Flags {
	return &Flags{
		DBAdapter: types.PostgresDB,
		DBURL:     DefaultPostgresDBURL,
		DBSchema:  "vent",
		GRPCAddr:  "localhost:10997",
		HTTPAddr:  "0.0.0.0:8080",
		LogLevel:  "debug",
		SpecFile:  "",
		SpecDir:   "",
		AbiFile:   "",
		AbiDir:    "",
		DBBlockTx: false,
	}
}

func DefaultSqliteFlags() *Flags {
	cfg := DefaultFlags()
	cfg.DBAdapter = types.SQLiteDB
	cfg.DBURL = "./vent.sqlite"
	return cfg
}

func DefaultPostgresFlags() *Flags {
	cfg := DefaultFlags()
	cfg.DBAdapter = types.PostgresDB
	cfg.DBURL = DefaultPostgresDBURL
	return cfg
}
