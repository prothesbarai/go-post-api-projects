package database

import (
	"context"
	"database/sql"
	"go-post-api-projects/logger"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	dsn	:= os.Getenv("DB_DSN")
	db, err := sql.Open("mysql",dsn)
	if(err != nil){
		logger.AppLogger.Error.Println("Database Connection Error : ",err)
	}

	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil{
		logger.AppLogger.Error.Println("Database unreachable : ",err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5*time.Minute)

	DB = db
	logger.AppLogger.Info.Println("Successfully Database Connected : ",db)
}