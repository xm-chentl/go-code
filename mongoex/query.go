package mongoex

import (
	"context"
	"errors"
	"reflect"

	"github.com/xm-chentl/go-code/dbfactory"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type query struct {
	dataBase *mongo.Database

	entry dbfactory.IDbModel
	args  []interface{}
}

func (q query) Count() (num int64, err error) {
	if len(q.args) == 0 {
		q.args = append(q.args, bson.M{})
	}
	num, err = q.dataBase.Collection("").CountDocuments(context.Background(), q.args[0])
	return
}

func (q query) ToArray(res interface{}) (err error) {
	resRt := reflect.TypeOf(res)
	resRv := reflect.ValueOf(res)
	if resRt.Kind() == reflect.Ptr {
		if resRt.Elem().Kind() == reflect.Slice {
			resRt = resRt.Elem().Elem()
		} else {
			// todo: err > 非切片
			err = errors.New("res is not slice")
			return
		}
	} else {
		// todo: err >  非指针
		err = errors.New("res is not ptr")
		return
	}

	if len(q.args) == 0 {
		q.args = append(q.args, bson.M{})
	}
	defer q.reset()

	ctx := context.Background()
	cursor, err := q.dataBase.Collection(
		reflect.TypeOf(q.entry).Name(),
	).Find(ctx, q.args[0])
	if err != nil {
		return
	}

	tempInst := reflect.New(resRt).Interface()
	tempSlice := reflect.MakeSlice(reflect.TypeOf(res).Elem(), 0, 0)
	for cursor.Next(ctx) {
		err := cursor.Decode(tempInst)
		if err != nil {
			break
		}
		tempSlice = reflect.Append(tempSlice, reflect.ValueOf(tempInst).Elem())
	}
	resRv.Elem().Set(tempSlice)

	return
}

func (q *query) Where(args ...interface{}) dbfactory.IQuery {
	q.args = append(q.args, args...)
	return q
}

func (q *query) reset() {
	q.args = make([]interface{}, 0)
}
