# gorm-rest-api
Simple REST API using Go Lang + Gorm + Mux


Example Architecture + Structure

## Requirements

- [Go](https://golang.org/)
- [Dep](https://github.com/golang/dep)
- [Docker](https://www.docker.com/)
- (Windows only: [make](http://gnuwin32.sourceforge.net/packages/make.htm)) 

---

## Run the proj

(Run the following commands from the project root folder)

1. Start DB (Docker): `docker-compose up`
2. Get Project's Dependencies: `dep ensure`
3. Run Application: `go run cmd/main.go`

## Build the proj

1. Test & Build: `make`
2. Only build: `make build`
3. Build & Run: `make run`

---

## Structure / Architecture

<pre>
.
+
|
+-+ cmd/
|
+-+ config/
|
+-+ db/
|
+-+ http/
|
+-+ router/
|
+-+ [domain]/
|
+
...

</pre>

### Cmd/
This folder contains the package 'main'. It is the entry point and sets up the Router, DB and Server.

### Config/
The configuration file can be found here. There is also a config.go that returns the configuration.

### Db/
Used to connect to the DB. It get's called by the main.go.

### Http/
Middleware and reponse structs can be found here.

### Router/
A route struct is stored here. It can be used to create routes for a domain.

### [domain]/
Every domain should have it's own folder containing the following structure:

- model.go
- routes.go
- handlers.go

---

## Ressources
[Gorm Doc](https://gorm.io/docs/index.html)
