package dbtype

type Value string

func (v Value) String() string {
	return string(v)
}

const (
	MySql Value = "mysql"
	Mongo Value = "mongo"
)
