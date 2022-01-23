package dbfactory

type IDbFactory interface {
	Uow() IUnitOfWork
	Db(...interface{}) IRepository
}
