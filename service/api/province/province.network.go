package province

import (
	"encoding/json"
	"net/http"

	"github.com/Cudesi-service-S-A-C/geo_service/service/connection"
	"github.com/Cudesi-service-S-A-C/geo_service/service/models"
	"github.com/Cudesi-service-S-A-C/geo_service/service/utils"
	"github.com/gorilla/mux"
)

func GetProvincesHandler(w http.ResponseWriter, r *http.Request) {
	var province []models.Province
	var count_province int64

	connection.DB.Find(&province).Count(&count_province)

	if count_province == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Provinces Not Found"))
		return
	}

	defer json.NewEncoder(w).Encode(&province)
}

func GetProvincesByDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var province []models.Province
	var count_province int64

	validation := utils.ValidationParamsString(params["department_id"], w)

	if !validation {
		return
	}

	connection.DB.Joins("left join geo_departments on geo_departments.dep_id = geo_provinces.dep_id").Where("geo_departments.dep_name = ? and geo_provinces.prv_status = true", params["department_id"]).Find(&province).Count(&count_province)

	if count_province == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Provinces Not Found"))
		return
	}

	defer json.NewEncoder(w).Encode(&province)
}

func GetProvincesByIdDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var province []models.Province
	var count_province int64

	validation := utils.ValidationParamsInt(params["department_id"], w)

	if !validation {
		return
	}

	connection.DB.Where("dep_id = $1 and prv_status = true", params["department_id"]).Find(&province).Count(&count_province)

	if count_province == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Provinces Not Found"))
		return
	}

	defer json.NewEncoder(w).Encode(&province)
}
