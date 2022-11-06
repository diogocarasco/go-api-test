package models

import (
	"gorm.io/gorm"
)

// FEIRAS data structure
type Feiras struct {
	gorm.Model
	ID         uint64  `json:"id" gorm:"primary_key"`
	LONG       float64 `json:"long"`
	LAT        float64 `json:"lat"`
	SETCENS    int64   `json:"setcens"`
	AREAP      int64   `json:"areap"`
	CODDIST    int     `json:"coddist"`
	DISTRITO   string  `json:"distrito"`
	CODSUBPREF int     `json:"codsubpref"`
	SUBPREF    string  `json:"subpref"`
	REGIAO5    string  `json:"regiao5"`
	REGIAO8    string  `json:"regiao8"`
	NOME_FEIRA string  `json:"nome_feira"`
	REGISTRO   string  `json:"registro"`
	LOGRADOURO string  `json:"logradouro"`
	NUMERO     string  `json:"numero"`
	BAIRRO     string  `json:"bairro"`
	REFERENCIA string  `json:"referencia"`
}

// FEIRAS data structure
type Feira struct {
	ID         uint64  `json:"id"`
	LONG       float64 `json:"long"`
	LAT        float64 `json:"lat"`
	SETCENS    int64   `json:"setcens"`
	AREAP      int64   `json:"areap"`
	CODDIST    int     `json:"coddist"`
	DISTRITO   string  `json:"distrito"`
	CODSUBPREF int     `json:"codsubpref"`
	SUBPREF    string  `json:"subpref"`
	REGIAO5    string  `json:"regiao5"`
	REGIAO8    string  `json:"regiao8"`
	NOME_FEIRA string  `json:"nome_feira"`
	REGISTRO   string  `json:"registro"`
	LOGRADOURO string  `json:"logradouro"`
	NUMERO     string  `json:"numero"`
	BAIRRO     string  `json:"bairro"`
	REFERENCIA string  `json:"referencia"`
}
