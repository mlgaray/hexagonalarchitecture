package mongo

import (
	"context"
	models2 "github.com/mlgaray/hexagonalarchitecture/internal/core/models"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	client   *mongo.Client
	users    []models2.User
	Incident []models2.Incident
}

func (s *UserRepository) GetAll() ([]models2.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	u := models2.User{
		Name:  "lorenzo",
		Email: "garay",
	}
	s.users = append(s.users, u)
	return s.users, nil

	db := s.client.Database("notfinancialadvice")
	col := db.Collection("incidents")

	// u.Password,_ = EncryptPassword(u.Password)
	// result, err := col.Find(ctx, bson.D{{}})

	/*	match := bson.D{
		{"category", "seguridad"},
	}*/
	pipeline := mongo.Pipeline{
		/*	{{"$match", match}},*/
		{{"$lookup", bson.M{
			"from":         "users",
			"localField":   "userId",
			"foreignField": "_id",
			"as":           "user",
		}}},
	}
	result1, err := col.Aggregate(ctx, pipeline)
	spew.Sprint(result1)
	if err != nil {
		return nil, err
	}

	/*	for result.Next(ctx) {
			//Create a value into which the single document can be decoded
			var elem models.User
			err := result.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}

			s.users = append(s.users, elem)

		}
		if err := result.Err(); err != nil {
			log.Fatal(err)
		}
		//Close the cursor once finished
		result.Close(context.TODO())*/

	for result1.Next(ctx) {
		// Create a value into which the single document can be decoded
		var elem models2.Incident
		err := result1.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		// s.Incident = elem
		s.Incident = append(s.Incident, elem)
	}
	if err := result1.Err(); err != nil {
		log.Fatal(err)
	}
	// Close the cursor once finished
	result1.Close(context.TODO())

	// ObjID,_ := result.InsertedID.(primitive.ObjectID)
	// return ObjID.String(), true, nil

	return s.users, nil
}

func (s *UserRepository) SignUp(user models2.User) error {
	return nil
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{
		client: client,
	}
}
