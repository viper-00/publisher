package database

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	TicketID  string `gorm:"type:varchar(100);unique_index"`
	Content   string `gorm:"type:varchar(100)"`
	Title     string `gorm:"type:varchar(100)"`
	Author    string `gorm:"type:varchar(100)"`
	Topic     string `gorm:"type:varchar(100)"`
	Watermark string `gorm:"type:varchar(100)"`
}

func Init(dbname, host, port, user, password, timeZone string) (*gorm.DB, error) {
	// look there: https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
	dsn := fmt.Sprintf("dbname=%s host=%s port=%s user=%s password=%s TimeZone=%s sslmode=disable", dbname, host, port, user, password, timeZone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, errors.New("don't open database connection")
	}

	return db, nil
}
