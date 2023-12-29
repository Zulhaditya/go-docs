package simple

type Database struct {
	Name string
}

type DatabasePostgreSQL Database
type DatabaseMongoDB Database

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	database := &Database{
		Name: "PostgreSQL",
	}
	return (*DatabasePostgreSQL)(database)
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	database := &Database{
		Name: "MongoDB",
	}
	return (*DatabaseMongoDB)(database)
}

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMongoDB    *DatabaseMongoDB
}

func NewDatabaseRepository(postgreSQL *DatabasePostgreSQL, mongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgreSQL: postgreSQL, DatabaseMongoDB: mongoDB}
}
