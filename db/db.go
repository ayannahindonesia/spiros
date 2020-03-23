package db

import (
	"fmt"
	"os"
	"spiros/models"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"

	// import db drivers
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB database instance
var DB *gorm.DB

func init() {
	DBinit()
	Migrate()
}

// DBinit initiates database connection
func DBinit() {
	var connectionString string
	dbAdapter := os.Getenv("SPIROS_DB_ADAPTER")
	switch dbAdapter {
	default:
		panic(fmt.Sprintf("invalid adapter %v", os.Getenv("SPIROS_DB_ADAPTER")))
	case "mysql":
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("SPIROS_DB_USERNAME"),
			os.Getenv("SPIROS_DB_PASSWORD"),
			os.Getenv("SPIROS_DB_HOST"),
			os.Getenv("SPIROS_DB_PORT"),
			os.Getenv("SPIROS_DB_TABLE"))
		break
	case "postgres":
		connectionString = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
			os.Getenv("SPIROS_DB_USERNAME"),
			os.Getenv("SPIROS_DB_PASSWORD"),
			os.Getenv("SPIROS_DB_HOST"),
			os.Getenv("SPIROS_DB_PORT"),
			os.Getenv("SPIROS_DB_TABLE"),
			os.Getenv("SPIROS_DB_SSL"))
		break
	}

	db, err := gorm.Open(dbAdapter, connectionString)
	if err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	if logmode := os.Getenv("SPIROS_DB_LOGMODE"); logmode == "true" || logmode == "false" {
		switch logmode {
		case "false":
			db.LogMode(false)
			break
		case "true":
			db.LogMode(true)
			break
		}
	}

	db.Exec(fmt.Sprintf("SET TIMEZONE TO '%s'", os.Getenv("SPIROS_TIMEZONE")))
	if maxLifeTime, err := strconv.ParseInt(os.Getenv("SPIROS_DB_MAXLIFETIME"), 10, 64); err != nil {
		db.DB().SetConnMaxLifetime(time.Second * time.Duration(maxLifeTime))
	}
	if maxIdleConnection, err := strconv.Atoi(os.Getenv("SPIROS_DB_MAXIDLECONNECTION")); err != nil {
		db.DB().SetMaxIdleConns(maxIdleConnection)
	}
	if maxOpenConnection, err := strconv.Atoi(os.Getenv("SPIROS_DB_MAXOPENCONNECTION")); err != nil {
		db.DB().SetMaxOpenConns(maxOpenConnection)
	}

	DB = db
}

// Migrate updates database structures
func Migrate() {
	err := DB.AutoMigrate(
		&models.Client{},
		&models.User{},
	).Error
	if err != nil {
		panic(err)
	}
}
