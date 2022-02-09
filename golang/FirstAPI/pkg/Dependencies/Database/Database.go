package Database

import (
	"FirstAPI/pkg/Users"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

type Database struct {
	Context *sql.DB
	UserRepository *Users.Repository
}

func InitDatabase() *Database {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "golang",
	}

	// Get a database handle.
	var err error
	var db *sql.DB

	fmt.Println("Start connecting to database")

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected to database!")

	return &Database{
        Context: db,
        UserRepository: Users.GetRepository(db),
    }
}
