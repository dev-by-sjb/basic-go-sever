package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	var port = os.Getenv("PORT")
	if port == "" {
		log.Fatal("There is no port in dotenv")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello People"))

	})

	println("Server is listening at http://%v", port)
	log.Fatal(http.ListenAndServe(port, nil))

}

// c := []byte("Hello my people")
