package server

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>wat</h1><div>ok</div>")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	log.Info("Starting Polsino Server")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
