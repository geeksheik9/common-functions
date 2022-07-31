package crud

type CRUD interface {
	InsertOne(document interface{}) int
	InsertMany(documents []interface{}) int
	GetOne(id string) interface{}
	GetMany(ids []string) []interface{}
	UpdateOne(id string, document interface{}) int
	UpdateMany(ids []string, documents []interface{}) int
	DeleteOne(id string) int
	DeleteMany(ids []string) int
}
