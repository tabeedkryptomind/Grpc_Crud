package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	db "Grpc-crud/app/database"

)

var (
	r = gin.Default()
)

const (
	port = ":8080"
)

func StartApp() {
	Url_maps()
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()
	r.Run(port)
}
