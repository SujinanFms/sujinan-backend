//  /sujinan-backend/handlers/recommendations.go
package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/SujinanFms/sujinan-backend/models"
	"github.com/SujinanFms/sujinan-backend/database"
)

func HandleRecommendations(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		rows, err := database.DB.Query("SELECT title FROM recommendations")
		if err != nil {
			http.Error(w, "Failed to fetch recommendations", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var recommendations []string
		for rows.Next() {
			var title string
			if err := rows.Scan(&title); err == nil {
				recommendations = append(recommendations, title)
			}
		}
		json.NewEncoder(w).Encode(recommendations)
		return
	}

	if r.Method == http.MethodPost {
		var rec models.Recommendation
		if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		_, err := database.DB.Exec("INSERT INTO recommendations (title) VALUES ($1)", rec.Title)
		if err != nil {
			http.Error(w, "Failed to save recommendation", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Recommendation saved",
		})
	}
}

// package handlers

// import (
// 	"encoding/json"
// 	"net/http"
// )

// func HandleRecommendations(w http.ResponseWriter, r *http.Request) {
// 	recommendations := []string{"Item A", "Item B"}
// 	json.NewEncoder(w).Encode(recommendations)
// }
