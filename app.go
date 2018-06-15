package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	r := chi.NewRouter()
	r.Get("/All_match_data", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://worldcup.sfg.io/matches")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		responseString := string(responseData)
		fmt.Println(responseString)

	})
	r.Get("/Today_matches", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://worldcup.sfg.io/matches/today")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		responseString := string(responseData)
		fmt.Println(responseString)
		fmt.Println()

	})
	r.Get("/Current_matches", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://worldcup.sfg.io/matches/current")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		responseString := string(responseData)
		fmt.Println(responseString)

	})
	http.ListenAndServe(":3000", r)
}
