package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Shrut26/ginFramework/controllers"
	"github.com/Shrut26/ginFramework/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userservice    services.Userservices
	usercontroller controllers.UserController
	ctx            context.Context
	usercollection *mongo.Collection
	mongoClient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()
	mongoConn := options.Client().ApplyURI("mongodb+srv://shrutayuaggarwal26:gin_go_Api@cluster0.w1hh9a1.mongodb.net/?retryWrites=true&w=majority")
	mongoClient, err = mongo.Connect(ctx, mongoConn)

	if err != nil {
		log.Fatal(err)
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo Connected")
	usercollection := mongoClient.Database("userdb").Collection("users")
	userservice = services.NewUserService(usercollection, ctx)
	usercontroller = controllers.New(&userservice)
	server = gin.Default()
}
func main() {
	defer mongoClient.Disconnect(ctx)
	basepath := server.Group("/vi")
	usercontroller.RegisterUserRoutes(basepath)
	log.Fatal(server.Run(":9090"))
}
