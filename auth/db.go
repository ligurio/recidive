package passwordless

import (
	"database/sql"
	"log"
	"os"

	"github.com/coopernurse/gorp"
)

var dbmap *gorp.DbMap

func init() {

	dbmap = &gorp.DbMap{
		Db:      db,
		Dialect: gorp.sqlite3{},
	}

	if os.Getenv("DEBUG") == "true" {
		dbmap.TraceOn("[gorp]", log.New(os.Stdout, "passwordless:", stdlog.Lmicroseconds))
	}

	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "Id")
}
