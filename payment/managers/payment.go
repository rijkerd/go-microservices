package managers

import "go-microservices/payment/models"

func GetPayments() []models.Payment {
	payments := []models.Payment{
		{ID: 1, Ref: "test1", Amount: 65.4, Currency: "GBP"},
		{ID: 2, Ref: "test2", Amount: 50.9, Currency: "GBP"},
		{ID: 3, Ref: "test3", Amount: 60.7, Currency: "GBP"},
	}
	return payments
}

func GetPaymentByID(id int) models.Payment {

	for _, item := range GetPayments() {
		if item.ID == id {
			return item
		}
	}

	return models.Payment{}
}
