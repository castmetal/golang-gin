package infra_db

import (
	"fmt"

	_config "golang-gin/src/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewDatabaseConnection() (*gorm.DB, error) {
	fmt.Println("Connecting into database ...")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		_config.SystemParams.DB_HOST,
		_config.SystemParams.DB_USER,
		_config.SystemParams.DB_PASSWORD,
		_config.SystemParams.DB_DATABASE_NAME,
		_config.SystemParams.DB_PORT,
		_config.SystemParams.DB_TIME_ZONE,
	)

	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   _config.SystemParams.DB_SCHEMA + ".", // schema name
			SingularTable: false,
		}}

	if _config.SystemParams.ENV == "production" {
		config = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}
	}

	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		return nil, err
	}

	return db, nil
}
