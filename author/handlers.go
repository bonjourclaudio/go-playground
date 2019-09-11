package author

import (
	"encoding/json"
	"github.com/claudioontheweb/gorm-rest-api/db"
	customHTTP "github.com/claudioontheweb/gorm-rest-api/http"
	"github.com/gorilla/mux"
	"net/http"
)

// Get One Author By ID
func GetAuthorsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var authors []Author
	db.DB.Find(&authors)
	json.NewEncoder(w).Encode(authors)
}

// Get all Authors
func GetAuthorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var author Author
	db.DB.First(&author, params["authorId"])
	json.NewEncoder(w).Encode(author)
}

// Create Author
func CreateAuthorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var author Author
	json.NewDecoder(r.Body).Decode(&author)
	err := db.DB.Create(&author).Error
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
		return
	}
	json.NewEncoder(w).Encode(&author)
}

// Delete Author
func DeleteAuthorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var author Author
	db.DB.Delete(&author, params["authorId"])
	w.Write([]byte("Successfully deleted Author with ID: " + params["authorId"]))
}