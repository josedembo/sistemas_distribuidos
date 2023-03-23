package main

import (
	"net/http"
)

func main(){
	http.HandleFunc("/me", func (w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("Ola mundo from me"))
	})
	http.ListenAndServe(":8080", nil)
}