package utils

import (
	"net/http"
)

func AcessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {

	return
}

//func SelectRandomListings(n int,listings []scrape.Listing)[]int{
//	rand.Seed(time.Now().UnixNano())
//	selected := []scrape.Listing{}
//	for _= range MakeRange(0,n){
//		selected = append(selected,listings[rand.Intn(len(listings)-1)])
//	}
//}
