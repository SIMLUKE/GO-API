package main

import (
	"airbd/api/routes/router"
	"airbd/api/server"
	"fmt"
	"strconv"
)

func main() {
	serv := server.OpenServer()
	//userhandling.Newuser(&db, "luc", "luc@gmail.com", "lise")
	//articleshandling.Newarticle(&db, "AMOGUS SHIRT", "M", 99.99, "New", "Balenciaga", "Red", "68b36685-15c2-4555-b58c-4f7448fd4423")
	fmt.Println(serv)

	router.ApplyRoutes(serv.Serv)

	serv.Serv.Run(fmt.Sprintf(":%s", strconv.Itoa(serv.Port)))
}
