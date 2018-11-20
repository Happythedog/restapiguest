package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

var config = Config{}
var dao = GuestsDAO{}

// GET list of guests
func AllGuestsEndPoint(w http.ResponseWriter, r *http.Request) {
	guests, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, guests)
}

// GET a guest by its ID
func FindGuestEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	guest, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, guest)
}

// POST a new guest
func CreateGuestEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var guest Guest
	if err := json.NewDecoder(r.Body).Decode(&guest); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	guest.ID = bson.NewObjectId()
	if err := dao.Insert(guest); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, guest)
}

// PUT update an existing guest
func UpdateGuestEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var guest Guest
	if err := json.NewDecoder(r.Body).Decode(&guest); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(guest); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing guest
func DeleteGuestEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var guest Guest
	if err := json.NewDecoder(r.Body).Decode(&guest); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(guest); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/guests", AllGuestsEndPoint).Methods("GET")
	r.HandleFunc("/guests", CreateGuestEndPoint).Methods("POST")
	r.HandleFunc("/guests", UpdateGuestEndPoint).Methods("PUT")
	r.HandleFunc("/guests", DeleteGuestEndPoint).Methods("DELETE")
	r.HandleFunc("/guests/{id}", FindGuestEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
