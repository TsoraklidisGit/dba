package main

import (
	"dba/controller"
	"dba/repository"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Serve static assets from the "frontend" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//Calling the repository to initiliaze and populate the Database
	repository.Repository()

	//REST API GET on a specific Army. example /getArmy?armyID=1
	http.HandleFunc("/getarmy", controller.GetArmy)

	// Serve the HTML page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	//Port for the calls
	port := ":8080"
	fmt.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
