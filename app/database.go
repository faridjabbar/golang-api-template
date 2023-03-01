package app

import (
	"log"
	"os"
	"time"
	"voltunes-chick-api-master-product/model/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase(user, host, password, port, db string) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	// dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + db + "?parseTime=true"
	dsn := "host= "+host+" user="+user+" password="+password+" dbname="+db+" port="+port+" sslmode=disable TimeZone=Asia/Jakarta"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	err = database.AutoMigrate(
		&domain.Bank{},
	)
	if err != nil {
		panic("failed to auto migrate schema")
	}

	return database
}
