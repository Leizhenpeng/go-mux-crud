package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type HotDryNoodle struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
	Brand *Brand `json:"brand"`
}

type Brand struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	location string `json:"location"`
}

var hotDryNoodles []HotDryNoodle

func fakeHotDryNoodles() {
	// "标准热干面" 5 "蔡林记"
	// "热干面加蛋" 8 "蔡林记"
	hotDryNoodles = append(hotDryNoodles, HotDryNoodle{ID: 1, Title: "标准热干面", Price: 5, Brand: &Brand{ID: 1, Name: "蔡林记", location: "王家墩18号"}})
	hotDryNoodles = append(hotDryNoodles, HotDryNoodle{ID: 2, Title: "热干面加蛋", Price: 8, Brand: &Brand{ID: 1, Name: "蔡林记", location: "王家墩18号"}})

}
func main() {
	//use mux
	r := mux.NewRouter()
	fakeHotDryNoodles()
	r.HandleFunc("/api/noodles", getHotDryNoodles).Methods("GET")
	r.HandleFunc("/api/noodles/{id}", getHotDryNoodleById).Methods("GET")
	r.HandleFunc("/api/noodles/{id}", deleteHotDryNoodle).Methods("DELETE")
	r.HandleFunc("/api/noodles", createHotDryNoodle).Methods("POST")
	r.HandleFunc("/api/noodles/{id}", updateHotDryNoodle).Methods("PUT")

	fmt.Println("Server is running on port 8000", "open http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func updateHotDryNoodle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range hotDryNoodles {
		id, _ := strconv.Atoi(params["id"])
		if item.ID == id {
			hotDryNoodles = append(hotDryNoodles[:index], hotDryNoodles[index+1:]...)

			var hotDryNoodle HotDryNoodle
			_ = json.NewDecoder(request.Body).Decode(&hotDryNoodle)
			hotDryNoodle.ID = id

			hotDryNoodles = append(hotDryNoodles, hotDryNoodle)
			json.NewEncoder(writer).Encode(hotDryNoodles)
			return
		}
	}

}

func deleteHotDryNoodle(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("deleteHotDryNoodle")
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range hotDryNoodles {
		id, _ := strconv.Atoi(params["id"])
		if item.ID == id {
			hotDryNoodles = append(hotDryNoodles[:index], hotDryNoodles[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(hotDryNoodles)

}

func createHotDryNoodle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var hotDryNoodle HotDryNoodle
	_ = json.NewDecoder(request.Body).Decode(&hotDryNoodle)
	if hotDryNoodle.Title == "" || hotDryNoodle.Price == 0 {
		json.NewEncoder(writer).Encode(hotDryNoodles)
		return
	}
	//if already exist, update
	for index, item := range hotDryNoodles {
		if item.ID == hotDryNoodle.ID {
			hotDryNoodles[index] = hotDryNoodle
			json.NewEncoder(writer).Encode(hotDryNoodles)
			return
		}
	}

	hotDryNoodles = append(hotDryNoodles, hotDryNoodle)
	json.NewEncoder(writer).Encode(hotDryNoodles)

}

func getHotDryNoodleById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range hotDryNoodles {
		id, _ := strconv.Atoi(params["id"])
		if item.ID == id {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
}

func getHotDryNoodles(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(hotDryNoodles)
}
