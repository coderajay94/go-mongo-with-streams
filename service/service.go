package service

import (
	"context"
	"fmt"

	"github.com/coderajay94/go-mongo-with-streams/db"
	"github.com/coderajay94/go-mongo-with-streams/model"
)

type Service interface {
	SaveUserData(ctx context.Context, user model.UserAccountRequest)error
}
type baseService struct {
	db db.MongoDatabase
}

func NewService(db db.MongoDatabase) *baseService {
	return &baseService{
		db:db,
	}
}

func(s *baseService) SaveUserData(ctx context.Context,  user model.UserAccountRequest) error {
	fmt.Println("Saving data into mongo db")
	return s.db.SaveUserData(ctx,user)
}