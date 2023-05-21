package main

import (
	"fmt"
	"net/http"

	routes "github.com/Shrut26/mongoApi/routers"
)

func main() {
	r := routes.Router()
	http.ListenAndServe(":4000", r)
	fmt.Println("Database connected")
	fmt.Println("Server started at PORT 4000")
}
