package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jump/technical-case/internal/technical_case/models"
	"jump/technical-case/internal/utils/connectors"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var user []models.Users
	result := connectors.Connector.Find(&user)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, user)
}

func findUser(id int) (error, *models.Users) {
	var user models.Users
	connectors.Connector.First(&user, id)
	if user.Id == 0 {
		return gorm.ErrRecordNotFound, &user
	}
	return nil, &user
}
