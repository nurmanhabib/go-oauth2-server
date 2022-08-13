package connection

import (
	"fmt"
	"strings"
)

// DSN is a struct for Postgres database connection DSN configuration.
type DSN struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  bool
	Timezone string
}

// ToMySQL export to DSN string format.
func (dsn DSN) ToMySQL() string {
	s := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dsn.User,
		dsn.Password,
		dsn.Host,
		dsn.Port,
		dsn.DBName,
	)

	return s
}

// ToPostgres export to DSN string format.
func (dsn DSN) ToPostgres() string {
	var s []string

	if dsn.Host != "" {
		s = append(s, fmt.Sprintf("host=%s", dsn.Host))
	} else {
		s = append(s, fmt.Sprintf("host=%s", "localhost"))
	}

	if dsn.Port != "" {
		s = append(s, fmt.Sprintf("port=%s", dsn.Port))
	} else {
		s = append(s, fmt.Sprintf("port=%s", "5432"))
	}

	if dsn.User != "" {
		s = append(s, fmt.Sprintf("user=%s", dsn.User))
	}

	if dsn.Password != "" {
		s = append(s, fmt.Sprintf("password=%s", dsn.Password))
	}

	if dsn.DBName != "" {
		s = append(s, fmt.Sprintf("dbname=%s", dsn.DBName))
	}

	if dsn.SSLMode {
		s = append(s, fmt.Sprintf("sslmode=%s", "require"))
	} else {
		s = append(s, fmt.Sprintf("sslmode=%s", "disable"))
	}

	if dsn.Timezone != "" {
		s = append(s, fmt.Sprintf("TimeZone=%s", dsn.Timezone))
	}

	return strings.Join(s, " ")
}
