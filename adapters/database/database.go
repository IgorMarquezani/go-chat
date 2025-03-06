package database

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     string
	port     string
	user     string
	password string
	dbName   string
	sslmode  string
)

var pool = sync.Pool{}

func Setup() {
	host = os.Getenv("DATABASE_HOST")
	port = os.Getenv("DATABASE_PORT")
	user = os.Getenv("DATABASE_USER")
	password = os.Getenv("DATABASE_PASSWORD")
	dbName = os.Getenv("DATABASE")
	sslmode = os.Getenv("SSLMODE")
}

func Connect(str string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(str), &gorm.Config{})
}

func GetDBConn(trys uint8) (*gorm.DB, error) {
	v := pool.Get()
	if v != nil {
		return v.(*gorm.DB), nil
	}

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbName, sslmode)

	db, err := Connect(conn)

	for i := 1; i < int(trys) && err != nil; i++ {
		db, err = Connect(conn)
	}

	pool.Put(db)

	return db, err
}
