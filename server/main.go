package main

import (
	"log"
	"net/http"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Film data structure
type Films struct {
	ID			int
	Name		string
	Categorie	string
}

// Setting up Database instance
var Instance *gorm.DB

// Helper function for crud operations
func checkIfFilmExist(filmId string) bool{
	var films Films
	Instance.First(&films, filmId)
	if films.ID == 0 {
		return false
	}
	return true
}


// Update film in database
func updateFilm(w http.ResponseWriter, r *http.Request) {
	filmId := mux.Vars(r)["id"]
	if checkIfFilmExist(filmId) == false {
		json.NewEncoder(w).Encode("Product not found!")
		return
	}
	var films Films
	Instance.First(&films, filmId)
	json.NewDecoder(r.Body).Decode(&films)
	Instance.Save(&films)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(films)
}


func getFilms(w http.ResponseWriter, r *http.Request) {
	var films []Films

	// Map all the avaible films into films variable
	Instance.Find(&films)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(films);
}

func getFilmById(w http.ResponseWriter, r *http.Request) {
	// Retrive data from query
	filmId := mux.Vars(r)["id"]
	if checkIfFilmExist(filmId) == false {
		json.NewEncoder(w).Encode("Film not found!")
		return
	}

	var films Films
	Instance.First(&films, filmId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(films);
}

// Handle the creating of a new film into Mysql
func setFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var films Films
	json.NewDecoder(r.Body).Decode(&films)
	Instance.Create(&films)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(films)
}

func deleteFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	filmId := mux.Vars(r)["id"]

	if checkIfFilmExist(filmId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product not found")
		return
	}
	var films Films
	Instance.Delete(&films, filmId)
	json.NewEncoder(w).Encode("Film deleted successfully!")
	w.WriteHeader(http.StatusOK)
}

// Setting up route for crud operations
func RegisterFilmsRoute(router *mux.Router) {
	router.HandleFunc("api/films", getFilms).Methods("GET")
	router.HandleFunc("api/films/{id}", getFilmById).Methods("GET")
	router.HandleFunc("api/films/{id}", setFilm).Methods("POST")
	router.HandleFunc("api/films/{id", deleteFilm).Methods("DELETE")
}

// Database connection
func migrate(){
	Instance.AutoMigrate(&Films{})
	log.Println("Database Migration Completed...")
}

func main() {

	//Setting up Mysql databas
	database, err := gorm.Open(mysql.Open("root:passwordisgood@tcp(127.0.0.1:3306)/films?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to open connection to database")
	}
	log.Println("Connected to Database", database)

	// Initialize router
	router := mux.NewRouter().StrictSlash(true)

	migrate()
	
	// Registering routes
	RegisterFilmsRoute(router)

	// Start server
	log.Fatal(http.ListenAndServe(":8000", router))
}

