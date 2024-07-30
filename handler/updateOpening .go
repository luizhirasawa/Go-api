package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizhirasawa/Go-api.git/schemas"
)

func UpdateOpeningHandler(c *gin.Context) {
	request := UpdateOpeningRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("Validation error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, "Missing id parameter")
		return
	}
	opening := schemas.Opening{}
	if err := db.Where("id = ?", id).First(&opening).Error; err != nil {
		logger.ErrorF("Error fetching opening: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Error fetching opening from database")
		return
	}
	if len(request.Role) > 0 {
		opening.Role = request.Role
	}
	if len(request.Company) > 0 {
		opening.Company = request.Company
	}
	if len(request.Location) > 0 {
		opening.Location = request.Location
	}
	if request.Remote != nil && request.Remote != &opening.Remote {
		opening.Remote = *request.Remote
	}
	if len(request.Link) > 0 {
		opening.Link = request.Link
	}
	if request.Salary > 0 && request.Salary != opening.Salary {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.ErrorF("Error updating opening: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Error updating opening on database")
		return
	}
	sendSuccess(c, "successfully updated", opening)
}
