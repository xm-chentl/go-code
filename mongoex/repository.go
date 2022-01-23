package mongoex

import (
	"context"
	"reflect"

	"github.com/xm-chentl/go-code/dbfactory"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	dataBase *mongo.Database
}

func (r repository) Create(entry dbfactory.IDbModel) (err error) {
	_, err = r.dataBase.Collection(
		r.getCollectionName(entry),
	).InsertOne(context.Background(), entry)
	return
}

func (r repository) Delete(entry dbfactory.IDbModel) (err error) {
	_, err = r.dataBase.Collection(
		r.getCollectionName(entry),
	).DeleteOne(context.Background(), bson.M{"_id": entry.GetID()})
	return
}

func (r repository) Update(entry dbfactory.IDbModel, args ...interface{}) (err error) {
	_, err = r.dataBase.Collection(
		r.getCollectionName(entry),
	).UpdateOne(context.Background(), bson.M{"_id": entry.GetID()}, entry)
	return
}

func (r repository) Query(entry dbfactory.IDbModel) dbfactory.IQuery {
	return &query{
		entry:    entry,
		dataBase: r.dataBase,
	}
}

func (r repository) getCollectionName(entry dbfactory.IDbModel) string {
	return reflect.TypeOf(entry).Name()
}
