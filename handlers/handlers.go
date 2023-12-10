package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"osprey-backend/db"
	"osprey-backend/models"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type ErrorLog struct {
	Key       string `json:"key"`
	ErrorType string `json:"error_type"`
	Message   string `json:"message"`
}

var logs = []ErrorLog{}

func NewLog(w http.ResponseWriter, r *http.Request) {
	client := db.GetClient()

	logsColl := client.Database("test").Collection("logs")
	projColl := client.Database("test").Collection("projects")

	//Get api key in url
	urlParams := r.URL.Query()
	apiKey := urlParams.Get("api_key")

	var project models.Project

	filter := bson.M{"api_key": apiKey}
	err := projColl.FindOne(context.TODO(), filter).Decode(&project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newLog models.ErrorLog
	newLog.Project = project.ID

	err = json.NewDecoder(r.Body).Decode(&newLog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := logsColl.InsertOne(context.TODO(), newLog)

	fmt.Println(newLog)

	fmt.Println(res)

	w.WriteHeader(http.StatusCreated)
}

func GetLogs(w http.ResponseWriter, r *http.Request) {
	client := db.GetClient()

	coll := client.Database("test").Collection("logs")

	var logs []models.ErrorLog

	params := mux.Vars(r)

	projId, _ := primitive.ObjectIDFromHex(params["projId"])

	filter := bson.M{"project": projId}
	fmt.Println(projId)

	cur, err := coll.Find(context.TODO(), filter)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var currLog models.ErrorLog
		err := cur.Decode(&currLog)
		if err != nil {
			log.Fatal(err)
		}
		logs = append(logs, currLog)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(logs)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(logs)

}

func CreateProject(w http.ResponseWriter, r *http.Request) {

	client := db.GetClient()

	coll := client.Database("test").Collection("projects")

	var newProject models.Project

	_ = json.NewDecoder(r.Body).Decode(&newProject)
	newProject.ApiKey = "123"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := coll.InsertOne(ctx, newProject)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)

}

func GetProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project

	client := db.GetClient()

	coll := client.Database("test").Collection("projects")

	params := mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}

	err := coll.FindOne(context.TODO(), filter).Decode(&project)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(project)
}
