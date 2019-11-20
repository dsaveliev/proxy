package types

import (
	"regexp"
	"sort"
	"strconv"
)

// Ingredient struct represents recipe's component
type Ingredient struct {
	Name      string `json:"name"`
	ImageLink string `json:"imageLink"`
}

// Recipe struct reveals different characteristics of the recipe
type Recipe struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Headline    string       `json:"headline"`
	Description string       `json:"description"`
	Difficulty  int          `json:"difficulty"`
	PrepTime    string       `json:"prepTime"`
	ImageLink   string       `json:"imageLink"`
	Ingredients []Ingredient `json:"ingredients"`
}

// RecipeList stands for a recipe collection
type RecipeList []*Recipe

// prepTimeRegexp describes how to get to the preparing time (lower bound for the range)
var prepTimeRegexp = regexp.MustCompile(`^PT(\d+).*M$`)

// Sort method orders a slice by PrepTime
func (l RecipeList) Sort() {
	sort.Slice(l, func(i, j int) bool {
		// Parse PrepTime
		si := prepTimeRegexp.FindStringSubmatch(l[i].PrepTime)
		sj := prepTimeRegexp.FindStringSubmatch(l[j].PrepTime)

		// Push recipe without PrepTime to the end of the slice
		if len(si) < 2 {
			return false
		}
		if len(sj) < 2 {
			return true
		}

		// Convert to minutes and compare the values
		ti, err := strconv.Atoi(si[1])
		if err != nil {
			return false
		}
		tj, err := strconv.Atoi(sj[1])
		if err != nil {
			return true
		}
		return ti < tj
	})
}
