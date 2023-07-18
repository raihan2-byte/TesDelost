package handler

import (
	endpointcount "delos/endpointCount"
	"delos/farm"
	"delos/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type farmHandler struct {
	farmService farm.Service
	endpointService endpointcount.StatisticsService
}

func NewFarmHandler(farmService farm.Service, endpointService endpointcount.StatisticsService) *farmHandler {
	return &farmHandler{farmService, endpointService}
}

func (h *farmHandler) CreateFarm(c *gin.Context) {
	var input farm.FarmInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newFarm, err := h.farmService.CreateFarm(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

    userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("POST /farm/create", userAgent)
    if err != nil {
        response := helper.APIresponse(http.StatusUnprocessableEntity, err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
    }

	response := helper.APIresponse(http.StatusCreated, farm.FormatterGetAllFarm(newFarm))
	c.JSON(http.StatusCreated, response)
}

func (h *farmHandler) UpdatedFarm(c *gin.Context) {
	var inputID farm.GetIdFarmInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData farm.UpdateFarmInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newFarm, err := h.farmService.UpdateFarms(inputID, inputData)

	if err != nil {
	
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userAgent := c.GetHeader("User-Agent")


	err = h.endpointService.IncrementCount("PUT /farm/put", userAgent)
    if err != nil {
        response := helper.APIresponse(http.StatusUnprocessableEntity, err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
    }

	response := helper.APIresponse(http.StatusOK, farm.FormatterGetAllFarm(newFarm))
	c.JSON(http.StatusOK, response)

}

func (h *farmHandler) DeletedFarm(c *gin.Context) {
	var input farm.GetIdFarmInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIresponse(http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}


	_, err = h.farmService.DeleteFarm(int(input.ID))
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userAgent := c.GetHeader("User-Agent")


	err = h.endpointService.IncrementCount("Delete /farm/delete", userAgent)
    if err != nil {
        response := helper.APIresponse(http.StatusUnprocessableEntity, err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
    }

	response := helper.APIresponse(http.StatusNoContent, "Your farm has been successfully deleted")
	c.JSON(http.StatusOK, response)
}

func (h *farmHandler) GetFarms(c *gin.Context) {
	input, _ := strconv.Atoi(c.Query("id"))
	// var input farm.GetIdFarmInput

	newFarm, err := h.farmService.GetFarms(input)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userAgent := c.GetHeader("User-Agent")


	err = h.endpointService.IncrementCount("Get /farm/", userAgent)
    if err != nil {
        response := helper.APIresponse(http.StatusUnprocessableEntity, err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
    }

	response := helper.APIresponse(http.StatusOK, farm.FormatterGetFarms(newFarm))
	c.JSON(http.StatusOK, response)
}

func (h *farmHandler) GetOneFarm(c *gin.Context) {
	var input farm.GetIdFarmInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newDel, err := h.farmService.GetFarmById(input.ID)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
		
	}

	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("GetByID /farm/delete", userAgent)
    if err != nil {
        response := helper.APIresponse(http.StatusUnprocessableEntity, err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
    }

	response := helper.APIresponse(http.StatusOK, farm.FormatterGetAllFarm(newDel))
	c.JSON(http.StatusOK, response)
}