package controllers_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	configs "go-api-test/config"
	"go-api-test/database"

	"github.com/instana/testify/assert"
)

func init() {
	database.InitializeDB()
}

func TestHome(t *testing.T) {

	router := configs.GetServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "\"Hello!!\"", w.Body.String())

}
func TestGetFeiras(t *testing.T) {

	router := configs.GetServer()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/feira", nil)
	if err != nil {
		fmt.Print(err.Error())
		t.Error(err)
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotEqual(t, "\"data:[]\"", w.Body.String())
}

func TestGetFeirasQueryString(t *testing.T) {

	router := configs.GetServer()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/feira?bairro=vl formosa", nil)
	if err != nil {
		fmt.Print(err.Error())
		t.Error(err)
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotEqual(t, "\"data:[]\"", w.Body.String())
}

func TestGetFeirasById(t *testing.T) {

	router := configs.GetServer()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/feira/1", nil)
	if err != nil {
		fmt.Print(err.Error())
		t.Error(err)
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotEqual(t, "\"data:[]\"", w.Body.String())

}

func TestUpdateFeiras(t *testing.T) {

	exp := []byte(`{"long": -46550166}`)

	router := configs.GetServer()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", "/feira/1", bytes.NewBuffer(exp))
	if err != nil {
		fmt.Print(err.Error())
		t.Error(err)
	}
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestDeleteFeiras(t *testing.T) {

	router := configs.GetServer()
	w := httptest.NewRecorder()

	id := rand.Intn(100)

	req, err := http.NewRequest("DELETE", "/feira/"+strconv.Itoa(id), nil)
	if err != nil {
		fmt.Print(err.Error())
		t.Error(err)
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}

func TestInsertFeiras(t *testing.T) {

	exp := []byte(`{"long": -46550164,"lat": -23558733,"setcens": 355030885000091,"areap": 3550308005040,"coddist": 87,"distrito": "VILA FORMOSA","codsubpref": 26,"subpref": "ARICANDUVA-FORMOSA-CARRAO","regiao5": "Leste","regiao8": "Leste 1","nome_feira": "VILA FORMOSA","registro": "4041-0","logradouro": "RUA MARAGOJIPE","numero": "S/N","bairro": "VL FORMOSA","referencia": "TV RUA PRETORIA"	}`)

	router := configs.GetServer()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/feira", bytes.NewBuffer(exp))
	if err != nil {
		fmt.Print(err.Error())
		t.Error(err)
	}
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}
