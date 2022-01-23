package dbfactory

type IQuery interface {
	Where(...interface{}) IQuery
	ToArray(res interface{}) error
	Count() (int64, error)
}
