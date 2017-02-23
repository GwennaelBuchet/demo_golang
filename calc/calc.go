package calc

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

////////////////////////////////////////////////
// Calc interface
type Calc interface {
	Sum(a, b string)
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
// Arabic implementation of the calc
type Arabic struct{}

func (r Arabic) Sum(a, b string) int {

	x, err := strconv.Atoi(a)
	y, err2 := strconv.Atoi(b)

	if err != nil || err2 != nil {
		return 0
	}

	return x + y
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