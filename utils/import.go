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

		emp := models.Feiras{

			DISTRITO: line[6],

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

		emp.ID, _ = strconv.ParseUint(line[0], 0, 32)
		emp.LONG, _ = strconv.ParseFloat(line[1], 64)
		emp.LAT, _ = strconv.ParseFloat(line[2], 64)
		emp.SETCENS, _ = strconv.ParseInt(line[3], 0, 64)
		emp.AREAP, _ = strconv.ParseInt(line[4], 0, 64)
		emp.CODDIST, _ = strconv.Atoi(line[5])
		emp.CODSUBPREF, _ = strconv.Atoi(line[7])

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
