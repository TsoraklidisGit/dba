DBA 0.1 

This project is a attempt for a web based application to act as a supplment for DBA 3.0. 
The goal is to access armies of the four different books, be able to see their composition, 
and to be able to see the statistics of the units.

The user can use an http call and request an army. 
Example: "localhost:8080/getArmy?armyID=1" 

The return should be a JSON with the composition of the army. 

The armies are currently saved in a SQLite3. 
There is are models of Unit and Army and their composition create different Armies. 
If the is no database in place, scripts are provided to initiate one on the start of the program. 

# Docker 
to create the docker image: 

docker build --pull --rm -f "Dockerfile" -t dba:latest "."

to run the docker image: 
docker run -p 8080:8080 dba:latest