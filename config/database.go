package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IMysqlInstance interface {
	Database() (*gorm.DB, error)
}

type mysqlInstance struct {
	db *gorm.DB
}

type MysqlConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func ConnectMysql(cfg MysqlConfig) (IMysqlInstance, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	Database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	Database.Set("gorm:table_options", "CHARSET=utf8mb4")

	return &mysqlInstance{db: Database}, nil
}

func (m *mysqlInstance) Database() (*gorm.DB, error) {
	return m.db, nil
}
