package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	host       string
	username   string
	password   string
	database   string
	connection *sql.DB
}

func NewDatabase(host string, username string, password string, database string) *Database {
	return &Database{
		host:     host,
		username: username,
		password: password,
		database: database,
	}
}

func (d *Database) Connect() {
	var con = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", d.username, d.password, d.host, d.database)
	db, err := sql.Open("mysql", con)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Connected to Database")
	d.connection = db
}

func (d *Database) GetDB() *sql.DB {
	return d.connection
}

func (d *Database) Insert(query string) {
	insert, err := d.connection.Query(query)
	if err != nil {
		fmt.Printf("Insert query failed %s", err)
	}
	defer insert.Close()
}
