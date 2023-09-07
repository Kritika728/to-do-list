package operation

import "net/http"

func Add(w http.ResponseWriter, r *http.Request) {
	request, err := r.GetBody()

}
