package database

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Singleton GORMDB
var GORMDB *gorm.DB

func CreateGormDB() (gormDB *gorm.DB, err error) {
	if GORMDB == nil {
		dbengine := "mysql"
		dsn := DataSource
		var sqlDB *sql.DB
		sqlDB, err = sql.Open(dbengine, dsn)
		if err != nil {
			return
		}
		GORMDB, err = gorm.Open(
			mysql.New(mysql.Config{Conn: sqlDB}),
			&gorm.Config{},
		)
		if err != nil {
			return
		}
	}
	gormDB = GORMDB
	return gormDB, err
}
