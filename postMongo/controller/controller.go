package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Shrut26/Postmongo/config"
	"github.com/Shrut26/Postmongo/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://postmongo:postmongo@cluster0.7cvdzl9.mongodb.net/?retryWrites=true&w=majority"
const dbName = "postmongo"
const colName = "mongo"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo Connected")
	collection = (*mongo.Collection)(client.Database(dbName).Collection(colName))

	fmt.Println("Collection ready")
}

func UserControllers(ctx *gin.Context) {
	users := models.User{}
	config.DB.Find(&users)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"users":   users,
	})
}
