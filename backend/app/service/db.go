package service

import (
	"log"
	"sync"

	"github.com/alex-1900/sparkling/app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	main *gorm.DB
}

var (
	serviceDB *DB
	onceDB    sync.Once
)

func GetDB() *DB {
	onceDB.Do(func() {
		serviceDB = &DB{
			main: connectMainDB(serviceDB),
		}
	})
	return serviceDB
}

// Main 获取主数据库实例
func (s *DB) Main() *gorm.DB {
	return s.main
}

// Migrate 用以执行数据库迁移脚本
func (s *DB) Migrate() {
	m := s.Main()
	// register migrations
	m.AutoMigrate(new(model.Authorization))
	m.AutoMigrate(new(model.User))
}

// connectMainDB for create the main db connection
func connectMainDB(d *DB) *gorm.DB {
	main, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:sparkling@tcp(mysql:3306)/sparkling?charset=utf8mb4", // data source name
		DefaultStringSize: 255,
	}), &gorm.Config{})

	if err == nil {
		setPool(main)
		return main
	}
	log.Fatalf("unable to connect to database: %v", err)
	return nil
}

// setPool setuo the connection pull for db
func setPool(db *gorm.DB) *gorm.DB {
	pool, err := db.DB()
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
	pool.SetConnMaxIdleTime(10)
	pool.SetMaxOpenConns(50)
	pool.SetConnMaxLifetime(0)
	return db
}
