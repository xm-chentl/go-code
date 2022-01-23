package mongoex

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type unitOfWork struct {
	session mongo.Session
}

func (u unitOfWork) Commit() (err error) {
	if u.session != nil {
		err = u.session.CommitTransaction(context.Background())
	}
	if err != nil {
		u.session.AbortTransaction(context.Background())
	}

	return
}
