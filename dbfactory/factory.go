package dbfactory

import (
	"fmt"

	"github.com/xm-chentl/go-code/dbfactory/dbtype"
)

type factory struct {
	dbTypeMap map[dbtype.Value]IDbFactory
}

func (f *factory) Db(args ...interface{}) IRepository {
	dbType := dbtype.MySql
	if len(args) > 0 {
		dbType = args[0].(dbtype.Value)
	}
	if _, ok := f.dbTypeMap[dbType]; !ok {
		panic(fmt.Errorf("[%s]类型dbsvc未注入", dbType.String()))
	}
	// todo: 此外工作单元传入的话，会报错，要处理下哟
	return f.dbTypeMap[dbType].Db(args...)
}

func (f factory) Uow() IUnitOfWork {
	return nil
}

func New(dbMap map[dbtype.Value]IDbFactory) IDbFactory {
	return &factory{
		dbTypeMap: dbMap,
	}
}
