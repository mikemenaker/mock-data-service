package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"io/ioutil"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func DataIndex(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("data.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, str)
}

func DataShow(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//todoId := vars["dataType"]

	genData := make([]int, 0)
	genData = append(genData, random(0, 100))
	genData = append(genData, random(0, 100))
	genData = append(genData, random(0, 100))
	genData = append(genData, random(0, 100))

	data := Data{Data: genData}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
	//fmt.Fprintln(w, "Todo show:", todoId, genData)
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
