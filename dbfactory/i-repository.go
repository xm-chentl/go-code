package dbfactory

type IRepository interface {
	Create(IDbModel) error
	Delete(IDbModel) error
	Update(IDbModel, ...interface{}) error
	Query(IDbModel) IQuery
	Close()
}
