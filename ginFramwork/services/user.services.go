package services

import (
	"github.com/Shrut26/ginFramework/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Userservices interface {
	CreateUser(*models.User) (*mongo.InsertOneResult, error)
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}
