package models

type Reserve struct {
	UserId     string  `json:"user_id"`
	ServiceId  string  `json:"service_id"`
	PurchaseId string  `json:"purchase_id"`
	Price      float64 `json:"price"`
}

type UserBalance struct {
	UserId  string  `json:"user_id"`
	Balance float64 `json:"balance"`
}

type AccrualMoney struct {
	UserId string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

type Report struct {
	UserId     string  `json:"user_id"`
	ServiceId  string  `json:"service_id"`
	PurchaseId string  `json:"purchase_id"`
	Price      float64 `json:"price"`
}
