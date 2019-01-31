package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Item is the schema of an item
type Item struct {
	id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// array of items
var items []Item

//GetItemByID gets an item by it's id
func GetItemByID(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	for _, item := range items {

		if item.id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// if not found return empty object with Item structure
	json.NewEncoder(w).Encode(&Item{})
}

//CreateItem creates a new item
func CreateItem(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var item Item
	_ = json.NewDecoder(req.Body).Decode(&item)
	item.id = params["id"]
	items = append(items, item)

	json.NewEncoder(w).Encode(items)

}

//GetItems retrieves all items
func GetItems(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(items)
}

//DeleteItem removes item from array
func DeleteItem(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	for index, item := range items {
		if item.id == params["id"] {
			items = append(items[:index], items[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(items)
}

func main() {
	fmt.Println("magic is happening on port 8081")

	//creating local array
	items = append(items, Item{id: "1", Name: "phone", Description: "a device used for calling people"})
	items = append(items, Item{id: "2", Name: "laptop", Description: "a mobile computing device"})

	router := mux.NewRouter()
	router.HandleFunc("/items", GetItems).Methods("GET")
	router.HandleFunc("/items/{id}", GetItemByID).Methods("GET")
	router.HandleFunc("/items/{id}", CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", router))

}
