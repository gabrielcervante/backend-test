# Japhy Backend Test in Golang

## Technical Stack
- Go
- Docker
- MySQL

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Docker](https://www.docker.com/products/docker-desktop/)
- [Git](https://git-scm.com/downloads)

## Tasks
you are a backend developer in a pet food startup. A new functionality will be implemented, we want to be able to easily manage the breeds of dogs and cats that are registered in the database in our back office,
you must implement a CRUD api to manage the breeds of pets. The breeds are stored in a CSV file located at `./breeds.csv`. 
the aim of this test is to demonstrate backend development skills using the Go programming language. The application implements a simple REST API for managing resources
you are free to take initiatives and make improvements to the codebase.
Have fun and good luck!

### you need to implement the following tasks:
- create a new table in the core database to store the breeds of pets, to do this you must create a new migration file in the `database_actions` directory.
- store the breeds of pets in the database (list of breeds are on `breeds.csv`).
- implement CRUD functionality for the breeds resource (GET, POST, PUT, DELETE).
- search functionality to filter breeds based on pet characteristics (weight and species).


## Installation

1. Fork the project repository
2. Copy the `.env.example` file to `.env`
3. Build the application `docker compose build`
4. Run docker compose to start the application `docker compose up -d`
5. Once the application is up and running, you can access the REST API at http://localhost:50010. Use tools like Postman or curl to interact with the API.
6. `curl -v http://localhost:50010/health` to ensure your application is running.
7. send us the link to your repository with the api.

## Observations

- All the endpoints are in /v1/
- To use Get we need to provide a query parameter as: /v1/?name=nameofyourpet or /v1/?id=yourbreedsid or /v1/?species=yourpetspecies or /v1/?pet_size=yourpetsize or /v1/?weight?=yourpetweight
- To use Delete we need to provide the breeds id as: /v1/?id=yourbreedsid
- To use the Post and Put you will need to provide the same json which is: 
{
	"id": 700,
	"species": "cat",
	"pet_size": "medium",
	"name": "Raj",
	"weight": 20000,
	"average_male_adult_weight": 12000,
	"average_female_adult_weight": 12000
}
- When using Put just change what you need to do and send the other data unchanged or it will be with the go zero value in the database.

- The last thing is the declaration of platform: linux/arm64 in the mysql-test service in the docker-compose file was removed because i don't use arm and it was crashing everything here so if you have arm just put it again.