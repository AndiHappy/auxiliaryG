package tests_test

import (
	"log"
	"math/rand"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	. "gorm.io/gorm/utils/tests"
)

var DB *gorm.DB
var (
	mysqlDSN = "root:@tcp(localhost:3630)/godb?charset=utf8&parseTime=True&loc=Local"
)

func init() {
	var err error
	if DB, err = OpenTestConnection(&gorm.Config{}); err != nil {
		log.Printf("failed to connect database, got error %v", err)
		os.Exit(1)
	} else {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Printf("failed to connect database, got error %v", err)
			os.Exit(1)
		}

		err = sqlDB.Ping()
		if err != nil {
			log.Printf("failed to ping sqlDB, got error %v", err)
			os.Exit(1)
		}

		RunMigrations()
	}
}

func OpenTestConnection(cfg *gorm.Config) (db *gorm.DB, err error) {
	dbDSN := os.Getenv("GORM_DSN")
	switch os.Getenv("GORM_DIALECT") {
	case "mysql":
		log.Println("testing mysql...")
		if dbDSN == "" {
			dbDSN = mysqlDSN
		}
		db, err = gorm.Open(mysql.Open(dbDSN), cfg)
	case "postgres":
		log.Println("testing postgres...")
	case "sqlserver":
		log.Println("testing sqlserver...")
	case "tidb":
		log.Println("testing tidb...")
	default:
		log.Println("testing sqlite3...")
	}

	if err != nil {
		return
	}

	if debug := os.Getenv("DEBUG"); debug == "true" {
		db.Logger = db.Logger.LogMode(logger.Info)
	} else if debug == "false" {
		db.Logger = db.Logger.LogMode(logger.Silent)
	}

	return
}

func RunMigrations() {
	var err error
	allModels := []interface{}{&User{}, &Account{}, &Pet{}, &Company{}, &Toy{}, &Language{}, &Coupon{}, &CouponProduct{}, &Order{}, &Parent{}, &Child{}, &Tools{}}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allModels), func(i, j int) { allModels[i], allModels[j] = allModels[j], allModels[i] })

	DB.Migrator().DropTable("user_friends", "user_speaks")

	if err = DB.Migrator().DropTable(allModels...); err != nil {
		log.Printf("Failed to drop table, got error %v\n", err)
		os.Exit(1)
	}

	if err = DB.AutoMigrate(allModels...); err != nil {
		log.Printf("Failed to auto migrate, but got error %v\n", err)
		os.Exit(1)
	}

	for _, m := range allModels {
		if !DB.Migrator().HasTable(m) {
			log.Printf("Failed to create table for %#v\n", m)
			os.Exit(1)
		}
	}
}
