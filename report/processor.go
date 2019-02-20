package report

import (
	"fmt"
)

type Report struct {
	OrderID           string    `json:"order_id"`
	SelectedHub int `json:"selected_hub"`
	HubWithLowestEta int `json:"hub_with_lowest_eta"`
	SelectedDelivery  *Delivery `json:"selected_delivery"`
	LowestEtaDelivery *Delivery `json:"delivery_with_lowest_eta"`
}

func Process(orders []Order) []Report {

	var reports []Report

	fmt.Println(len(orders))
	for i := 0; i < len(orders); i++ {
		order := orders[i]
		deliveryOptionsLength := len(order.Options.DeliveryOptions)
		if deliveryOptionsLength == 0 {
			fmt.Println("No delivery option in ", i, "line!")
			continue
		}

		if deliveryOptionsLength == 1 {
			fmt.Println("Only 1 delivery option in ", i, "line!")
			continue
		}

		lowestEtaIndex := 0
		lowestEta := order.Options.DeliveryOptions[lowestEtaIndex].ETA.Low

		var selectedDelivery *Delivery

		for j := 0; j < len(order.Options.DeliveryOptions); j++ {
			delivery := order.Options.DeliveryOptions[j]

			if j == 0 {
				selectedDelivery = order.Options.DeliveryOptions[j]
			}
			if delivery.ETA.Low < lowestEta && delivery.ETA.Method != "lower_bound"{
				lowestEtaIndex = j
				lowestEta = delivery.ETA.Low
			}
		}

		if lowestEtaIndex > 0 {
			lovestEtaDelivery := order.Options.DeliveryOptions[lowestEtaIndex]

			report := Report{
				order.OrderId,
				selectedDelivery.Vehicle,
				lovestEtaDelivery.Vehicle,
				selectedDelivery,
				lovestEtaDelivery,
			}

			reports = append(reports, report)
		}
	}

	return reports
}
