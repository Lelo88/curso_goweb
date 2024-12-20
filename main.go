package main

import (
	//"fmt"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

// para eeste proyecto vamos a crear una página y su contenido,

type Page struct {
	Title string
	Body []byte // slice de bytes para el contenido de la página.
}

// crearemos un método para la estructura Page

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

// crearemos una función para cargar la página

func loadPage(title string) (*Page, error){
	filename := title + ".txt" // nombre del archivo que contiene la página a cargar
	body, err := os.ReadFile(filename) // leemos el archivo
	if err != nil {
		return nil, err // si hay un error retornamos nil y el error que ocurrió al leer el archivo.
	}

	return &Page{Title: title, Body: body}, nil
}

// crearemos un handler para mostrar el contenido de la página.
func handler (w http.ResponseWriter){
	fmt.Fprintf(w, "Hola %s!", "mundo")
}

// crearemos un handler para mostrar el contenido de la página.
// vamos a cambiar el renderizado de la página a un template.
func viewHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/view/"):] // obtenemos el título de la página a cargar.
	p, err := loadPage(title) // cargamos la página.
	if err != nil {
		p = &Page{Title: title} // si hay un error, creamos una nueva página con el título.
	}

	renderTemplate(w, "view", p)
}

// crearemos un handler para editar el contenido de la página.
// endpoint: localhost:3000/edit/TestPage
func editHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/edit/"):] // obtenemos el título de la página a editar.
	p, err := loadPage(title) // cargamos la página.
	if err != nil {
		p = &Page{Title: title} // si hay un error, creamos una nueva página con el título.
	}
	
	renderTemplate(w, "edit", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page){
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ejecutamos el template.
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/save/"):] // obtenemos el título de la página a editar.
	
	body := r.FormValue("body") // obtenemos el contenido de la página.
	
	p := &Page{Title: title, Body: []byte(body)} // creamos la página con el título y el contenido
	
	err := p.save() // guardamos la página.
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	http.Redirect(w, r, "/view/" + title, http.StatusFound) // redirigimos al usuario a la página.
}

func main(){
	/* p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body)) // si no lo casteo a string, me imprime el slice de bytes. */

	//http.HandleFunc("/", handler)

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
	
}