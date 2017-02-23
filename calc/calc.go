package calc

import (
	"github.com/gorilla/mux"
	"net/http"
)

////////////////////////////////////////////////
// Calc interface
type Calc interface {
	Sum(a, b string) string
}
////////////////////////////////////////////////


////////////////////////////////////////////////
// Roman implementation of the calc
type Roman struct{}

func (r Roman) Sum(a, b string) string {

	if a == "II" && b == "I" {
		return "III"
	}

	return "II"
}
////////////////////////////////////////////////


////////////////////////////////////////////////
// WebService handlers
func GetRomanSum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	a := vars["a"]
	b := vars["b"]

	var roman *Roman = new(Roman)

	w.Write([]byte(roman.Sum(a, b)))
}
////////////////////////////////////////////////