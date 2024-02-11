package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Read the content of README.md
		content, err := ioutil.ReadFile("README.md")
		if err != nil {
			http.Error(w, "Could not read README.md", http.StatusInternalServerError)
			return
		}

		// Write the content to the response
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "%s", content)
	})

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
