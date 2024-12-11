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

func main(){
	/* p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body)) // si no lo casteo a string, me imprime el slice de bytes. */

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":3000", nil))
	
}