package operation

import (
	"encoding/json"
	"net/http"
)

func Sample(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "plain/text")
	type person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p := person{Name: "kritika", Age: 25}
	response, err := json.Marshal(p)
	if err != nil {
		http.Error(w, "server Error", 400)
	}

	//data := []byte("Hi kritika this side")
	if _, err := w.Write(response); err != nil {
		http.Error(w, "Server Error", 400)
	}
	//w.WriteHeader(http.StatusOK)
}
