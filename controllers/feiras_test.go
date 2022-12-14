package controllers_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	configs "github.com/diogocarasco/go-api-test/config"
	"github.com/diogocarasco/go-api-test/database"
	"github.com/gin-gonic/gin"
	"github.com/instana/testify/assert"
)

func init() {
	database.InitializeDB("mock")

}

var w = httptest.NewRecorder()
var C, _ = gin.CreateTestContext(w)

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

	database.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "feiras" WHERE "feiras"."deleted_at" IS NULL`)).WillReturnRows(sqlmock.NewRows([]string{"800"}))

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
	req, err := http.NewRequest("GET", "/feira/4", nil)
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
	req, err := http.NewRequest("PATCH", "/feira/4", bytes.NewBuffer(exp))
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

	//id := rand.Intn(100)

	req, err := http.NewRequest("DELETE", "/feira/4", nil)
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
