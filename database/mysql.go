package database

import (
	"fmt"
	"go-template/entity"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitMySQL() (*gorm.DB, error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))

	dbConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}
	if (os.Getenv("GORM_DEBUG") == "true"){
		dbConfig.Logger = logger.Default.LogMode(logger.Info)
	}
	
	db, err := gorm.Open(mysql.Open(dsn), dbConfig)
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 COLLATE=utf8_bin")
	
	if err != nil {
		return nil, err
	}
	return db, nil
}


func MigrateMySQL(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.User{},
	)
}