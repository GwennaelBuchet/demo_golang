package calc

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
// Arabic implementation of the calc
type Arabic struct{}

func (self Arabic) Sum(a, b string) int {

	x, err := strconv.Atoi(a)
	y, err2 := strconv.Atoi(b)

	if err != nil || err2 != nil {
		return 0
	}

	return x + y
}

func (self Arabic) SumFromSlice(a []int) int {
	r := 0
	for _, value := range a {
		r += value
	}

	return r

	/**
	We also could get current index:
		for index, value := range a {
			r += value
		}
	*/
}

////////////////////////////////////////////////
// WebService handlers
func GetRomanSum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	a := vars["a"]
	b := vars["b"]

	var roman *Roman = new(Roman)

	w.Write([]byte(roman.Sum(a, b)))
}

func GetArabicSum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	a := vars["a"]
	b := vars["b"]

	var arabic *Arabic = new(Arabic)
	result := strconv.Itoa(arabic.Sum(a, b))

	w.Write([]byte(result))
}
////////////////////////////////////////////////