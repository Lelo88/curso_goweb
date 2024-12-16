package main

import (
	//"fmt"
	"fmt"
	"log"
	"net/http"
	"os"
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

func handler (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hola %s!", "mundo")
}

func viewHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/view/"):] // obtenemos el título de la página a cargar.
	p, err := loadPage(title) // cargamos la página.
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound) // si hay un error, redireccionamos al editor de la página.
		return
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body) // si no hay error, mostramos el contenido de la página.
}

func main(){
	/* p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body)) // si no lo casteo a string, me imprime el slice de bytes. */

	//http.HandleFunc("/", handler)

	http.HandleFunc("/view/", viewHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
	
}