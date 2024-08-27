package service

import (
	"github.com/coderajay94/go-mongo-with-streams/db"
	"github.com/coderajay94/go-mongo-with-streams/model"
)

type Service interface {
	SaveUserData(model.UserAccountRequest) error
}
type baseService struct {
	db db.MongoDatabase
}
