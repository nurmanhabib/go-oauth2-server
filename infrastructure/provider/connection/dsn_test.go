package connection_test

import (
	"testing"

	"github.com/nurmanhabib/go-oauth2-server/infrastructure/provider/connection"
	"github.com/stretchr/testify/assert"
)

func TestDSN_ToMySQL(t *testing.T) {
	t.Run("if minimal config", func(t *testing.T) {
		dsn := &connection.DSN{
			User:    "john",
			DBName:  "db_test",
			SSLMode: false,
		}

		assert.Equal(t, "john:@tcp(:)/db_test?charset=utf8&parseTime=True&loc=Local", dsn.ToMySQL())
	})

	t.Run("if all options are filled correctly", func(t *testing.T) {
		dsn := &connection.DSN{
			Host:     "192.168.0.1",
			Port:     "5432",
			User:     "john",
			Password: "doe",
			DBName:   "db_test",
			SSLMode:  false,
			Timezone: "Asia/Jakarta",
		}

		assert.Equal(t, "john:doe@tcp(192.168.0.1:5432)/db_test?charset=utf8&parseTime=True&loc=Local", dsn.ToMySQL())
	})

	t.Run("if the password is empty", func(t *testing.T) {
		dsn := &connection.DSN{
			Host:     "192.168.0.1",
			Port:     "5432",
			User:     "john",
			DBName:   "db_test",
			SSLMode:  false,
			Timezone: "Asia/Jakarta",
		}

		assert.Equal(t, "john:@tcp(192.168.0.1:5432)/db_test?charset=utf8&parseTime=True&loc=Local", dsn.ToMySQL())
	})
}

func TestDSN_ToPostgres(t *testing.T) {
	t.Run("if minimal config", func(t *testing.T) {
		dsn := &connection.DSN{
			User:    "john",
			DBName:  "db_test",
			SSLMode: false,
		}

		assert.Equal(t, "host=localhost port=5432 user=john dbname=db_test sslmode=disable", dsn.ToPostgres())
	})

	t.Run("if all options are filled correctly", func(t *testing.T) {
		dsn := &connection.DSN{
			Host:     "192.168.0.1",
			Port:     "5432",
			User:     "john",
			Password: "doe",
			DBName:   "db_test",
			SSLMode:  false,
			Timezone: "Asia/Jakarta",
		}

		assert.Equal(t, "host=192.168.0.1 port=5432 user=john password=doe dbname=db_test sslmode=disable TimeZone=Asia/Jakarta", dsn.ToPostgres())
	})

	t.Run("if the password is empty", func(t *testing.T) {
		dsn := &connection.DSN{
			Host:     "192.168.0.1",
			Port:     "5432",
			User:     "john",
			DBName:   "db_test",
			SSLMode:  false,
			Timezone: "Asia/Jakarta",
		}

		assert.Equal(t, "host=192.168.0.1 port=5432 user=john dbname=db_test sslmode=disable TimeZone=Asia/Jakarta", dsn.ToPostgres())
	})
}
