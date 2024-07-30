package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizhirasawa/Go-api.git/schemas"
)

//BasePath

// @summary Create Opening
// @description Create a new Opening
// @ tags Openings
// @accept json
// @produce json
// @param request body CreateOpeningRequest true "Request Body"
// @Success 200 {object} CreateOpeningResponse
// Failure 400 {object} ErrorResponse
// Failure 400 {object} ErrorResponse
// @router /openings [post]
func PostOpeningHandler(c *gin.Context) {
	request := CreateOpeningRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("Validation error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.ErrorF("Error creating opening: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Error creating opening on database")
		return
	}
	sendSuccess(c, "Created", opening)
}
