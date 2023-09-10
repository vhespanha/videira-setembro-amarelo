package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var mu sync.Mutex

func main() {
	http.HandleFunc("/save-feeling", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		feeling := r.FormValue("feeling")
		if feeling == "" {
			http.Error(w, "Sentimento vazio", http.StatusBadRequest)
			return
		}

		mu.Lock()
		defer mu.Unlock()

		// Abra o arquivo CSV para escrita, ou crie se não existir
		file, err := os.OpenFile("sentiments.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			http.Error(w, "Erro ao abrir o arquivo CSV", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Crie um escritor CSV
		writer := csv.NewWriter(file)
		defer writer.Flush()

		// Escreva o sentimento no CSV
		if err := writer.Write([]string{feeling}); err != nil {
			http.Error(w, "Erro ao escrever no arquivo CSV", http.StatusInternalServerError)
			return
		}

		// Responda ao cliente
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Sentimento salvo com sucesso!"))
	})

	fmt.Println("Servidor Go em execução na porta 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
