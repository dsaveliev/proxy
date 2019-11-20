package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/hellofreshdevtests/dsaveliev-golang-test/internal/client"
	"github.com/hellofreshdevtests/dsaveliev-golang-test/internal/config"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func initClient(server *httptest.Server) {
	client.SetClient(server.Client())
}

func initConfig(server *httptest.Server) {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}

	cfg := &config.Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Fatal("Error load env variables: " + err.Error())
	}
	cfg.ProxyEndpoint = server.URL
	config.SetConfig(cfg)
}

func readRecipeJSON(id int) []byte {
	recipe, err := ioutil.ReadFile(fmt.Sprintf("../../test/recipe_%d.json", id))
	if err != nil {
		log.Fatal("Error reading recipe's stub file: " + err.Error())
	}
	return recipe
}

func TestFetchRecipe(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write(readRecipeJSON(1))
	}))
	defer server.Close()

	initClient(server)
	initConfig(server)

	r, err := fetchRecipe(1)
	if err != nil {
		t.Errorf("Response was incorrect, got %#v %#v.", r, err)
	}
	if r.ID != "1" {
		t.Errorf("Recipe's ID was incorrect, got %s.", r.Name)
	}
	if r.Name != "Parmesan-Crusted Pork Tenderloin" {
		t.Errorf("Recipe's name was incorrect, got %s.", r.Name)
	}
}

func TestFetchRecipeList(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() == "/1" {
			rw.Write(readRecipeJSON(1))
		} else if req.URL.String() == "/2" {
			rw.Write(readRecipeJSON(2))
		}
	}))
	defer server.Close()

	initClient(server)
	initConfig(server)

	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancel()

	rs := fetchRecipeList(ctx, []int{1, 2})
	rs.Sort()

	if len(rs) != 2 {
		t.Errorf("Response was incorrect, got %#v.", rs)
	}
	if rs[0].ID != "1" {
		t.Errorf("Recipe's ID was incorrect, got %s.", rs[0].ID)
	}
	if rs[0].Name != "Parmesan-Crusted Pork Tenderloin" {
		t.Errorf("Recipe's name was incorrect, got %s.", rs[0].Name)
	}
	if rs[1].ID != "2" {
		t.Errorf("Recipe's ID was incorrect, got %s.", rs[1].ID)
	}
	if rs[1].Name != "Melty Monterey Jack Burgers" {
		t.Errorf("Recipe's name was incorrect, got %s.", rs[1].Name)
	}
}
