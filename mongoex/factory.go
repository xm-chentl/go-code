package mongoex

import (
	"context"
	"fmt"
	"os"

	"github.com/xm-chentl/go-code/dbfactory"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
	ctx := context.Background()
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(uri),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to mongodb: %v\n", err)
		// os.Exit(1)
		panic(err)
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	return &factory{
		client:   client,
		dataBase: client.Database(dbName),
		dbName:   dbName,
		uri:      uri,
	}
}
