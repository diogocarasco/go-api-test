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

// @BasePath /feiras

// GetFeirasById godoc
// @Summary     Fetches a fair
// @Description get fair record by ID
// @Tags        feira
// @Accept      json
// @Produce     json
// @Param       id  path     int          true "Fair ID"
// @Success     200 {object} models.Feira "Single fair row"
// @Failure     400
// @Router      /feira/{id} [get]
// GetFeiras retrieve registers from FEIRAS based on the passed ID
func GetFeirasById(c *gin.Context) {

	var feiras *models.Feiras

	err := database.DB.Where("id = ?", c.Param("id")).First(&feiras).Error
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		logrus.Info("data:" + "[]")
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": feiras})
	logrus.Debug(feiras)

}

// GetFeiras godoc
// @Summary     Fetch fair data
// @Description get fair records based on the querystring filters and pagination
// @Tags        feira
// @Accept      json
// @Produce     json
// @Param       distrito   query    string false "District"
// @Param       regiao5    query    string false "Region"
// @Param       nome_feira query    string false "Fair name"
// @Param       bairro     query    string false "Neighborhood"
// @Param       page       query    int    false "Page"
// @Success     200        {object} models.Feira
// @Failure     400
// @Failure     500
// @Router      /feira [get]
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
	err = database.DB.Model(&feiras).Where(whereMap).Count(&bookCount).Error
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		logrus.Info("data:" + "400 Bad request - Unable to count rows from FEIRAS")
		return
	}

	booksPerPage := 15
	pageCount := int(math.Ceil(float64(bookCount) / float64(booksPerPage)))
	if pageCount == 0 {
		pageCount = 1
	}
	if page < 1 || page > pageCount {
		c.AbortWithStatus(http.StatusBadRequest)
		logrus.Info("data:" + "400 Bad request - Invalid pages calculation")
		return
	}

	offset := (page - 1) * booksPerPage

	err = database.DB.Debug().Where(whereMap).Offset(offset).Limit(booksPerPage).Find(&feiras).Error
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		logrus.Info("data: 400 Bad request - " + err.Error())
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

}

// InsertFeiras godoc
// @Summary     Insert fair data
// @Description Create a new fair record
// @Tags        feira
// @Accept      json
// @Produce     json
// @Param       request body     models.Feira true "FEIRAS model"
// @Success     200     {string} {id:1}       " JSON with ID value"
// @Failure     400
// @Router      /feira [post]
// InsertFeiras create a new register into FEIRAS
func InsertFeiras(c *gin.Context) {

	var feirasinput models.Feiras

	err := c.ShouldBindJSON(&feirasinput)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		logrus.Info("error:" + err.Error())
		return
	}

	database.DB.Create(&feirasinput)
	c.JSON(http.StatusOK, gin.H{"id": feirasinput.ID})
	logrus.Debug(feirasinput.ID)

}

// DeleteFeiras godoc
// @Summary     Delete a fair
// @Description Delete a single fair record based on ID
// @Tags        feira
// @Accept      json
// @Produce     json
// @Param       id  path     int            true "Fair ID"
// @Success     200 {string} teste "Success"
// @Failure     400
// @Router      /feira/{id} [delete]
// DeleteFeiras remove a FEIRAS register by ID
func DeleteFeiras(c *gin.Context) {

	var feiras models.Feiras

	err := database.DB.Where("id = ?", c.Param("id")).First(&feiras).Error
	if err != nil {
		logrus.Debug(gin.H{"data": err.Error()})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	database.DB.Delete(&feiras)

	c.JSON(http.StatusOK, gin.H{"data": true})
	logrus.Debug(gin.H{"data": true})

}

// UpdateFeiras godoc
// @Summary     Update fair data
// @Description Create a new fair record
// @Tags        feira
// @Accept      json
// @Produce     json
// @Param       distrito   query    string false "District"
// @Param       regiao5    query    string false "Region"
// @Param       nome_feira query    string false "Fair name"
// @Param       bairro     query    string false "Neighborhood"
// @Param       page                       query int   false "Page"
// @Success     200        {string} {id:1} " JSON with ID value"
// @Failure     400
// @Router      /feira [patch]
// UpdateFeiras update row from FEIRAS
func UpdateFeiras(c *gin.Context) {

	var feirasinput models.Feiras
	err := c.ShouldBindJSON(&feirasinput)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		logrus.Info("error:" + err.Error())
		return
	}

	var feiras models.Feiras
	err = database.DB.Where("id = ?", c.Param("id")).First(&feiras).Error
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
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

// ImportFeiras godoc
// @Summary     Import fair data
// @Description Import a CSV containing fair data
// @Tags        feira
// @Accept      mpfd
// @Produce     json
// @Param       file               formData                        string                     true "File"
// @Success     200     {string} {data:true}       " File uploaded"
// @Failure     400
// @Router      /feira/upload [post]
// ImportFeiras imports the CSV from the request file to the database
func ImportFeiras(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		logrus.Debug(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, "data/import/"+filename); err != nil {
		logrus.Debug(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	content, errors := utils.ImportFile("data/import/" + filename)
	if errors != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		err := database.DB.Create(&content)
		if err.Error != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			logrus.Debug(err.Error)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}

}
