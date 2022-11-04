package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/diogocarasco/go-api-test/models"
)

func ImportFile(filename string) ([]models.Feiras, error) {

	logrus.Debug("Proccessing " + filename)

	csvFile, err := os.Open(filename)
	if err != nil {
		logrus.Debug(err.Error())
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		logrus.Debug(err.Error())
		fmt.Println(err)
	}

	var lines []models.Feiras

	for ii, line := range csvLines {

		if ii == 0 {
			continue
		}

		//id, _ = strconv.ParseUint(line[0], 0, 32)
		long, _ := strconv.ParseFloat(line[1], 64)
		lat, _ := strconv.ParseFloat(line[2], 64)
		setcens, _ := strconv.ParseInt(line[3], 0, 64)
		areap, _ := strconv.ParseInt(line[4], 0, 64)
		coddist, _ := strconv.Atoi(line[5])
		codsubpref, _ := strconv.Atoi(line[7])

		emp := models.Feiras{

			LONG:       long,
			LAT:        lat,
			SETCENS:    setcens,
			AREAP:      areap,
			CODDIST:    coddist,
			DISTRITO:   line[6],
			CODSUBPREF: codsubpref,
			SUBPREF:    line[8],
			REGIAO5:    line[9],
			REGIAO8:    line[10],
			NOME_FEIRA: line[11],
			REGISTRO:   line[12],
			LOGRADOURO: line[13],
			NUMERO:     line[14],
			BAIRRO:     line[15],
			REFERENCIA: line[16],
		}

		lines = append(lines, emp)

	}

	e := os.Remove(filename)
	if e != nil {
		logrus.Debug(err.Error())
		log.Fatal(e)
		err = e
	}

	return lines, err

}
