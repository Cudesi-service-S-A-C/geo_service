package district

import (
	"encoding/json"
	"net/http"

	"github.com/Cudesi-service-S-A-C/geo_service/service/connection"
	"github.com/Cudesi-service-S-A-C/geo_service/service/models"
	"github.com/Cudesi-service-S-A-C/geo_service/service/utils"
	"github.com/gorilla/mux"
)

func GetDistrictsHandler(w http.ResponseWriter, r *http.Request) {
	var district []models.District
	var count_district int64

	connection.DB.Find(&district).Count(&count_district)

	if count_district == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Districts Not Found"))
		return
	}

	defer json.NewEncoder(w).Encode(&district)
}

func GetDistrictsByProvinceHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var district []models.District
	var count_district int64

	validation := utils.ValidationParamsString(params["province_id"], w)

	if !validation {
		return
	}

	connection.DB.Joins("left join geo_provinces on geo_provinces.prv_id = geo_districts.prv_id").Where("geo_provinces.prv_name = ? and geo_districts.dis_status = true", params["province_id"]).Find(&district).Count(&count_district)

	if count_district == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("District Not Found"))
		return
	}

	defer json.NewEncoder(w).Encode(&district)
}

func GetDistrictsByIdProvinceHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var district []models.District
	var count_district int64

	validation := utils.ValidationParamsInt(params["province_id"], w)

	if !validation {
		return
	}

	connection.DB.Where("prv_id = $1 and dis_status = true", params["province_id"]).Find(&district).Count(&count_district)

	if count_district == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("District Not Found"))
		return
	}

	defer json.NewEncoder(w).Encode(&district)
}
