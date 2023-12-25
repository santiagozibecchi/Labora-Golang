package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"jsonApp/utils"
)

func main() {
	// Configura los manejadores antes de iniciar el servidor
	http.HandleFunc("/api/v1/sumEvens", handlerEvens)
	http.HandleFunc("/api/v1/sumOdds", handlerOdds)

	// Inicia el servidor en el puerto 8080
	fmt.Println("Iniciando el servidor en http://localhost:8080...")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}

type Num struct {
	Num int `json:"num"`
}

type response struct {
	EvenSum int `json:"EvenSum"`
	OddSum int `json:"OddSum"`
}

func handlerEvens(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		stringNumber := r.URL.Query().Get("number")
		number, err := strconv.Atoi(stringNumber)

		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		sum := utils.SumaPares(number)

		// Envía la respuesta al cliente
		response := fmt.Sprintf("La suma de los pares es: %d", sum)
		w.Write([]byte(response))
	} 

	if r.Method == "POST" {
		var sumaPar response

		err := json.NewDecoder(r.Body).Decode(&sumaPar)

		if err != nil {
			http.Error(w, "No se puede convertir a número", http.StatusBadRequest)
			return
		}

		evenSum := utils.SumaPares(sumaPar.EvenSum)
		sumaPar.EvenSum = evenSum

		// Codificar la respuesta actualizada como JSON y enviarla al cliente
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(sumaPar)
		if err != nil {
			http.Error(w, "Error al codificar la respuesta como JSON", http.StatusInternalServerError)
			return
		}
	}
	
}

func handlerOdds(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		stringNumber := r.URL.Query().Get("number")
		number, err := strconv.Atoi(stringNumber)

		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Llamada a la función que suma números impares
		sum := utils.SumaImpares(number)

		// Envía la respuesta al cliente
		response := fmt.Sprintf("La suma de los impares es: %d", sum)
		w.Write([]byte(response))
	}
}


