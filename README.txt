DBA 0.1 

# Index 

In this file you will find: 

How to read the README
Description for the application. 
Golang: version and how to use.
Docker: how to make an image and run it.

# How to read the README

In this file there are some example of code. The lines before the command orders, the text is annotated with "-".
The parts without the "-" annotation can be copied. In lines with "- OR" it shows an alternative way and its not part
of the command lines to be executed.  

# Description 

This project is a attempt for a web based application to act as a supplment for DBA 3.0. 
The goal is to access armies of the four different books, be able to see their composition, 
and to be able to see the statistics of the units.

The user can use an http call and request an army. 
Example: "localhost:8080/getarmy?armyid=1" 

The return should be a JSON with the composition of the army. 

The armies are currently saved in a SQLite3. 
There is are models of Unit and Army and their composition create different Armies. 
If the is no database in place, scripts are provided to initiate one on the start of the program. 

# Golang 

- This application uses 1.21.3 golang. To run the program use: 
go run main.go 

- OR 

go build dba 
go run .  
- OR 
go run ./build

- Use an http call to: 
localhost:8080/getarmy?armyid=1

- Example: 
curl "http://localhost:8080/getarmy?armyid=1"

For more about Golang:
https://go.dev/doc/install

For more about SQLite3:
https://pkg.go.dev/github.com/mattn/go-sqlite3

# Docker 

- to create the docker image: 
docker build --pull --rm -f "Dockerfile" -t dba:latest "."

- to run the docker image: 
docker run -p 8080:8080 dba:latest

For more about Docker:
https://docs.docker.com/engine/install/