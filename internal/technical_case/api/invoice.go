package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jump/technical-case/internal/technical_case/models"
	"jump/technical-case/internal/utils/connectors"
	"jump/technical-case/internal/utils/validators"
	"net/http"
)

func PostInvoice(c *gin.Context) {
	var queryData = validators.InvoiceRequest{}
	if err := c.ShouldBind(&queryData); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err, user := findUser(queryData.UserId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	invoice := models.Invoices{UserId: user.Id, Status: "pending", Label: queryData.Label, Amount: roundFloat(queryData.Amount)}
	inserted := connectors.Connector.Create(&invoice)
	if inserted.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.AbortWithStatus(http.StatusNoContent)
}

func findInvoice(id int) (error, *models.Invoices) {
	var invoice models.Invoices
	connectors.Connector.First(&invoice, id)
	if invoice.Id == 0 {
		return gorm.ErrRecordNotFound, &invoice
	}
	return nil, &invoice
}
