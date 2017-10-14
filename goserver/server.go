// Server running at litbit.in

package main

import (
	"fmt"
	"net/http"
)

func handleSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "400 Bad Request\n")
	} else {
		err := r.ParseForm();
		if err != nil {
			panic(err);
		}

		fmt.Fprintf(w, "Form submission detected!\n");
		fmt.Fprintf(w, "firstname: %s\n", r.FormValue("firstname"));
		fmt.Fprintf(w, "middlename: %s\n", r.FormValue("middlename"));
		fmt.Fprintf(w, "lastname: %s\n", r.FormValue("lastname"));
		fmt.Fprintf(w, "birthdate: %s\n", r.FormValue("birthdate"));
		fmt.Fprintf(w, "condition: %s\n", r.FormValue("condition"));
		fmt.Fprintf(w, "phonenumber1: %s\n", r.FormValue("phonenumber1"));
		fmt.Fprintf(w, "phonenumber2: %s\n", r.FormValue("phonenumber2"));
		fmt.Fprintf(w, "phonenumber3: %s\n", r.FormValue("phonenumber3"));
		fmt.Fprintf(w, "message: %s\n", r.FormValue("message"));
		fmt.Fprintf(w, "plan: %s\n", r.FormValue("plan"));
	}
}

func main() {
	fmt.Printf("Listening on :80\n")

	http.HandleFunc("/submit/", handleSubmit)
	http.Handle("/", http.FileServer(http.Dir("./")))
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
