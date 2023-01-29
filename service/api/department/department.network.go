package department

import (
	"encoding/json"
	"net/http"

	"github.com/Cudesi-service-S-A-C/geo_service/service/connection"
	"github.com/Cudesi-service-S-A-C/geo_service/service/models"
	"github.com/Cudesi-service-S-A-C/geo_service/service/utils"
	"github.com/gorilla/mux"
)

func GetDepartmentsHandler(w http.ResponseWriter, r *http.Request) {

	var department []models.Department
	var count_deparment int64

	connection.DB.Find(&department).Count(&count_deparment)

	if count_deparment == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Departments Not Found"))
		return
	}

	defer json.NewEncoder(w).Encode(&department)
}

func GetDepartmentByCountryHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var department []models.Department
	var count_deparment int64

	validation := utils.ValidationParamsString(params["country_id"], w)

	if !validation {
		return
	}

	connection.DB.Joins("left join geo_countries on geo_countries.cnt_id = geo_departments.cnt_id").Where("geo_countries.cnt_name = ? and geo_departments.dep_status = true", params["country_id"]).Find(&department).Count(&count_deparment)

	if count_deparment == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Departments Not Found"))
		return
	}

	defer json.NewEncoder(w).Encode(&department)

}

func GetDepartmentByIdCountryHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var department []models.Department
	var count_deparment int64

	validation := utils.ValidationParamsInt(params["country_id"], w)

	if !validation {
		return
	}

	connection.DB.Where("cnt_id = $1 and dep_status = true", params["country_id"]).Find(&department).Count(&count_deparment)

	if count_deparment == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Departments Not Found"))
		return
	}

	defer json.NewEncoder(w).Encode(&department)

}
