// cmd/server/main.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type GenerateRequest struct {
    APIToken  string `json:"apiToken"`
    RecipeID  string `json:"recipeId"`
}

func main() {
    port := os.Getenv("PORT") // For Heroku
    if port == "" {
        port = "8080"
    }

    http.HandleFunc("/api/generate", handleGenerate)
    http.Handle("/", http.FileServer(http.Dir("web/static")))
    
    log.Printf("Server starting on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req GenerateRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // TODO: Implement Workato API call
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Successfully received request",
    })
}