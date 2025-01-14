package network

import (
	"fmt"
	"net/http"
)

const port string = ":8080"

var fileServer http.Handler = http.FileServer(http.Dir("./web/source"))

// Serve Web Application
func ServeWebApplication() {
	fmt.Println("Serving on Port: 8080")
	http.ListenAndServe(port, fileServer)
}
