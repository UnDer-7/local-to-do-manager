package main

import (
	"fmt"
	"github.com/UnDer-7/local-to-do-manager/internal/config"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	envs := config.LoadEnvs()

	FileServerSPA(router, envs.Frontend.BasePath, "/home/under7/Workspace/go/src/github.com/UnDer-7/local-to-do-manager/website/dist")

	http.ListenAndServe(":3000", router)
}

// FileServerSPA source: https://stackoverflow.com/a/76534692
func FileServerSPA(router chi.Router, public, static string) {
	fmt.Println("Frontend base path: " + public)
	if public == "" {
		panic("Frontend base path cannot be empty, it needs at least to be an slash (/)")
	}

	if strings.ContainsAny(public, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	root, _ := filepath.Abs(static)
	if _, err := os.Stat(root); os.IsNotExist(err) {
		panic("Static Documents Directory Not Found | " + static)
	}

	fs := http.StripPrefix(public, http.FileServer(http.Dir(root)))

	if public != "/" && public[len(public)-1] != '/' {
		router.Get(public, http.RedirectHandler(public+"/", 301).ServeHTTP)
		public += "/"
	}

	router.Get(public+"*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file := strings.Replace(r.RequestURI, public, "/", 1)
		if _, err := os.Stat(root + file); os.IsNotExist(err) {
			http.ServeFile(w, r, path.Join(root, "index.html"))
			return
		}
		fs.ServeHTTP(w, r)
	}))
}
