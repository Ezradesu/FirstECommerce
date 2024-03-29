package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Menetapkan handler untuk rute "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file HTML
		file, err := os.Open("home.html")
		if err != nil {
			http.Error(w, "Unable to open HTML file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Menyalin isi file HTML ke ResponseWriter
		_, err = io.Copy(w, file)
		if err != nil {
			http.Error(w, "Unable to copy HTML content", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("styles.css", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file CSS
		file, err := os.Open("styles.css")
		if err != nil {
			http.Error(w, "Unable to open CSS file", http.StatusInternalServerError)
			return
		}
		defer file.Close()
	
		// Menyalin isi file CSS ke ResponseWriter
		_, err = io.Copy(w, file)
		if err != nil {
			http.Error(w, "Unable to copy CSS content", http.StatusInternalServerError)
			return
		}
	})

	// nentuin port yang dipake
	addr := "localhost:8080"

	// biar ketauan webnya jalan
	fmt.Printf("Server sedang berjalan di http://%s\n", addr)

	// start server
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
