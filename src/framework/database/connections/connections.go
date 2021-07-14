package connections

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mixnote/mixnote-api-go/configs"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (
	IDatabase interface {
		ConnectDB() (*gorm.DB, error)
	}
)

type (
	mySql struct {
		DATABASE_URL string
		DB_HOST      string
		DB_PORT      int
		DB_USERNAME  string
		DB_PASSWORD  string
		DB_NAME      string
	}

	sqLite3 struct {
		DATABASE_PATH string
	}

	pgSql struct {
		PG_DATABASE_URL string
		PG_HOST         string
		PG_PORT         int
		PG_USERNAME     string
		PG_PASSWORD     string
		PG_NAME         string
	}
)

func MySql() (mysql *mySql) {
	mysql = &mySql{
		DATABASE_URL: os.Getenv("DATABASE_URL"),
		DB_HOST:      os.Getenv("DB_HOST"),
		DB_USERNAME:  os.Getenv("DB_USERNAME"),
		DB_PASSWORD:  os.Getenv("DB_PASSWORD"),
		DB_NAME:      os.Getenv("DB_NAME"),
	}

	if mysql.DB_HOST == "" {
		mysql.DB_HOST = "127.0.0.1"
	}

	port, err := (func() (int, error) {
		port := os.Getenv("DB_PORT")
		if port != "" {
			return strconv.Atoi(port)
		}
		return 3306, nil
	})()

	if err != nil {
		panic("Invalid database port fot mysql")
	}
	mysql.DB_PORT = port
	return
}

func (m *mySql) ConnectDB() (g *gorm.DB,err error) {
	dns := (func() string {
		if m.DATABASE_URL != "" {
			return m.DATABASE_URL
		}
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.DB_USERNAME, m.DB_PASSWORD, m.DB_HOST, m.DB_PORT, m.DB_NAME)
	})()
	g, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dns,
		DefaultStringSize: 256,
	}))
	g.Set("gorm:table_options", "ENGINE=InnoDB")
	return
}

func SqLite3() (sqlite *sqLite3) {
	sqlite = &sqLite3{
		DATABASE_PATH: os.Getenv("DATABASE_PATH"),
	}
	if sqlite.DATABASE_PATH == "" {
		sqlite.DATABASE_PATH = configs.App.DatabasePath
	}
	return
}

func (sq *sqLite3) ConnectDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(sq.DATABASE_PATH))
}

func PgSql() (pg *pgSql) {
	pg = &pgSql{
		PG_DATABASE_URL: os.Getenv("PG_DATABASE_URL"),
		PG_HOST:         os.Getenv("PG_HOST"),
		PG_USERNAME:     os.Getenv("PG_USERNAME"),
		PG_PASSWORD:     os.Getenv("PG_PASSWORD"),
		PG_NAME:         os.Getenv("PG_NAME"),
	}

	if pg.PG_HOST == "" {
		pg.PG_HOST = "127.0.0.1"
	}

	port, err := (func() (int, error) {
		port := os.Getenv("PG_PORT")
		if port != "" {
			return strconv.Atoi(port)
		}
		return 5432, nil
	})()

	if err != nil {
		panic("Invalid database port address for postgres")
	}
	pg.PG_PORT = port
	return
}

func (pg *pgSql) ConnectDB() (*gorm.DB, error) {
	dns := (func() string {
		if pg.PG_DATABASE_URL != "" {
			return pg.PG_DATABASE_URL
		}
		return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", pg.PG_USERNAME, pg.PG_PASSWORD, pg.PG_HOST, pg.PG_PORT, pg.PG_NAME)
	})()
	return gorm.Open(postgres.Open(dns))
}
