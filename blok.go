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
		file, err := os.Open("html/index.html")
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


	http.HandleFunc("/style/style.css", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file CSS
		file, err := os.Open("style/style.css")
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
	
	// Menetapkan handler untuk rute "/"
	http.HandleFunc("/html/snk.html", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file HTML
		file, err := os.Open("html/snk.html")
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

	// Menetapkan handler untuk rute "/"
	http.HandleFunc("/html/layanankami.html", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file HTML
		file, err := os.Open("html/layanankami.html")
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

		// Menetapkan handler untuk rute "/"
		http.HandleFunc("/html/sewamobil.html", func(w http.ResponseWriter, r *http.Request) {
			// Membaca file HTML
			file, err := os.Open("html/sewamobil.html")
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
			// Menetapkan handler untuk rute "/"
			http.HandleFunc("/html/sewamotor.html", func(w http.ResponseWriter, r *http.Request) {
			// Membaca file HTML
			file, err := os.Open("html/sewamotor.html")
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

	// Menetapkan handler untuk rute "/"
	http.HandleFunc("/html/mobil/daihatsu-grandmax-blind-van.html", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file HTML
		file, err := os.Open("html/mobil/daihatsu-grandmax-blind-van.html")
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

	// Menetapkan handler untuk rute "/"
	http.HandleFunc("/html/mobil/daihatsu-xenia.html", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file HTML
		file, err := os.Open("html/mobil/daihatsu-xenia.html")
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

	// Menetapkan handler untuk rute "/"
	http.HandleFunc("/html/mobil/honda-brio.html", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file HTML
		file, err := os.Open("html/mobil/honda-brio.html")
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

	// Menetapkan handler untuk rute "/"
	http.HandleFunc("/html/mobil/honda-mobillio.html", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file HTML
		file, err := os.Open("html/mobil/honda-mobillio.html")
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

	// Menetapkan handler untuk rute "/"
	http.HandleFunc("/html/mobil/mitshubishi-pajero.html", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file HTML
		file, err := os.Open("html/mobil/mitshubishi-pajero.html")
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

	// Menetapkan handler untuk rute "/"
	http.HandleFunc("/html/mobil/suzuki-ertiga.html", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file HTML
		file, err := os.Open("html/mobil/suzuki-ertiga.html")
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

	// Menetapkan handler untuk rute "/"
	http.HandleFunc("/html/mobil/toyota-avanza.html", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file HTML
		file, err := os.Open("html/mobil/toyota-avanza.html")
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

	// Menetapkan handler untuk rute "/"
	http.HandleFunc("/html/mobil/toyota-innova-reborn.html", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file HTML
		file, err := os.Open("html/mobil/toyota-innova-reborn.html")
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

	// Menetapkan handler untuk rute "/"
	http.HandleFunc("/html/mobil/toyota-innova-zenix.html", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file HTML
		file, err := os.Open("html/mobil/toyota-innova-zenix.html")
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

	// Menetapkan handler untuk rute "/"
	http.HandleFunc("/html/motor/honda-beat.html", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file HTML
		file, err := os.Open("html/motor/honda-beat.html")
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

	// Menetapkan handler untuk rute "/"
	http.HandleFunc("/html/motor/honda-scoopy.html", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file HTML
		file, err := os.Open("html/motor/honda-scoopy.html")
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

	// Menetapkan handler untuk rute "/"
	http.HandleFunc("/html/motor/honda-vario.html", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file HTML
		file, err := os.Open("html/motor/honda-vario.html")
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

	// Menetapkan handler untuk rute "/"
	http.HandleFunc("/html/motor/suzuki-gsx-r150.html", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file HTML
		file, err := os.Open("html/motor/suzuki-gsx-r150.html")
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

	// Menetapkan handler untuk rute "/"
	http.HandleFunc("/html/motor/vespa-primavera.html", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file HTML
		file, err := os.Open("html/motor/vespa-primavera.html")
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

	// Menetapkan handler untuk rute "/"
	http.HandleFunc("/html/motor/yamaha-aerox.html", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file HTML
		file, err := os.Open("html/motor/yamaha-aerox.html")
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

	// Menetapkan handler untuk rute "/"
	http.HandleFunc("/html/motor/yamaha-nmax.html", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file HTML
		file, err := os.Open("html/motor/yamaha-nmax.html")
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

	// Menetapkan handler untuk rute "/"
	http.HandleFunc("/html/motor/yamaha-r15.html", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file HTML
		file, err := os.Open("html/motor/yamaha-r15.html")
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

	// Menetapkan handler untuk rute "/"
	http.HandleFunc("/html/motor/yamaha-xsr.html", func(w http.ResponseWriter, r *http.Request) {
		// Membaca file HTML
		file, err := os.Open("html/motor/yamaha-xsr.html")
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
