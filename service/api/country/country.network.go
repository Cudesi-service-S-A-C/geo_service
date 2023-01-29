package country

import (
	"encoding/json"
	"net/http"

	"github.com/Cudesi-service-S-A-C/geo_service/service/connection"
	"github.com/Cudesi-service-S-A-C/geo_service/service/models"
	"github.com/gorilla/mux"
)

func GetCountrysHandler(w http.ResponseWriter, r *http.Request) {
	var country []models.Country
	connection.DB.Find(&country)
	defer json.NewEncoder(w).Encode(&country)
}

func GetCountrysByIdHandler(w http.ResponseWriter, r *http.Request) {
	var country models.Country
	params := mux.Vars(r)

	connection.DB.First(&country, "cnt_id = ?", params["country_id"])

	if country.CNT_ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Country Not Found"))
		return
	}

	defer json.NewEncoder(w).Encode(&country)
}

func GetCountrysByNameHandler(w http.ResponseWriter, r *http.Request) {
	var country models.Country
	params := mux.Vars(r)

	connection.DB.First(&country, "cnt_name = ?", params["country_id"])

	if country.CNT_ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Country Not Found"))
		return
	}

	defer json.NewEncoder(w).Encode(&country)
}
