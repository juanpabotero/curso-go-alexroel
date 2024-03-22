package main

import (
	"log"
	"net/http"
	"rpsweb/handlers"
)

func main() {
	// crear enrutador
	router := http.NewServeMux()

	// manejo de archivos estáticos
	// FileServer devuelve un servidor de archivos que sirve archivos desde el sistema de archivos dado
	fs := http.FileServer(http.Dir("static"))

	// ruta para archivos estáticos
	// StripPrefix devuelve un controlador que sirve el contenido de un sistema de archivos en la ruta dada
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	// rutas
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	// levantar servidor
	port := ":8080"
	log.Printf("Server running on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
