# gorm-rest-api
Simple REST API using Go Lang + Gorm + Mux


Example Architecture + Structure

---

## Structure / Architecture

<pre>
.
+
|
+-+ cmd
|
+-+ config
|
+-+ db
|
+-+ http
|
+-+ router
|
+-+ [domain]
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
