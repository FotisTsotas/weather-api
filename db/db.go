package db

import (
	"database/sql"
	"fmt"
	"log"

	"weather-api/config"
	"weather-api/migrations"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(cfg config.Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Could not connect to database: ", err)
	}

	DB.SetMaxOpenConns(cfg.DBMaxOpenConns)
	DB.SetMaxIdleConns(cfg.DBMaxIdleConns)

	if err = DB.Ping(); err != nil {
		log.Fatal("Could not ping database: ", err)
	}

	if err := migrations.Run(DB); err != nil {
		log.Fatal("Could not run migrations: ", err)
	}
}
