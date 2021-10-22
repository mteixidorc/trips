# Trips

An example of a simple API, just a way to make a practice applying DDD, TDD, hexagonal architecture,  coding with Golang and applying SOLID principles.

The intention is to encapsulate and push the application logic to the center (domain) and use interfaces where it is possible in order to replace current data access and other input / output operative.
The way we obtain information from our domain (http rest services, console commands, and so on) should be independent from domain logic itself. 

As you see project is named Trips but it will contain the code for several microservices, each one representing a bounded context of our solution and our first bounded context is named trips as well ;-)

Thoughts about the future:
Right now we have the functionality to perform simple CRUD operations with trips but in the future we could have also a solution to work with means of transport, routes, a backoffice, who knows, and these concepts whould become several bounded context deployed as distinct microservice.


## Clone the project

First things first, clone the project:
```
git clone https://github.com/mteixidorc/trips.git
```

## Requirements

You will need GO compiler to build and/or execute the application 
[Golang](https://golang.org)


## Build / Execute 

In order to start `trip` service, go to `/apps/httpserver/servers/trip` folder and execute:

```
go build
```

And run it:
```
./trip
```

Server will run using port 8080 this parameter is configurable via environment variable, so use the port you want.

```
PORT=9090 ./trip
```

Alternatively you could execute project directly without a previous compilation

```
go run main.go 
```

One way to test endpoints is with curl:

POST Trip Creation 
```
curl -X POST http://localhost:8080/trip -H 'Content-Type: application/json' -d '{ "originId": 2, "destinationId": 1, "dates": "Sat Sun", "price": 40.55}'
```
GET All Trips
```
curl -X GET http://localhost:8080/trip 
```

GET a trip by ID 
```
curl http://localhost:8080/trip/6de052bc-94f9-4bfc-8fa5-0962ff9b4b99
```

### Tests

Tests are made at several levels:
- domain: Value Objects mainly
- application: Use cases, queries, updaters.
- apps: http server handlers


Run all tests just with this:
```
go test ./...
```

### Architecture and project structure ###

This project is an opportunity to improve DDD acknowledge, applying hexagonal architecture with the particularity of each language, in our case Golang.

And because this, files are organized using this structure:

```
.
├── apps
│   └── httpserver
│       ├── controllers
│       │   ├── shared
│       │   └── trip
│       └── servers
│           ├── shared
│           └── trip
├── data
└── internal
    ├── shared
    │   └── domain
    │       ├── errors
    │       └── value
    └── trips
        ├── application
        │   ├── city
        │   ├── mock
        │   └── trip
        ├── domain
        └── infrastructure
            └── repository
                └── mock
```

- apps: All applications consuming our domain should be put in this folder, the idea is to create an application per bounded context (aka microservice), in this example an API REST for the Tips context. Theres only one http service for all the solution and controllers for each context.
- data: Data files and other resources, currently we have `cities.txt` (cities data source).
- internal: Here we will put all code related to our domain (domain, application and infrastructure folders as usual). Each bounded context will have its own group of folders and communication between contexts should be done using DTOs to protect domain of the outside. Also we have a `shared` folder to put code used by any context.

### REST API Description

At the moment I've just implemented the Rest Service for this requests:

| Method | Endpoint  | Description          |
|--------|-----------|----------------------|
| GET    | /trip     | List all trips       |
| POST   | /trip     | Add a new trip       |
| GET    | /trip/:id | Get trip with ID :id |

The trips will be obtained in the following format:

```json
{
    "origin": "Barcelona",
    "destination": "Seville",
    "dates": "Mon Tue Wed Fri",
    "price": 40.55
}
```

Whereas to add a trip, we would send the following:

```json
{
    "originId": 2,
    "destinationId": 1,
    "dates": "Sat Sun",
    "price": 40.55
}
```

The list of cities is in a text file in folder `data` and is automatically loaded each time service starts. Maybe in the future we will use a DB instead.


### TODO

- Connect infrastructure to a real services.
- Implement an API versioning system 
