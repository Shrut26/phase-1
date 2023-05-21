package services

import (
	"context"
	"errors"

	"github.com/Shrut26/ginFramework/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) Userservices {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) (*mongo.InsertOneResult, error) {
	result, err := u.usercollection.InsertOne(u.ctx, user)
	return result, err
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	var user *models.User
	filter := bson.D{bson.E{Key: "user_name", Value: name}}
	err := u.usercollection.FindOne(u.ctx, filter).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(u.ctx) {
		var user models.User
		err := cursor.Decode(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}
	return users, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "user_name", Value: user.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "user_name", Value: user.Name}, bson.E{Key: "user_age", Value: user.Age}, bson.E{Key: "user_address", Value: user.Address}}}}
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched user found")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	filter := bson.D{bson.E{Key: "user_name", Value: name}}
	result, _ := u.usercollection.DeleteOne(u.ctx, filter)

	if result.DeletedCount != 1 {
		return errors.New("no matched found user")
	}
	return nil
}
