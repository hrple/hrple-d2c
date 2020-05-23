package main

import (
	"fmt"
	"net/http"

	web "github.com/hrple/common/server"
)

func main() {
	web.Get("/hello", helloWorldStd)
	logger := web.GetLogger()
	if err := web.Start(""); err != nil {
		logger.Fatalf("Could not listen: %v\n", err)
	}
}

func helloWorldStd(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "hello %v \n", r.URL)
}
