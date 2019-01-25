package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ImageAdd struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImgURL      string `json:"img_url"`
	Dimension   string `json:"dimension"`
	Recommended string `json:"recommended"`
}

var db *gorm.DB
var err error

func main() {
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	dbURL := os.Getenv("DB_URL")
	db, err = gorm.Open(
		"postgres", dbURL)
	// "postgres", "host=localhost user=bsinkule dbname=personal-chef sslmode=disable")

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&ImageAdd{})

	router.HandleFunc("/images/", GetImages).Methods("GET")
	router.HandleFunc("/images/{id}", GetImage).Methods("GET")
	router.HandleFunc("/images/", AddImage).Methods("POST")
	router.HandleFunc("/images/{id}", DeleteImage).Methods("DELETE")
	router.HandleFunc("/images/{id}", UpdateImage).Methods("PUT")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(headers, methods, origins)(router)))
	// log.Fatal(http.ListenAndServe(":5000", handlers.CORS(headers, methods, origins)(router)))
}

func GetImages(w http.ResponseWriter, r *http.Request) {
	var images []ImageAdd
	db.Order("id desc").Find(&images)
	json.NewEncoder(w).Encode(&images)
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var image ImageAdd
	db.First(&image, params["id"])
	json.NewEncoder(w).Encode(&image)
}

func UpdateImage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var image ImageAdd
	db.First(&image, params["id"])
	json.NewDecoder(r.Body).Decode(&image)
	db.Save(&image)
	json.NewEncoder(w).Encode(&image)
}

func AddImage(w http.ResponseWriter, r *http.Request) {
	var image ImageAdd
	json.NewDecoder(r.Body).Decode(&image)
	db.Create(&image)
	json.NewEncoder(w).Encode(&image)
}

func DeleteImage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var image ImageAdd
	db.First(&image, params["id"])
	db.Delete(&image)

	var images []ImageAdd
	db.Find(&images)
	json.NewEncoder(w).Encode(&images)
}
