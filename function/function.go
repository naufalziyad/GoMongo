package function

import (
	"fmt"
	"context"
	"log"
	"models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"

)


var db *mongo.Database

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(CONNECTION))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database(DBNAME)
}


func InsertOne(user models.User) {
	fmt.Println(user)
	_,err := db.Collection (COLLECTIONNAME).InsertOne(context.Background(), user)
	if err !=nil {
		log.Fatal(err)
	}
}

func InsertMany(person []models.User) {
	var dataPerson []interface{}
	for _, pers := range person {
		dataPerson = append(dataPerson, pers)
	}
	_, err := db.Collection(COLLECTIONNAME).InsertMany(context.Background(), dataPerson)
	if err != nil{
		log.Fatal(err)
	}
}

func GetAllUser() []models.User {
	personal, err := db.Collection(COLLECTIONNAME).Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var arrayPersonal []models.User
	var idxPersonal models.User

	for personal.Next(context.Background()) {
		err := personal.Decode(&idxPersonal)
		if err !=nil {
			log.Fatal(err)
		}
		arrayPersonal = append(arrayPersonal, idxPersonal)
	}

	if err := personal.Err(); err !=nil {
		log.Fatal(err)
	}

	personal.Close(context.Background())
	return arrayPersonal
}
