package types

import "testing"

func TestSort(t *testing.T) {
	recipes := RecipeList{
		&Recipe{ID: "1", PrepTime: "PT25M"},
		&Recipe{ID: "2", PrepTime: "PT5M"},
		&Recipe{ID: "3", PrepTime: "PT2-3M"},
		&Recipe{ID: "4", PrepTime: ""},
		&Recipe{ID: "5", PrepTime: "PT120M"},
	}
	expected := []string{"3", "2", "1", "5", "4"}

	recipes.Sort()

	for i, r := range recipes {
		if r.ID != expected[i] {
			t.Errorf("Recipes order was incorrect, got: %s, want: %s.", r.ID, expected[i])
		}
	}
}
