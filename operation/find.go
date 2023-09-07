package operation

import (
	"context"
	"encoding/json"
	"fmt"
	"go-quickstart/to-do-list/database"
	"go-quickstart/to-do-list/model"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	var response []model.Task
	client := database.DB()
	db := client.Database("ToDo").Collection("Tasks")
	cursor, err := db.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err)
	}
	if err := cursor.All(context.TODO(), &response); err != nil {
		panic(err)
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "issue in marshaling the struct ", http.StatusInternalServerError)
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}

}
