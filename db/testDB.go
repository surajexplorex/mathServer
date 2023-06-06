package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"mathOperation/config"
	"strconv"
	"time"
)

// Global vars to store the DB open connections
var db *gorm.DB
var sqlDB *sql.DB

// InitTestDB initialize the db connection
func InitTestDB() error {

	c := config.InitEnvConfigs()

	// print the env variables
	log.Printf("DB Port: %d \n", c.DBPort)
	log.Printf("DB HostName :%s \n", c.DBHostName)

	var err error
	dsn := c.DBUserName + ":" + c.DBDatabasePassword + "@tcp" + "(" + c.DBHostName + ":" + strconv.Itoa(c.DBPort) + ")/" + c.DBDatabaseName + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("error connecting to database : Err : %w", err)
	}

	sqlDB, err = db.DB()

	sqlDB.SetMaxOpenConns(c.DBMaxOpenConnection)
	sqlDB.SetMaxIdleConns(c.DBMaxIdleConnection)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = sqlDB.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return fmt.Errorf("errors pinging DB: %w", err)
	}
	log.Printf("Connected to DB %s successfully\n", c.DBDatabaseName)

	return nil
}

// GetDB get the db instance.
func GetDB() *gorm.DB {
	return db
}
