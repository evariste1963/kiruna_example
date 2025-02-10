package main

import (
	"fmt"
	"html/template"
	"kiruna_example/internal/platform"
	"net/http"

	"github.com/sjc5/kiruna"
)

func main() {
	// Health check endpoint
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Serve static files from "dist/kiruna/static/public" directory, accessible at "/public/"
	http.Handle("/public/", platform.Kiruna.MustGetServeStaticHandler("/public/", true))

	// Serve an HTML file using html/template
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		privateFS, err := platform.Kiruna.GetPrivateFS()
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFS(privateFS, "index.go.html")
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, struct{ Kiruna *kiruna.Kiruna }{
			Kiruna: platform.Kiruna,
		})
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
		}
	})

	port := kiruna.MustGetPort()

	fmt.Printf("Starting server on: http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
