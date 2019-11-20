package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hellofreshdevtests/dsaveliev-golang-test/internal/client"
	"github.com/hellofreshdevtests/dsaveliev-golang-test/internal/config"
	"github.com/hellofreshdevtests/dsaveliev-golang-test/internal/types"
)

func fetchRecipe(id int) (*types.Recipe, error) {
	cfg := config.GetConfig()
	client := client.GetClient()

	url := fmt.Sprintf("%s/%d", cfg.ProxyEndpoint, id)
	response, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Recipe #%d: response error: %w", id, err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Recipe #%d: wrong response status: %s", id, response.Status)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Recipe #%d: error reading response body: %w", id, err)
	}

	recipe := &types.Recipe{}
	err = json.Unmarshal(data, recipe)
	if err != nil {
		return nil, fmt.Errorf("Recipe #%d: error unmarshal response body: %w", id, err)
	}

	return recipe, nil
}

func fetchRecipeList(ctx context.Context, ids []int) types.RecipeList {
	result := types.RecipeList{}
	cfg := config.GetConfig()

	// Setup a semaphore channel in order to limit concurrency.
	// This helps to avoid resources exhaustion (e.g. file descriptors)
	// and improve performance due to decreasing context switching
	semaphoreCh := make(chan struct{}, cfg.ConcurrencyLimit)
	successCh := make(chan *types.Recipe)
	errorCh := make(chan error)

	defer func() {
		close(successCh)
		close(errorCh)
		close(semaphoreCh)
	}()

	for _, id := range ids {
		// In the case of context timeout,
		// all goroutines should stop their work
		go func(ctx context.Context, recipeID int) {
			select {
			case <-ctx.Done():
				errorCh <- errors.New("Timeout")
				return
			case semaphoreCh <- struct{}{}:
				recipe, err := fetchRecipe(recipeID)
				if err == nil {
					successCh <- recipe
				} else {
					errorCh <- err
				}
				<-semaphoreCh
			}
		}(ctx, id)
	}

	for range ids {
		select {
		case recipe := <-successCh:
			result = append(result, recipe)
		case err := <-errorCh:
			// Ignore timeout errors to avoid log flood
			if err.Error() != "Timeout" {
				log.Println("Response Error : ", err)
			}
		}
	}

	return result
}
