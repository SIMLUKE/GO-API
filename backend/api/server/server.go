package server

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Server struct {
	Host string
	Port int
	Serv *gin.Engine
}

func OpenServer() Server {
	var serv Server
	var err error
	godotenv.Load()

	serv.Host = os.Getenv("SERVER_HOST")
	serv.Port, err = strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		panic(err)
	}
	serv.Serv = gin.Default()
	return serv
}
