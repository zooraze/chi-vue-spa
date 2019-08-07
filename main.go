package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi"
)

var staticDir string
var serverPort string

func main() {

	staticDir := "web/dist"
	
	// TODO(zooraze): Correctly resolve direct URI access attempts by user
	// Router
	router := chi.NewRouter()
	routes := Routes{
		staticDir: staticDir,
		apiKey: "key",
		apiSecret: "secret",
	}

	router.Get("/", func(resp http.ResponseWriter, req *http.Request) {
		http.ServeFile(resp, req, staticDir + "/index.html")
	})

	// API Routes
	router.HandleFunc("/api/data", routes.apiDataRoute)

	// Serve static files
	FileServer(router, "/", http.Dir(staticDir))

	// Start server
	serverPort = os.Getenv("PORT")
	if len(serverPort) == 0 {
		serverPort = "3333"
	}

	fmt.Printf("### Starting server listening on %v\n", serverPort)
	fmt.Printf("### Serving static content from '%v'\n", staticDir)
	fmt.Printf("### Browse: http://localhost:3333\n")
	http.ListenAndServe(":"+serverPort, router)
}

// Serve static files
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		fs.ServeHTTP(resp, req)
	}))
}
