package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/Cudesi-service-S-A-C/geo_service/service/api/country"
	"github.com/Cudesi-service-S-A-C/geo_service/service/api/department"
	"github.com/Cudesi-service-S-A-C/geo_service/service/api/district"
	"github.com/Cudesi-service-S-A-C/geo_service/service/api/province"
	"github.com/Cudesi-service-S-A-C/geo_service/service/connection"
	"github.com/Cudesi-service-S-A-C/geo_service/service/utils"
)

var portServer = utils.ReadEnv("SERVER_PORT_DEV")

func getRoutesV1(router *mux.Router) *mux.Router {

	route_v1 := router.PathPrefix("/v1").Subrouter()
	route_v1.HandleFunc("/countrys", country.GetCountrysHandler).Methods(http.MethodGet)
	route_v1.HandleFunc("/country/byid/{country_id}", country.GetCountrysByIdHandler).Methods(http.MethodGet)
	route_v1.HandleFunc("/country/byname/{country_id}", country.GetCountrysByNameHandler).Methods(http.MethodGet)

	route_v1.HandleFunc("/departments", department.GetDepartmentsHandler).Methods(http.MethodGet)
	route_v1.HandleFunc("/departments/country/{country_id}", department.GetDepartmentByCountryHandler).Methods(http.MethodGet)
	route_v1.HandleFunc("/departments/idcountry/{country_id}", department.GetDepartmentByIdCountryHandler).Methods(http.MethodGet)

	route_v1.HandleFunc("/provinces", province.GetProvincesHandler).Methods(http.MethodGet)
	route_v1.HandleFunc("/provinces/department/{department_id}", province.GetProvincesByDepartmentHandler).Methods(http.MethodGet)
	route_v1.HandleFunc("/provinces/iddepartment/{department_id}", province.GetProvincesByIdDepartmentHandler).Methods(http.MethodGet)

	route_v1.HandleFunc("/districts", district.GetDistrictsHandler).Methods(http.MethodGet)
	route_v1.HandleFunc("/districts/province/{province_id}", district.GetDistrictsByProvinceHandler).Methods(http.MethodGet)
	route_v1.HandleFunc("/districts/idprovince/{province_id}", district.GetDistrictsByIdProvinceHandler).Methods(http.MethodGet)

	return route_v1
}

func main() {
	connection.Connection()

	router := mux.NewRouter()

	router = getRoutesV1(router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://demo.localhost:3000"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	handler := c.Handler(router)

	http.ListenAndServe(":"+portServer, handler)
}
