package handler

import (
	endpointcount "delos/endpointCount"
	"delos/helper"
	"delos/pounds"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type poundsHandler struct {
	poundsService pounds.Service
	endpointService endpointcount.StatisticsService

}

func NewPoundsHandler(poundsService pounds.Service, endpointService endpointcount.StatisticsService) *poundsHandler {
	return &poundsHandler{poundsService, endpointService}
}

func (h *poundsHandler) CreatePounds(c *gin.Context) {
	var input pounds.PoundsInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newFarm, err := h.poundsService.CreatePounds(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("Post /pounds/", userAgent)
    if err != nil {
        response := helper.APIresponse(http.StatusUnprocessableEntity, err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
    }

	response := helper.APIresponse(http.StatusCreated, pounds.FormatterPounds(newFarm))
	c.JSON(http.StatusCreated, response)
}

func (h *poundsHandler) UpdatedPounds(c *gin.Context) {
	var inputID pounds.GetIdPoundsInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData pounds.UpdatePoundsInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	farm, err := h.poundsService.UpdatePounds(inputID, inputData)

	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userAgent := c.GetHeader("User-Agent")
	
	err = h.endpointService.IncrementCount("PUT /pounds/", userAgent)
    if err != nil {
        response := helper.APIresponse(http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
    }


	response := helper.APIresponse(http.StatusOK, pounds.FormatterGetPound(farm))
	c.JSON(http.StatusOK, response)

}

func (h *poundsHandler) DeletedPounds(c *gin.Context) {
	var input pounds.GetIdPoundsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIresponse(http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}


	_, err = h.poundsService.DeletePounds(int(input.ID))
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("DELETE /pounds/", userAgent)
    if err != nil {
        response := helper.APIresponse(http.StatusUnprocessableEntity, err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
    }

	response := helper.APIresponse(http.StatusNoContent, "your pounds has been successfully deleted")
	c.JSON(http.StatusOK, response)
}

func (h *poundsHandler) GetPounds(c *gin.Context) {
	input, _ := strconv.Atoi(c.Query("id"))
	// var input farm.GetIdFarmInput

	newPounds, err := h.poundsService.GetPounds(input)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("GET /pounds/", userAgent)
    if err != nil {
        response := helper.APIresponse(http.StatusUnprocessableEntity, err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
    }

	response := helper.APIresponse(http.StatusOK, pounds.FormatterGetPounds(newPounds))
	c.JSON(http.StatusOK, response)
}

func (h *poundsHandler) GetOnePounds(c *gin.Context) {
	var input pounds.GetIdPoundsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newDel, err := h.poundsService.GetPoundsById(input.ID)
	if err != nil {

		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
		
	}

	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("GET /poundsByID/", userAgent)
    if err != nil {
        response := helper.APIresponse(http.StatusUnprocessableEntity, err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
    }

	response := helper.APIresponse(http.StatusOK, pounds.FormatterGetPound(newDel))
	c.JSON(http.StatusOK, response)
}