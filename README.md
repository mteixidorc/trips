# Trips

An example of a simple API, just a way to make a practice applying DDD and coding with Golang.
The intention is to encapsulate and push the application logic to the center (domain) and use interfaces where it is possible in order to replace current data access and other input / output operative.
The way we obtain information from our domain (http rest services, console commands, and so on) should be independent from domain logic itself. 

## Clone the project

First things first, clone the project:
```
git clone https://github.com/mteixidorc/trips.git
```

## Requirements

You will need GO compiler to build and/or execute the application 
[Golang](https://golang.org)


## Build / Execute 

In order to start service, go to /apps/httpserver folder and execute:
```
go build
```

And run it:
```
./httpserver
```

Server will run using port 8080, right now this parameter is not configurable.

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

Run all tests just with this:

```
go test ./...
```

Tests are made at distinct levels:
- domain: Value Objects
- application: Use cases, queries, updaters.
- apps: http server handlers

### Architecture ###

This project is an opportunity to improve DDD acknowledge, applying hexagonal architecture with the particularity of each language, in this case Golang.

Files are organized using this structure:

- apps: All applications that use our domain will be implemented in this folder. The idea is to create on application per bounded context (aka microservice), in this example an API REST for the Tips context.
- data: Data files and other resources.
- internal: Here we will put all code related to our domain (domain, application and infrastructure).

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
- Make a API versioning system 
