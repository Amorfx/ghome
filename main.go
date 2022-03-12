package main

import (
	"clementdecou/ghome/device"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Env struct {
	db *mongo.Client
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Init Mongo db connection
	env := Env{db: InitMongoDbConnection()}

	router := gin.Default()
	router.LoadHTMLFiles("index.html")
	dm := device.InitDeviceManager()

	router.GET("/", func(c *gin.Context) {
		allDevices := device.GetAllDevices(env.db)
		c.HTML(http.StatusOK, "index.html", gin.H{"titre": "Mon titre", "allDevices": allDevices, "deviceTypes": dm.Types})
	})

	router.POST("/devices/add", func(c *gin.Context) {
		if len(c.PostForm("device_ip")) > 0 {
			device.AddDevice(env.db, c.PostForm("device_type"), c.PostForm("device_name"), c.PostForm("device_ip"))
			c.Redirect(http.StatusFound, "/")
		}
	})

	router.Run()
}

func InitMongoDbConnection() *mongo.Client {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return client
}
