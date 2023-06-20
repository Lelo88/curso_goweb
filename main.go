package main

import (
	"html/template"
	"log"
	"net/http"
)

var plantilla = template.Must(template.ParseGlob("public/*.html"))

func main() {

	css := http.FileServer(http.Dir("assets"))
	/*
	   Indicamos al servidor que traera los archivos css para que la pagina pueda tener estilos css
	   Estos deben de estar en una carpeta aparte del public
	*/
	http.Handle("/assets/", http.StripPrefix("/assets", css))
	// Rutas
	// Ruta de index
	http.HandleFunc("/", index)

	//Mensajes de confirmacion
	log.Println("Iniciando servidor")
	log.Println("Cargando Paginas")
	log.Println("Carga finalizada")
	//Iniciar servidor

	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	plantilla.ExecuteTemplate(w, "index.html", nil)
}
