package postgres

import (
	"errors"
	"github.com/Mona-bele/rote-notify/infra/database/drive"
	util "github.com/Mona-bele/rote-notify/utils"
	"github.com/jinzhu/gorm"
)

func PrimaryConnectionPostgres(env *util.Env) (*gorm.DB, error) {
	postgres := NewPostgres(env)
	driver := drive.Drive(postgres)
	if driver == "" {
		return nil, errors.New("driver is empty")
	}

	dns := postgres.Connect()
	if dns == "" {
		return nil, errors.New("dns is empty")
	}

	db, err := NewConnection(dns)
	if err != nil {
		return nil, err
	}
	dbConn := db.GetDatabaseConnection()
	if dbConn == nil {
		return nil, errors.New("dbConn is nil")
	}

	return dbConn, nil
}
