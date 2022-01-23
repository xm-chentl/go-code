package gormex

import (
	"reflect"

	"github.com/xm-chentl/go-code/dbfactory"

	"gorm.io/gorm"
)

type query struct {
	db    *gorm.DB
	entry dbfactory.IDbModel
	args  []interface{}
}

func (q query) Count() (num int64, err error) {
	_ = q.db.Model(
		reflect.New(
			reflect.TypeOf(q.entry),
		).Interface(),
	).Count(&num)
	q.reset()
	return
}

func (q query) ToArray(res interface{}) (err error) {
	if len(q.args) > 0 {
		q.db.Where(q.args[0], q.args[1:]...)
	}
	_ = q.db.Find(res)
	q.reset()
	return
}

func (q *query) Where(args ...interface{}) dbfactory.IQuery {
	if len(args) > 0 {
		q.args = append(q.args, args...)
	}
	return q
}

func (q *query) reset() {
	q.args = make([]interface{}, 0)
}
