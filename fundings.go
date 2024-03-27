package routes

import (
	"crowdfunding/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getFundings(context *gin.Context) {
	fundings, err := models.GetAllFundings()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch models."})
		return
	}
	if len(fundings) < 1 {
		context.JSON(http.StatusOK, gin.H{"message": "There are no fundings available."})
		return
	}
	context.JSON(http.StatusOK, fundings)
}

func getFunding(context *gin.Context) {
	fundingID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse funding id."})
		return
	}
	funding, err := models.GetFundingByID(fundingID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch funding."})
		return
	}
	context.JSON(http.StatusOK, funding)
}

func createFunding(context *gin.Context) {
	var funding models.Funding
	err := context.ShouldBindJSON(&funding)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request."})
		return
	}

	userID := int64(1)
	funding.UserID = userID

	err = funding.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create funding" + err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Funding created"})
}

func deleteFunding(context *gin.Context) {
	fundingID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse funding id."})
		return
	}
	funding, err := models.GetFundingByID(fundingID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch funding."})
		return
	}
	err = funding.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete funding."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": gin.H{"message": "Funding deleted succesfully!"}})
}
