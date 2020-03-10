package main

import (
	"encoding/json"
	"net/http"
	"log"

	"github.com/gorilla/mux"
)
const api_url = "https://api.exchangeratesapi.io/latest"
var mongoStore = MongoStore{}

func main(){
	log.Println("Initialize mongo db connection")
	session := initialiseMongo()
	mongoStore.session = session

	log.Println("Start routing setup")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":9090", router))
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	keys, ok := r.URL.Query()["base"]
	url := api_url

    if ok && len(keys[0]) > 0 {
		log.Println(">>> Base Param: " + keys[0])
		url = url + "?base=" + keys[0]
    }

	log.Println("Api request URL: " + url)

	response := http_get(url)

	// Log Response Base currency And Date
	log.Println(response.Base)
	log.Println(response.Date)
	log.Println(response.Rates)

	exchangeRate := ExchangeRate{}
	exchangeRate.BaseCurrency = response.Base
	exchangeRate.Date = response.Date

	// insert some meta data into mongo
	insert(&mongoStore, &exchangeRate)

    log.Println("Set content-type http header")  
    w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

    log.Println("Send response")
    json.NewEncoder(w).Encode(response)
}

