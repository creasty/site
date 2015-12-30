package store

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/creasty/site/utils"
)

var Database *gorm.DB

func initDatabase() {
	db, err := gorm.Open("postgres", utils.Config.DatabaseUrl)

	if err != nil {
		panic(err)
	}

	// health check
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}

	db.LogMode(true)

	Database = &db

	databaseAutoMigrate()
}

type DatabaseTransactionHandler func(*gorm.DB) error

func DatabaseTransaction(fn DatabaseTransactionHandler) (err error) {
	tx := Database.Begin()
	err = fn(tx)

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}

	return
}

func DatabaseWithoutLogging(fn func()) {
	Database.LogMode(false)
	defer Database.LogMode(true)

	fn()
}

func databaseAutoMigrate() {
	conf, err := goose.NewDBConf("db", utils.Config.Env, "")
	if err != nil {
		panic(err)
	}

	target, err := goose.GetMostRecentDBVersion(conf.MigrationsDir)
	if err != nil {
		panic(err)
	}

	if err := goose.RunMigrations(conf, conf.MigrationsDir, target); err != nil {
		panic(err)
	}
}
