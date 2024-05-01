package postgresql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DataBaseConnection interface {
	Connect() *sql.DB
}

type dataBaseConnection struct{}

func (c *dataBaseConnection) Connect() *sql.DB {
	// Configura la cadena de conexión
	/*dbEndpoint := "incident-management-db-instance.cwc5c9nusj4r.us-east-2.rds.amazonaws.com"
	dbPort := "3306"
	dbUser := "admin"
	dbPassword := "As4551603"
	dbName := "incident_management_db"*/

	//dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbEndpoint, dbPort, dbName)

	db, err := sql.Open("postgres", "user=postgres.rzvwzzmdylhgpbcymjjw password=As4551603!? host=aws-0-us-west-1.pooler.supabase.com port=5432 dbname=postgres")
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	// Verifica la conexión
	// err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexión exitosa a la base de datos!")

	return db
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
	//return nil
}

func NewDataBaseConnection() *dataBaseConnection {
	return &dataBaseConnection{}
}
