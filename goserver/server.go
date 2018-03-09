// Server running at litbit.in

package main

import (
	"fmt"
	"net/http"
	"os"
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
			command := exec.Command("./send_long_range", "\"" + r.FormValue("firstname") + " " + r.FormValue("lastname") + "\"", r.FormValue("birthdate"), r.FormValue("phonenumber1"), r.FormValue("state"), msg)
			output, err := command.CombinedOutput()
			
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				fmt.Printf("Output:\n%s\n", output);
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "500 Internal Server Error\n")
			} else {
				http.Redirect(w, r, "/success/", 200)
			}
		}
	}
}

func handleFile(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s: %s\n", r.Method, r.URL)
	if _, err := os.Stat(r.URL.Path[1:]); os.IsNotExist(err) {
		http.ServeFile(w, r, "index.html");
	} else {
		http.ServeFile(w, r, r.URL.Path[1:]);
	}
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s: %s -> http://192.168.42.1/%s\n", r.Method, r.URL, r.RequestURI)
	http.Redirect(w, r, "http://192.168.42.1" + r.RequestURI, http.StatusMovedPermanently)
}

func main() {
	fmt.Printf("Listening on :80\n")

	http.HandleFunc("/submit/", handleSubmit)
	http.HandleFunc("/", handleFile)
	go func() {
		err := http.ListenAndServe(":80", nil)
		if err != nil {
			panic(err)
		}
	}()
	
	fmt.Printf("Listening on :443\n")
	err := http.ListenAndServeTLS(":443", "cert.cert", "cert.key", http.HandlerFunc(handleRedirect))
	if err != nil {
		panic(err)
	}
}
