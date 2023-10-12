package main

import (
	"dba/controller"
	"dba/repository"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Calling the repository to initiliaze and populate the Database
	repository.Repository()

	//REST API GET on a specific Army. example /getArmy?armyID=1
	http.HandleFunc("/getarmy", controller.GetArmy)

	//Port for the calls
	port := ":8080"
	fmt.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
