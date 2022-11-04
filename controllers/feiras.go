package controllers

import (
	"fmt"
	"math"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/diogocarasco/go-api-test/database"
	"github.com/diogocarasco/go-api-test/models"
	"github.com/diogocarasco/go-api-test/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetFeiras retrieve registers from FEIRAS based on the passed ID
func GetFeirasById(c *gin.Context) {

	var feiras *models.Feiras

	err := database.DB.Where("id = ?", c.Param("id")).First(&feiras).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "{}"})
		logrus.Info("data:" + "[]")
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": feiras})
	logrus.Debug(feiras)

}

// GetFeiras retrieve registers from FEIRAS based on the passed query string filter
func GetFeiras(c *gin.Context) {

	var feiras []models.Feiras
	whereMap := make(map[string]interface{})

	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	distrito := c.Query("distrito")
	if distrito != "" {
		whereMap["distrito"] = distrito
	}

	regiao5 := c.Query("regiao5")
	if regiao5 != "" {
		whereMap["regiao5"] = regiao5
	}

	nome_feira := c.Query("nome_feira")
	if nome_feira != "" {
		whereMap["nome_feira"] = nome_feira
	}

	bairro := c.Query("bairro")
	if bairro != "" {
		whereMap["bairro"] = bairro
	}

	var bookCount int64
	err = database.DB.Debug().Model(&feiras).Where(whereMap).Count(&bookCount).Error
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	booksPerPage := 15
	pageCount := int(math.Ceil(float64(bookCount) / float64(booksPerPage)))
	if pageCount == 0 {
		pageCount = 1
	}
	if page < 1 || page > pageCount {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	offset := (page - 1) * booksPerPage

	err = database.DB.Debug().Where(whereMap).Offset(offset).Limit(booksPerPage).Find(&feiras).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "{}"})
		logrus.Info("data:" + "[]")
		return
	}

	var prevPage, nextPage string
	if page > 1 {
		prevPage = fmt.Sprintf("%d", page-1)
	}
	if page < pageCount {
		nextPage = fmt.Sprintf("%d", page+1)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      feiras,
		"pageCount": pageCount,
		"page":      page,
		"prevPage":  prevPage,
		"nextPage":  nextPage,
	})
	logrus.Debug(feiras)

}

// InsertFeiras create a new register into FEIRAS
func InsertFeiras(c *gin.Context) {

	var feirasinput models.Feiras

	err := c.ShouldBindJSON(&feirasinput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logrus.Info("error:" + err.Error())
		return
	}

	database.DB.Create(&feirasinput)
	c.JSON(http.StatusOK, gin.H{"id": feirasinput.ID})
	logrus.Debug(feirasinput.ID)

}

// DeleteFeiras remove a FEIRAS register by ID
func DeleteFeiras(c *gin.Context) {

	var feiras models.Feiras

	err := database.DB.Where("id = ?", c.Param("id")).First(&feiras).Error
	if err != nil {
		logrus.Debug(gin.H{"data": err.Error()})
		c.JSON(http.StatusBadRequest, gin.H{"data": "{}"})
		return
	}

	database.DB.Delete(&feiras)

	c.JSON(http.StatusOK, gin.H{"data": true})
	logrus.Debug(gin.H{"data": true})

}

// UpdateFeiras update row from FEIRAS
func UpdateFeiras(c *gin.Context) {

	var feirasinput models.Feiras
	err := c.ShouldBindJSON(&feirasinput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logrus.Info("error:" + err.Error())
		return
	}

	var feiras models.Feiras
	err = database.DB.Where("id = ?", c.Param("id")).First(&feiras).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "{}"})
		logrus.Info("error:" + err.Error())
		return
	}

	upderr := database.DB.Model(&feiras).Omit("id").Updates(feirasinput)
	if upderr != nil {
		logrus.Debug(upderr.Error)
	}

	c.JSON(http.StatusOK, gin.H{"data": feiras})
	logrus.Debug(gin.H{"data": feiras})

}

func ImportFeiras(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		logrus.Debug(err.Error())
		c.String(http.StatusBadRequest, fmt.Sprintf("File %s processed unsuccessfully.", file.Filename))
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, "data/import/"+filename); err != nil {
		logrus.Debug(err.Error())
		c.String(http.StatusBadRequest, fmt.Sprintf("File %s processed unsuccessfully.", file.Filename))
		return
	}

	content, errors := utils.ImportFile("data/import/" + filename)
	if errors != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("File %s processed unsuccessfully.", file.Filename))
	} else {
		err := database.DB.Create(&content)
		if err.Error != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("File %s processed unsuccessfully.", file.Filename))
			logrus.Debug(err.Error)
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("File %s processed successfully.", file.Filename))
	}

}
