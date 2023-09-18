package validators

type InvoiceRequest struct {
	UserId int     `form:"user_id" json:"user_id" binding:"required"`
	Amount float64 `form:"amount" json:"amount" binding:"required"`
	Label  string  `form:"label" json:"label" binding:"required"`
}

type TransactionRequest struct {
	InvoiceId int     `form:"invoice_id" json:"invoice_id" binding:"required"`
	Amount    float64 `form:"amount" json:"amount" binding:"required"`
	Reference string  `form:"reference" json:"reference" binding:"required"`
}
