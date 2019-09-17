package main

import (
	"github.com/claudioontheweb/gorm-rest-api/author"
	"github.com/claudioontheweb/gorm-rest-api/db"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
	customRouter "github.com/claudioontheweb/gorm-rest-api/router"
	"log"
	"github.com/claudioontheweb/gorm-rest-api/config"
)

func NewRouter() *mux.Router {

	// Init Router
	router := mux.NewRouter()

	// Append author routes
	customRouter.AppRoutes = append(customRouter.AppRoutes, author.Routes)

	for _, route := range customRouter.AppRoutes {

		// Create subroute
		routePrefix := router.PathPrefix(route.Prefix).Subrouter()

		// Loop through each subroute

		for _, r := range route.SubRoutes {

			var handler http.Handler
			handler = r.HandlerFunc

			routePrefix.Path(r.Pattern).Handler(handler).Methods(r.Method).Name(r.Name)
		}
	}

return router

}

func main() {

	config.GetConfig()

	port := viper.GetString("PORT")
	router := NewRouter()

	// Setup DB
	db.DB = db.ConnectDB()

	if !db.DB.HasTable("authors") {
		db.DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&author.Author{})
	}


	defer db.DB.Close()

	// HTTP Server
	log.Fatal(http.ListenAndServe(":" + port, router))

}
