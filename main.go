package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type List struct {
	ID       string `json:"_id,omitempty" bson:"_id,omitempty"`
	ItemName string `json:"itemName,omitempty" bson:"itemName, omitempty"`
}
type Res struct {
	ID string
}

var client *mongo.Client

func addItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Add("Content-type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var item List
	err = json.Unmarshal(body, &item)
	if err != nil {
		panic(err)
	}
	log.Println(item.ItemName)
	collection := client.Database("todo-challenge").Collection("todo")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, _ := collection.InsertOne(ctx, item)

	json.NewEncoder(w).Encode(result)

}

func getItems(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Add("Content-type", "application/json")
	log.Println("Requested")
	collection := client.Database("todo-challenge").Collection("todo")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var items []List
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Print("Error GET items")
		return
	}

	if err = cur.All(ctx, &items); err != nil {
		log.Print("Error GET items 2")
		return
	}
	log.Print(items)
	json.NewEncoder(w).Encode(items)

}

func showMain(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	path := "./client/index.html"
	http.ServeFile(w, r, path)

}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	router := httprouter.New()
	router.POST("/todo", addItem)
	router.GET("/todo", getItems)
	router.GET("/", showMain)
	// fs := http.FileServer(http.Dir("./client/static"))
	// httprouter.Router.Handle("/static/", http.StripPrefix("/static", fs))
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	log.Fatal(http.ListenAndServe(":3000", router))

}
