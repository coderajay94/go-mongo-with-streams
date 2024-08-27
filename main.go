package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/coderajay94/go-mongo-with-streams/db"
	"github.com/coderajay94/go-mongo-with-streams/model"
	"github.com/coderajay94/go-mongo-with-streams/service"
	"github.com/pborman/uuid"
)

func main() {

	//fmt.Println("starting the golang mongo with streams...")

	log.Fatal( http.ListenAndServe(":8080", nil))
	fmt.Println("http service listening on port:8080")
	//var dbService db.MongoDatabase
	dbService, error := db.NewClient("mongodb://localhost:27017", "admin", "admin")

	defer dbService.Close(context.Background())

	if error != nil{
		fmt.Println("error occured while connecting to MongoDB",error)
	}
	fmt.Println("starting the golang mongo server...")
	var s service.Service
	s = service.NewService(dbService)


	user := model.UserAccountRequest{
		//UserId: "UserId123",
		Name: "ajay Kumar",
		Email:"ajay@test.com",
		Password: "pwd@123",

	}

	user.UserId = uuid.NewUUID().String()
	err := s.SaveUserData(context.Background(), user)

	if err != nil{
		fmt.Println("Failed to save user data")
	}

	fmt.Println("User was inserted successfully into mongodb...")

}