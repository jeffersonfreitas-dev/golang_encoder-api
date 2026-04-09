package database

import (
	"log"

	"github.com/glebarez/sqlite"
	"github.com/jeffersonfreitas-dev/encoder-api/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	AutoMigrateDb bool
	Env           string
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	instance := NewDb()
	instance.Env = "test"
	instance.DbTypeTest = "sqlite3"
	instance.DsnTest = ":memory:"
	instance.AutoMigrateDb = true

	connection, err := instance.Connect()

	if err != nil {
		log.Fatalf("test db error: %v", err)
	}

	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	if d.Env != "test" {
		d.Db, err = gorm.Open(postgres.Open(d.Dsn), &gorm.Config{})
	} else {
		d.Db, err = gorm.Open(sqlite.Open(d.DsnTest), &gorm.Config{})
	}

	if err != nil {
		return nil, err
	}

	if d.AutoMigrateDb {
		d.Db.AutoMigrate(&domain.Video{}, &domain.Job{})
	}

	return d.Db, nil
}
