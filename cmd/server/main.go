package server

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/hellofreshdevtests/dsaveliev-golang-test/internal/config"
)

func handleRecipes(w http.ResponseWriter, req *http.Request) {
	// Check http method
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "Unsupporthed request method.")
		return
	}

	// Parse query string
	cfg := config.GetConfig()
	skip := parseSkipValue(req, cfg.DefaultSkipValue)
	top := parseTopValue(req, cfg.DefaultTopValue)
	ids := parseIdsValue(req)
	sortByPrepTime := false

	// We must sort recipes by cooking time, if the ids exist
	if len(ids) > 0 {
		sortByPrepTime = true
	} else {
		for i := skip + 1; i <= skip+top; i++ {
			ids = append(ids, i)
		}
	}

	// Context with timeout limit the total request time
	ctx, cancel := context.WithTimeout(
		req.Context(),
		time.Duration(cfg.ServerTimeout)*time.Millisecond,
	)
	defer cancel()

	// Fetch recipes synchronously
	recipes := fetchRecipeList(ctx, ids)

	// Sort recipes by cooking time, if necessary
	if sortByPrepTime {
		recipes.Sort()
	}

	// Marshal response
	response, err := json.Marshal(recipes)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	// Return recipes slice as a JSON response
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Run proxy server with recipes handler
func Run() {
	http.HandleFunc("/recipes", handleRecipes)

	if err := http.ListenAndServe(config.GetConfig().Addr, nil); err != nil {
		log.Fatal(err)
	}
}
