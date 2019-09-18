package main

import (
	"github.com/claudioontheweb/go-playground/author"
	"github.com/claudioontheweb/go-playground/config"
	"github.com/claudioontheweb/go-playground/db"
	customRouter "github.com/claudioontheweb/go-playground/router"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func NewRouter() *mux.Router {

	// Init Router
	r := mux.NewRouter()

	// Append author routes
	customRouter.AppRoutes = append(customRouter.AppRoutes, author.Routes)

	for _, route := range customRouter.AppRoutes {

		// Create subroute
		routePrefix := r.PathPrefix(route.Prefix).Subrouter()

		// Loop through each subroute

		for _, r := range route.SubRoutes {

			var handler http.Handler
			handler = r.HandlerFunc

			routePrefix.Path(r.Pattern).Handler(handler).Methods(r.Method).Name(r.Name)
		}
	}

return r

}

func main() {

	config.GetConfig()

	port := viper.GetString("PORT")
	router := NewRouter()

	// Setup DB
	db.DB = db.ConnectDB()

	defer db.DB.Close()

	// Create table
	if !db.DB.HasTable("authors") {
		db.DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&author.Author{})
	}

	// HTTP Server
	log.Fatal(http.ListenAndServe(":" + port, router))

}
