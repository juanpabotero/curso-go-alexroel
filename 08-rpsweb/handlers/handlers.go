package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"rpsweb/rps"
	"strconv"
)

type PLayer struct {
	Name string
}

var player PLayer

const (
	dirTemplate  = "templates/"
	BaseTemplate = dirTemplate + "base.html"
)

// http.ResponseWriter se utiliza para escribir la respuesta
// http.Request es una estructura que representa la petición http
func Index(w http.ResponseWriter, r *http.Request) {
	resetValues()
	RenderTemplate(w, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	resetValues()
	RenderTemplate(w, "new-game.html", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// ParseForm: analiza el formulario y los parámetros de la consulta
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Form.Get: devuelve el primer valor para la clave del formulario
		player.Name = r.Form.Get("name")
	}
	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusSeeOther)
	}
	RenderTemplate(w, "game.html", player)
}

// Controlador de jugar
func Play(w http.ResponseWriter, r *http.Request) {
	// Get: devuelve el primer valor asociado con la clave
	// strconv.Atoi: convierte la cadena a un entero
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	// MarshalIndent: serializa un valor en formato JSON
	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	// Header().Set: establece el valor de la clave en el encabezado
	w.Header().Set("Content-Type", "application/json")
	// Write: escribe la respuesta en el cuerpo de la respuesta
	w.Write(out)
}

func About(w http.ResponseWriter, r *http.Request) {
	resetValues()
	RenderTemplate(w, "about.html", nil)
}

func RenderTemplate(w http.ResponseWriter, page string, data any) {
	// Must: envuelve un valor y un error, si el error no es nil, Must ejecuta panic
	// ParseFiles carga un archivo de plantilla y lo analiza
	tmpl := template.Must(template.ParseFiles(BaseTemplate, dirTemplate+page))

	// ExecuteTemplate renderiza la plantilla y escribe su salida en w
	err := tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panicln(err)
		return
	}
}

// reiniciar valores
func resetValues() {
	player.Name = ""
	rps.ComputerScore = 0
	rps.PlayerScore = 0
}
