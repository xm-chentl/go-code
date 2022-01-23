package gormex

import (
	"github.com/xm-chentl/go-code/dbfactory"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Option func() *gorm.DB

type factory struct {
	dsn string

	db *gorm.DB
}

func (f *factory) Db(args ...interface{}) dbfactory.IRepository {
	if len(args) > 0 {
		return &repository{
			db: f.db.Begin(),
		}
	}

	return &repository{
		db: f.db,
	}
}

func (f factory) Uow() dbfactory.IUnitOfWork {
	return &unitOfWork{
		tx: f.db,
	}
}

func New(dsn string) dbfactory.IDbFactory {
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}

	return &factory{
		db:  db,
		dsn: dsn,
	}
}
