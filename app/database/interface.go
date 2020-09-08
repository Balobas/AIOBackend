package database

/*
Интерфейс базы данных.
Любую структуру, наследующую эти методы, можно использовать в коде как базу данных
 */

type Database interface {
	Init(login, password, url string) error
	Set (key string, obj interface{}) error
	Get (key string) (map[string]interface{}, error)
	Delete (key string) error
    Unmarshal(obj interface{}, objMap map[string]interface{}) (interface{}, error)
    QueryAllFieldsWithSelector(selector string) ([]map[string] interface{}, error)
}