package database

type CRUD interface {
	InsertOne(document interface{}) (int, error)
	InsertMany(documents []interface{}) (int, error)
	GetOne(id string) (interface{}, error)
	GetMany(ids []string) ([]interface{}, error)
	UpdateOne(id string, document interface{}) (int, error)
	UpdateMany(ids []string, documents []interface{}) (int, error)
	DeleteOne(id string) (int, error)
	DeleteMany(ids []string) (int, error)
}
