// Server running at litbit.in

package main

import (
	"fmt"
	"net/http"
	"os/exec"
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

		if r.FormValue("firstname") == "" || r.FormValue("lastname") == "" || r.FormValue("phonenumber1") == "" {
			http.Redirect(w, r, "../", 400);
		} else {
			msg := ""
			if r.FormValue("message") != "" {
				msg = r.FormValue("message")
			}
			fmt.Printf("./send_long_range \"" + r.FormValue("firstname") + " " + r.FormValue("lastname") + "\" " + r.FormValue("birthdate") + " " + r.FormValue("phonenumber1") + " " + r.FormValue("state") + " " + msg + "\n")
			command := exec.Command("./send_long_range", "\"" + r.FormValue("firstname") + " " + r.FormValue("lastname") + "\"", r.FormValue("birthdate"), r.FormValue("phonenumber1"), r.FormValue("state"), msg)
			output, err := command.CombinedOutput()
			
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				fmt.Printf("Output: %s\n", output);
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "500 Internal Server Error\n")
			} else {
				http.Redirect(w, r, "/success/", 200)
			}
		}
/*
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
*/
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
