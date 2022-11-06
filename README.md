# GO API TEST  [![Go Report Card](https://goreportcard.com/badge/github.com/diogocarasco/go-api-test)](https://goreportcard.com/report/github.com/diogocarasco/go-api-test) [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/diogocarasco/go-api-test) 
<img align="right" width="159px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">
A GIN-GONIC REST API.


The requirement to use this repo is `docker` and `docker-compose >=v1.27`, if you need
to install this plugin, follow these steps:

- install [Docker](https://docs.docker.com/engine/install/ubuntu/)
- install [Docker compose](https://docs.docker.com/compose/install/)


## How to run
```
$ docker compose up
```

## Run tests
```
$ go test ./...
```

## [Open API](https://github.com/swaggo/swag) documentation
http://localhost:8080/swagger/index.html

## Examples

 - GET all fair results
```
$ curl --location --request GET 'http://localhost:8080/feira'
```

- GET fair using filters and pagination
```
$ curl --location --request GET 'http://localhost:8080/feira?distrito=PERDIZES&page=1' \
       --header 'Content-Type: multipart/form-data'
```

- POST to create a new fair record 
```
$ curl --location --request POST 'http://localhost:8080/feira' \
       --header 'Content-Type: application/json' \
       --data-raw '{
           "long" : -46674080,
           "lat" : -23486660,
           "setcens" : 355030850000100,
           "areap" : 3550308005060,
           "coddist" : 50,
           "distrito" : "LIMAO",
           "codsubpref" : 4,
           "subpref" : "CASA VERDE-CACHOEIRINHA",
           "regiao5" : "Norte",
           "regiao8" : "Norte 1",
           "nome_feira" : "SANTA MARIA",
           "registro" : "3077-5",
           "logradouro" : "RUA TOMAZ ANTONIO VILANE",
           "numero" : "417.000000",
           "bairro" : "VL STA MARIA",
           "referencia" : "TV AV DEP EMILIO CARLOS"
         }'
```

- PATCH to partially update a fair record
```
$ curl --location --request PATCH 'http://localhost:8080/feira/1' \
       --header 'Content-Type: application/json' \
       --data-raw '{
           "distrito" : "BAIRRO DO LIMAO",
         }'
```

- DELETE a fair record
```
$ curl --location --request DELETE 'http://localhost:8080/feira/1'
```



- POST file to upload
```
$ curl --location --request POST 'http://localhost:8080/feira/upload' \
       --form 'file=@"/home/diogo/DEINFO_AB_FEIRASLIVRES_2014.csv"'
```
