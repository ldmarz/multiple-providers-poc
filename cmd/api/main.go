package main

import (
	"fmt"
	"net/http"

	"providers_poc/cmd/api/middleware"
)

func main() {
	http.Handle("/verify", middleware.SelectProvider(middleware.RetrieveData(middleware.Response())))

	fmt.Println("Server started.")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(fmt.Errorf("server fails %w", err))
	}
}
