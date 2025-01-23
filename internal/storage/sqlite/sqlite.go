package sqlite

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const storage_path = "./internal/storage/"

func Connect(dbname string) (*sql.DB, error) {
	path := storage_path + dbname

	log.Println("starting sqlite database at", path)
	database, err := sql.Open("sqlite3", path)

	if err != nil {
		log.Println("error check", err)
		return nil, err
	}

	err = database.Ping()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}

	database.SetMaxIdleConns(10)
	database.SetConnMaxLifetime(time.Minute * 5)

	return database, nil
}

/* func Test(storage_path string) {
	const create string = `
		CREATE TABLE IF NOT EXISTS activities (
			id INTEGER NOT NULL PRIMARY KEY,
			time DATETIME NOT NULL,
			description TEXT
		);`

	database, err := sql.Open("sqlite3", storage_path)

	sql, err := database.Exec(create)

	log.Println("Create res", sql)

	if err != nil {
		log.Println("Create error", err)
	}
}
*/
