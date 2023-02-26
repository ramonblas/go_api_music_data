package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"go_api/src/model"
)

func responseSong(w http.ResponseWriter, status int, results model.Song) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

func responseSongs(w http.ResponseWriter, status int, results []model.Song) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

var collection = getSession().DB("my_music").C("Songs")


func getSession() *mgo.Session {
	
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
		}

	return session
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Song DB")
}

func SongList(w http.ResponseWriter, r *http.Request) {
	var results []model.Song
	err := collection.Find(nil).Sort("-_id").All(&results)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Resultados: ", results)
	}

	responseSongs(w, 200, results)
}

func SongShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Song_id := params["id"]

	if !bson.IsObjectIdHex(Song_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(Song_id)

	results := model.Song{}
	err := collection.FindId(oid).One(&results)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	responseSong(w, 200, results)
}

func SongAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var Song_data model.Song
	err := decoder.Decode(&Song_data)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	err = collection.Insert(Song_data)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	responseSong(w, 200, Song_data)
}

func SongUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Song_id := params["id"]

	if !bson.IsObjectIdHex(Song_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(Song_id)

	decoder := json.NewDecoder(r.Body)

	var Song_data model.Song
	err := decoder.Decode(&Song_data)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	}

	defer r.Body.Close()

	document := bson.M{"_id": oid}
	change := bson.M{"$set": Song_data}
	err = collection.Update(document, change)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	responseSong(w, 200, Song_data)
}

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (this *Message) setStatus(data string) {
	this.Status = data
}

func (this *Message) setMessage(data string) {
	this.Message = data
}

func SongRemove(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Song_id := params["id"]

	if !bson.IsObjectIdHex(Song_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(Song_id)

	err := collection.RemoveId(oid)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	//results := Message{"success", "La Cancion con ID "+Song_id+" ha sido borrada correctamente"}
	message := new(Message)

	message.setStatus("success")
	message.setMessage("La Cancion con ID " + Song_id + " ha sido borrada correctamente")

	results := message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}
