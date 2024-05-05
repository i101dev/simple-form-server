package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST was successful\n")

	name := r.FormValue("user_name")
	email := r.FormValue("user_email")
	subject := r.FormValue("user_subject")
	message := r.FormValue("user_message")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Email = %s\n", email)
	fmt.Fprintf(w, "Subject = %s\n", subject)
	fmt.Fprintf(w, "Message = %s\n", message)
}

func hellowHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hellow!")
}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", hellowHandler)

	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Println("failed to boot server")
		log.Fatal(err)
	}
}
