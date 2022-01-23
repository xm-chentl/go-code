package mongoex

import (
	"context"

	"github.com/xm-chentl/go-code/dbfactory"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type factory struct {
	client   *mongo.Client
	dataBase *mongo.Database

	dbName string
	uri    string
}

func (f *factory) Db(args ...interface{}) dbfactory.IRepository {
	return &repository{
		dataBase: f.dataBase,
	}
}

func (f factory) Uow() dbfactory.IUnitOfWork {
	session, err := f.client.StartSession()
	if err != nil {
		panic(err)
	}

	return &unitOfWork{
		session: session,
	}
}

func New(dbName, uri string) dbfactory.IDbFactory {
	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(uri),
	)
	if err != nil {
		panic(err)
	}

	return &factory{
		client:   client,
		dataBase: client.Database(dbName),
		dbName:   dbName,
		uri:      uri,
	}
}
