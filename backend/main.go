package main

import (
	"airbd/api/routes/router"
	"airbd/api/server"
	"fmt"
	"strconv"
)

func main() {
	serv := server.OpenServer()

	router.ApplyRoutes(serv.Serv)

	serv.Serv.Run(fmt.Sprintf(":%s", strconv.Itoa(serv.Port)))
}
