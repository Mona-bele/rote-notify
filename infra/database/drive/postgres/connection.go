package postgres

import (
	"os"
	"sync"

	"github.com/jinzhu/gorm"
)

type Connection interface {
	GetDatabaseConnection() *gorm.DB
}

type connection struct {
	db *gorm.DB
}

var (
	db   *gorm.DB
	once = &sync.Once{}
)

func (c *connection) GetDatabaseConnection() *gorm.DB {
	return c.db
}

func NewConnection(dns string) (Connection, error) {

	var err error
	db, err = gorm.Open("postgres", dns)
	if err != nil {
		panic(err)
	}

	db.LogMode(os.Getenv("LOG_DEBUG") == "true")

	return &connection{db}, nil
}
