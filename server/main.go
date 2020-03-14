package server

import (
	"fmt"
	"net/http"

	"github.com/mneil/polsino/server/middleware"
	"github.com/mneil/polsino/server/request"
	"github.com/sirupsen/logrus"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

	}
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	logrus.Info("Starting Polsino Server")
	auth := middleware.Auth{}
	common := []request.Handler{
		auth.Verify,
	}
	http.HandleFunc("/", request.Handlers(common))
	logrus.Fatal(http.ListenAndServe(":9000", nil))
}
