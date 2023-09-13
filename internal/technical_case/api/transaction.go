package api

import (
	"github.com/gin-gonic/gin"
	"jump/technical-case/internal/technical_case/models"
	"jump/technical-case/internal/util/connectors"
	"jump/technical-case/internal/util/validators"
	"net/http"
)

func PostTransaction(c *gin.Context) {
	var queryData = validators.TransactionRequest{}
	if err := c.ShouldBind(&queryData); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err, invoice := findInvoice(queryData.InvoiceId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	transaction := models.Transactions{InvoiceId: invoice.Id, Reference: queryData.Reference, Status: "rejected", Amount: roundFloat(queryData.Amount)}
	if roundFloat(invoice.Amount) != roundFloat(queryData.Amount) {
		connectors.Connector.Save(&transaction)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if invoice.Status == "paid" {
		connectors.Connector.Save(&transaction)
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}
	invoice.Status = "paid"
	transaction.Status = "accepted"

	err, user := findUser(invoice.UserId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	user.Balance = roundFloat(user.Balance + invoice.Amount)
	connectors.Connector.Save(&transaction)
	connectors.Connector.Save(&invoice)
	connectors.Connector.Save(&user)
	c.AbortWithStatus(http.StatusNoContent)
}
