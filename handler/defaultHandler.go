package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
Default handler displaying service information.
*/
func defaultHandler(w http.ResponseWriter, r *http.Request) {

	template = template.Must(template.ParseFiles("site/mySite.html"))

	http.Handle("/static/", http.FileServer(http.Dir("static")))
	http.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := template.ExecuteTemplate(w, "mySite.html") {
			http.Error(w, http.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println(http.ListenAndServe(":8080"))

	// Define content type, so browser renders links correctly
	w.Header().Add("content-type", "text/html")

	// Prepare output returned to client
	output := "This service only offers the <a href=\"/diag\">/diag endpoint</a> that shows comprehensive HTTP " +
		"information for any received request. Please redirect any request to that endpoint."

	// Write output to client
	_, err := fmt.Fprintf(w, "%v", output)

	// Deal with error if any
	if err != nil {
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
	}
}
