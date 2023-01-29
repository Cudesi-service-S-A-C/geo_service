package utils

import (
	"fmt"
	"net/http"
	"strconv"
)

func ValidationParamsString(param string, w http.ResponseWriter) bool {

	value, error := strconv.Atoi(param)

	if error == nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Param Not Permit \n"))
		fmt.Println(value)
		return false
	}

	return true
}

func ValidationParamsInt(param string, w http.ResponseWriter) bool {

	value, error := strconv.Atoi(param)

	if error != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Param Not Permit \n"))
		fmt.Println(value)
		return false
	}

	return true
}
