package main

import (
	"fmt"
	"math"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {

	z := 0
	y := 16.0

	for z < 100000 {

		y += executando(y)
		z++
	}

	fmt.Fprintf(w, "Done!")
}

func executando(x float64) float64 {

	return math.Sqrt(x)
}
