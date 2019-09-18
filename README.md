# Go Playground
Simple REST API + DB (Docker)


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
4. Server listens on [localhost:8080](http://localhost:8080) (Specify port in config/config.json)

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

## Useful Go Articles and Ressources

### Go Basics
* [Learn go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI)
* [Go Crash Course](https://www.youtube.com/watch?v=SqrbIlUwR0U)
* [Go Tour](https://tour.golang.org/welcome/1)

### Architecture
* **[Clean Architecture](https://medium.com/@eminetto/clean-architecture-using-golang-b63587aa5e3f)**
* [Clean Architecture 2](https://hackernoon.com/golang-clean-archithecture-efd6d7c43047)
* [Clean Architecture 3](https://geeks.uniplaces.com/putting-clean-architecture-into-practice-20c47d8c76de)
* [Clean Architecture 4](https://medium.com/@matthieujacquot/clean-architecture-in-golang-7ebafe4c70db)
* [Web Services Architecture](https://boobo94.xyz/web-service/webservice-architecture-golang/)
* [Go and Package Focused Design Article](https://blog.gopheracademy.com/advent-2016/go-and-package-focused-design/)
* [Structuring Go Apps](https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091#.kj6eym1u4)

### Go + DB
* [MongoDB + Go Best Practices](https://www.nicolasmerouze.com/how-to-render-json-api-golang-mongodb/)

### Concurrency
* [Concurrency Patterns](https://www.youtube.com/watch?v=YEKjSzIwAdA)

---

## Dependencies
-> [Repo](https://godoc.org/-/subrepo)
### HTTP / REST API / Web
* [Mux](https://github.com/gorilla/mux)
  * [REST API using Mux](https://www.youtube.com/watch?v=SonwZ6MF5BE)
* [Echo](https://github.com/labstack/echo)

### DB
* [Gorm (ORM Library)](https://github.com/jinzhu/gorm)
* [Goose (DB Migration tool)](https://github.com/steinbacher/goose)

### Dpendency Managers
* [Dep](https://github.com/golang/dep)
* [Glide](https://github.com/Masterminds/glide)

### Logging / Error
* [Logrus](https://github.com/sirupsen/logrus)
* [Errors](https://github.com/juju/errors)

### Testing
* [Testify](https://github.com/stretchr/testify)

### Others
* [Colors](https://github.com/fatih/color)
* [Now (Time toolkit)](https://github.com/jinzhu/now)

---

## Frameworks
* [Toolkit for microservices (A la Spring Boot)](https://gokit.io/examples/)
* [Beego](https://beego.me/)

---

## Cheat Sheets
* [Cheat Sheet](https://devhints.io/go)

