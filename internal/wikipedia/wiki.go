package wiki

import (
	"io"
	"log"
	"net/http"
)

func getMovies() []byte {
	resp, err := http.Get("")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return body
}
