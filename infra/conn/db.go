package conn

import (
	"first-api/app/model"
	"first-api/infra/config"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	conf := config.Db()
	log.Info("connecting to mysql at " + conf.Host + ":" + conf.Port + "...")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", conf.User, conf.Password, conf.Host, conf.Port, conf.DBName)
	dB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		panic(err)
	}
	// sqlDb, err = dB.DB()
	// if err != nil {
	// 	panic(err)
	// }

	db = dB
	db.AutoMigrate(
		&model.User{},
	)
}

func Db() *gorm.DB {
	return db
}
