package utils

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

func StringInArray(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func Json200(payload any, w http.ResponseWriter) {
	res, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetStr(n int) string {
	if n >= 100_000 {
		k := math.Floor(float64(n) / 1000)
		return fmt.Sprintf("> %gk", k)
	} else {
		h := math.Floor(float64(n) / 100)
		h /= 10
		return fmt.Sprintf("> %gk", h)
	}
}
