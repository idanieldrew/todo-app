package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"rest/database"
	"rest/entity"
)

/*func Index(rw http.ResponseWriter,r *http.Request)  {
	var posts []entity.Post
	database.Connector.Find(&posts)
	rw.Header().Set("Content-Type","application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(posts)
}*/

func Store(rw http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var post entity.Post
	json.Unmarshal(body, &post)
	database.Connector.Create(&post)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(201)
	json.NewEncoder(rw).Encode(post)
}

func Show(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var posts []entity.Post
	database.Connector.First(&posts, key)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(posts)
}

func Test(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("kjn")
	rw.WriteHeader(500)
}
