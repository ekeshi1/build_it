package domain

import "time"

const (
	POStatusCreated           = "CREATED"
	POStatusSent              = "SENT"
	PODeliveryStatusCreated   = "CREATED"
	PODeliveryStatusDelivered = "DELIVERED"
)

type PlantHire struct {
	Id                 int64     `json:"id"`
	PlantId            int64     `json:"plantId"`
	ConstructionSiteId int64     `json:"constructionSiteId"`
	SupplierId         int64     `json:"supplierId"`
	SiteEngineerId     int64     `json:"siteEngineerId"`
	PlantArrivalDate   string    `json:"plantArrivalDate"`
	PlantReturnDate    string    `json:"plantReturnDate"`
	PlantTotalPrice    float64   `json:"plantTotalPrice"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	Status             string    `json:"status"`
}

type PurchaseOrder struct {
	Id              int64     `json:"id"`
	PlantHireId     int64     `json:"plantHireId"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	Creator         string    `json:"creator"`
	DeliveryAddress string    `json:"deliveryAddress"`
	DeliveryStatus  string    `json:"deliveryStatus"`
	Status          string    `json:"status"`
}

type Invoice struct {
	Id               int64     `json:"id"`
	PurchaseOrderId  int64     `json:"purchaseOrderId"`
	CreatedAt        time.Time `json:"createdDate"`
	UpdatedAt        time.Time `json:"updatedDate"`
	LastReminderDate time.Time `json:"lastReminderDate"`
	PaymentStatus    string    `json:"paymentStatus"`
}
