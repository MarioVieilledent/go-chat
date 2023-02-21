package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var url string = "0.0.0.0"

// Setup du port dans env pour Heroku
var port string = getPort()

var messages []string = []string{}

func main() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.Use(static.Serve("/", static.LocalFile("client/dist", false)))

	router.GET("/mes", getAllMessages)
	router.PUT("/mes", sendMessage)

	fmt.Println("API on " + url + ":" + port)
	router.Run(url + ":" + port)
}

func getAllMessages(c *gin.Context) {
	data, err := json.Marshal(messages)

	if err != nil {
		c.JSON(400, string(data))
		return
	}

	c.JSON(200, string(data))
}

func sendMessage(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, "Error read JSON")
		return
	}
	var elem string
	err = json.Unmarshal([]byte(jsonData), &elem)
	if err != nil {
		c.JSON(400, "Error Unmarshal JSON")
		return
	}
	messages = append(messages, elem)
	c.JSON(200, "OK")
}

// Get either port of env else 3000
func getPort() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return
}
