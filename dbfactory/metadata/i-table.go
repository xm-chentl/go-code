package metadata

type ITable interface {
	GetValueMap(value interface{}) map[string]interface{}
}
