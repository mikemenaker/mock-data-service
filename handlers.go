package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
	"strconv"
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

func DataShowStats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	countStr := vars["count"]
	count, _ := strconv.ParseInt(countStr, 10, 64)

	genData := make([]int, 0)

	for i := 0; i <= int(count); i++ {
		genData = append(genData, random(0, 100))
	}

	data := StatsData{Data: genData}

	sendResponse(w, data)
}

func DataShowText(w http.ResponseWriter, r *http.Request) {
	randNum := random(0, 100)
	data := TextData{Data: strconv.Itoa(randNum) + " request per hour for last 24 hours"}

	sendResponse(w, data)
}

func DataShowTable(w http.ResponseWriter, r *http.Request) {
	data := TableData{Data: []TableDataItem{
		{Name: "user1", Value: strconv.Itoa(random(0, 100)), Value2: strconv.Itoa(random(0, 100))},
		{Name: "user2", Value: strconv.Itoa(random(0, 100)), Value2: strconv.Itoa(random(0, 100))},
		{Name: "user3", Value: strconv.Itoa(random(0, 100)), Value2: strconv.Itoa(random(0, 100))},
		{Name: "user4", Value: strconv.Itoa(random(0, 100)), Value2: strconv.Itoa(random(0, 100))},
	},
	}

	sendResponse(w, data)
}

func sendResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
