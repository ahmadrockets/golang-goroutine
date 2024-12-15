package config

import (
	"database/sql"
	"log"
	"os"
)

type SqlLite struct {
	Client *sql.DB
	DBName string
}

func InitSqlLite(config Config) *SqlLite {
	if config.IsResetDB {
		// Remove database
		os.Remove(config.DBName)

		// Recreate database
		log.Println("Creating " + config.DBName + "...")
		file, err := os.Create(config.DBName) // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println(config.DBName + " created")
	}

	sqliteDatabase, _ := sql.Open("sqlite3", "./"+config.DBName) // Open the created SQLite File

	if config.IsResetDB {
		err := createStudentTable(sqliteDatabase)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	return &SqlLite{
		Client: sqliteDatabase,
		DBName: config.DBName,
	}
}

func createStudentTable(dbClient *sql.DB) (err error) {
	createStudentTableSQL := `CREATE TABLE student (
		"idStudent" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"code" TEXT,
		"name" TEXT,
		"program" TEXT		
	  );`

	log.Println("Create student table...")
	statement, err := dbClient.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		return
	}
	_, err = statement.Exec() // Execute SQL Statements
	if err != nil {
		return
	}
	log.Println("Success create student table")
	return
}
