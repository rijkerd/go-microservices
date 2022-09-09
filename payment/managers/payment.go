package managers

import "payment/models"

func GetPayments() []models.Payment {
	payments := []models.Payment{
		{ID: 1, Ref: "RF091ZSHS", Amount: 1000.0, Currency: "RUB"},
		{ID: 2, Ref: "JK123DFZ", Amount: 500.0, Currency: "RUB"},
		{ID: 3, Ref: "KL09012112", Amount: 600.0, Currency: "RUB"},
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
