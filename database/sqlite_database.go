package database

import (
	"github.com/go-xorm/xorm"
	"github.com/imamfzn/bukaresep-go/entity"

	// it required for xorm
	_ "github.com/mattn/go-sqlite3"
)

const dbDriver = "sqlite3"

// CreateDatabase will create an implementation of xorm database
// using sqlite3 driver
func CreateDatabase(dbFilename string) (*xorm.Engine, error) {
	db, err := xorm.NewEngine(dbDriver, dbFilename)

	if err != nil {
		return nil, err
	}

	err = db.Sync(new(entity.Recipe))

	return db, err
}
