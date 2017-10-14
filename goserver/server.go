// Server running at litbit.in

package main

import (
	"fmt"
	"net/http"
)

func handleSubmit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Submitted form!");
}

func main() {
	fmt.Printf("Listening on :80\n")

	http.Handle("/", http.FileServer(http.Dir("./")));
	http.HandleFunc("/submit/", handleSubmit)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
