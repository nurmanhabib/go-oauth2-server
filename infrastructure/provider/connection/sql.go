package connection

import (
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/nurmanhabib/go-oauth2-server/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	driverMysql    = "mysql"
	driverPostgres = "postgres"
)

// NewDBConnection is a function to establish a database connection.
func NewDBConnection(config *config.Config) (*gorm.DB, error) {
	dsn := &DSN{
		Host:     config.Database.DBHost,
		Port:     config.Database.DBPort,
		User:     config.Database.DBUser,
		Password: config.Database.DBPassword,
		DBName:   config.Database.DBName,
		SSLMode:  false,
		Timezone: config.Database.DBTimeZone,
	}

	switch config.Database.DBDriver {
	case driverMysql:
		return gorm.Open(mysql.Open(dsn.ToMySQL()), &gorm.Config{})

	case driverPostgres:
		return gorm.Open(postgres.Open(dsn.ToPostgres()), &gorm.Config{})

	default:
		return nil, errors.New("common.error.unknown_database_driver")
	}
}
