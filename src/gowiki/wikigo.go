package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"
)

type Page struct {
	Tittle string
	Body   []byte
}

func (p *Page) save() error {
	filename := p.Tittle + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(tittle string) (*Page, error) {
	filename := tittle + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Tittle: tittle, Body: body}, nil
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// func getTittle(w http.ResponseWriter, r *http.Request) (string, error) {
// 	m := validPath.FindStringSubmatch(r.URL.Path)
// 	if m == nil {
// 		http.NotFound(w, r)
// 		return "", errors.New("Titulo de pagina invalido")
// 	}
// 	return m[2], nil
// }

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}

}

func viewHandler(w http.ResponseWriter, r *http.Request, tittle string) {

	// tittle, err := getTittle(w, r)
	// if err != nil {
	// 	return
	// }
	p, err := loadPage(tittle)
	if err != nil {
		http.Redirect(w, r, "/edit/"+tittle, http.StatusFound)
		return
	}
	renderTemplates(w, "view", p)

}

func editHandler(w http.ResponseWriter, r *http.Request, tittle string) {
	// tittle, err := getTittle(w, r)
	// if err != nil {
	// 	return
	// }
	p, err := loadPage(tittle)
	if err != nil {
		p = &Page{Tittle: tittle}
	}

	renderTemplates(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, tittle string) {
	// tittle, err := getTittle(w, r)
	// if err != nil {
	// 	return
	// }
	body := r.FormValue("body")
	p := &Page{Tittle: tittle, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+tittle, http.StatusFound)
	renderTemplates(w, "edit", p)
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplates(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// p1 := &Page{Tittle: "TestPage", Body: []byte("Esta es una página de muestra")}
	// p1.save()
	// p2, _ := loadPage("TestPage")

	// fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))

}
