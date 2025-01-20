package main

import (
	"api/clients"
	"api/dto"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

// Función principal
func main() {

	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	r := chi.NewRouter()

	// cors
	r.Use(corsOptions.Handler)

	// route
	r.Get("/search", searchItems)
	// server
	fmt.Println("Iniciando el servidor en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Manejador para buscar ítems por nombre
func searchItems(w http.ResponseWriter, r *http.Request) {
	// Obtener el parámetro de búsqueda desde la URL
	value := r.URL.Query().Get("value")
	field := r.URL.Query().Get("field")

	fmt.Print(value)
	fmt.Print(field)

	response,err := clients.ZincSearchClient(field,value)
	var responseDTO dto.Response
	/*if(err != nil){
		
	}*/
	// Deserializar el JSON en la variable 'person'
	err = json.Unmarshal([]byte(response), &responseDTO)
	if err != nil {
		log.Fatal(err)
	}

	// Responder con los ítems encontrados
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseDTO)
}