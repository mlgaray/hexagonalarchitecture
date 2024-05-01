package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type DataBaseConnection interface {
	Connect() *mongo.Client
}

type dataBaseConnection struct{}

func (c *dataBaseConnection) Connect() *mongo.Client {
	/*	var clientOptions = options.Client().ApplyURI("mongodb+srv://mlgaray:As455160352907891@cluster0.9kxvl.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal(err.Error())
			return client
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err.Error())
			return client
		}
		log.Println("conexion exitosa con la BD")
		return client*/
	return nil
}

func NewDataBaseConnection() *dataBaseConnection {
	return &dataBaseConnection{}
}
