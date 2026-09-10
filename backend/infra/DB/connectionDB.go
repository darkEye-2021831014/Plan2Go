package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() (*sql.DB, error) {
	// Format: username:password@tcp(host:port)/database_name
	host := os.Getenv("MYSQLHOST")
	port := os.Getenv("MYSQLPORT")
	user := os.Getenv("MYSQLUSER")
	pass := os.Getenv("MYSQLPASSWORD")
	name := os.Getenv("MYSQLDATABASE")

	// TiDB Cloud Serverless rejects any non-TLS connection. Local MySQL
	// (127.0.0.1/localhost) doesn't need it, so only enable for real hosts.
	useTLS := host != "127.0.0.1" && host != "localhost" && host != ""

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)
	if useTLS {
		dsn += "?tls=true"
	}

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error creating DB handle:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	fmt.Println("Connected to MySQL Database successfully! 🟢")
	return DB, nil
}
