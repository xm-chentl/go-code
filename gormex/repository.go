package gormex

import (
	"github.com/xm-chentl/go-code/dbfactory"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r repository) Create(entry dbfactory.IDbModel) (err error) {
	_ = r.db.Create(entry)
	return
}

func (r repository) Delete(entry dbfactory.IDbModel) (err error) {
	_ = r.db.Delete(entry)
	return
}

func (r repository) Update(entry dbfactory.IDbModel, fields ...interface{}) (err error) {
	if len(fields) > 0 {
		// table := metadata.Get(entry)
		_ = r.db.Model(entry).Updates(nil)
	} else {
		_ = r.db.Save(entry)
	}

	return
}

func (r repository) Query(entry dbfactory.IDbModel) dbfactory.IQuery {
	return &query{
		db:    r.db,
		entry: entry,
	}
}
