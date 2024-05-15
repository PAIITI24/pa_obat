package helper

import (
	"github.com/hakushigo/pa_c_obat/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

var (
	config_db = struct {
		username string
		password string
		hostname string
		port     int
		dbname   string
	}{
		"root",
		"",
		"localhost",
		3306,
		"apotek_obat",
	}
)

func DB() *gorm.DB {
	db, err := gorm.Open(mysql.New(
		mysql.Config{
			DSN: config_db.username + ":" + config_db.password + "@tcp(" + config_db.hostname + ":" + strconv.Itoa(config_db.port) + ")/" + config_db.dbname + "?parseTime=true",
		}))

	if err != nil {
		panic(err)
	}

	return db
}

func Migrator() {
	db := DB()

	err := db.AutoMigrate(
		&model.Obat{},
		&model.KategoriObat{},
	)

	if err != nil {
		panic(err)
	}
}
