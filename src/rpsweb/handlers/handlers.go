package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"rpsweb/rps"
	"strconv"
	"text/template"
)

const (
	templateDir  = "templates/"
	templateBase = templateDir + "base.html"
)

type Player struct {
	Name string
}

var player Player

func renderTemplate(w http.ResponseWriter, page string, data any) {
	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page))
	// tpl, err := template.ParseFiles("templates/base.html", "templates/index.html")
	// if err != nil {
	// 	http.Error(w, "Error al analizar plantillas", http.StatusInternalServerError)
	// 	return
	// }

	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func restartValue() {
	player.Name = ""
	rps.ComputerScore = 0
	rps.PlayerScore = 0
}

func Index(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, "index.html", nil)
}
func NewGame(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, "new_game.html", nil)
}
func Game(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error Parsing Form", http.StatusBadRequest)
			log.Println(err)
			return
		}
		player.Name = r.Form.Get("name")
	}

	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusFound)
	}

	renderTemplate(w, "game.html", player)
}
func Play(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
func About(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, "about.html", nil)
}
