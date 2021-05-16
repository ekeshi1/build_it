package domain

import "time"

const (
	POStatusCreated           = "CREATED"
	POStatusSent              = "SENT"
	PODeliveryStatusCreated   = "CREATED"
	PODeliveryStatusDelivered = "DELIVERED"
	InvStatusPaid             = "PAID"
	InvStatusCreated          = "CREATED"
	InvStatusApproved         = "APPROVED"
	PHApproved                = "APPROVED"
	PHCreated                 = "CREATED"
)

type PlantHire struct {
	Id                 int64     `json:"id"`
	PlantId            int64     `json:"plantId"`
	ConstructionSiteId int64     `json:"constructionSiteId"`
	SupplierId         int64     `json:"supplierId"`
	SiteEngineerId     int64     `json:"siteEngineerId"`
	PlantArrivalDate   string    `json:"plantArrivalDate"`
	PlantReturnDate    string    `json:"plantReturnDate"`
	PlantDailyPrice    float64   `json:"plantDailyPrice"`
	PlantTotalPrice    float64   `json:"plantTotalPrice"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	Status             string    `json:"status"`
}

type PurchaseOrder struct {
	Id             int64     `json:"id"`
	PlantHireId    int64     `json:"plantHireId"`
	Description    string    `json:"description"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Creator        string    `json:"creator"`
	DeliveryStatus string    `json:"deliveryStatus"`
	Status         string    `json:"status"`
}

type Invoice struct {
	Id              int64     `json:"id"`
	PurchaseOrderId int64     `json:"purchaseOrderId"`
	CreatedAt       time.Time `json:"createdDate"`
	UpdatedAt       time.Time `json:"updatedDate"`
	//PaymentDate     time.Time `json:"paymentDate"`
	PaymentStatus string `json:"paymentStatus"`
}

type PlantHireExtensionDTO struct {
	PlantReturnDate string `json:"plantReturnDate"`
}

type RemittanceAdviceDTO struct {
	CompanyName   string    `json:"companyName"`
	InvoiceNumber int64     `json:"invoiceNumber"`
	PaymentDate   time.Time `json:"paymentDate"`
	Amount        float64   `json:"amount"`
}
