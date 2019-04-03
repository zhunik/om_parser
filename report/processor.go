package report

import (
	"fmt"
)

type Report struct {
	OrderID                  string    `json:"order_id"`
	UserID                   int       `json:"user_id"`
	SelectedHub              int       `json:"selected_hub"`
	HubWithLowestEta         int       `json:"hub_with_lowest_eta"`
	HubWithShortestDistance  int       `json:"hub_with_shortest_distance"`
	SelectedDelivery         *Delivery `json:"selected_delivery"`
	LowestEtaDelivery        *Delivery `json:"delivery_with_lowest_eta"`
	ShortestDistanceDelivery *Delivery `json:"delivery_with_lowest_eta"`
	shouldBeSortedBy         string    `json:"should_be_sorted_by"`
	sortedByDriveTime        bool      `json:"sorted_by_drive_time"`
	sortedByDistance         bool      `json:"sorted_by_distance"`
}

func Process(orders []Order) []Report {

	var reports []Report

	users := []int{
		44926, 9568, 28133, 25246, 37587, 12402, 33666, 44129, 26998, 40025, 40635, 44798,
		35271, 40131, 44530, 41050, 44173, 44927, 42805, 23442, 9418, 36970, 44172, 30605, 43271, 29279,
		31839, 35929, 511, 8077, 39374, 44180, 42194, 9153, 27381, 46367, 35899, 36117, 39299, 21782, 36400,
		36767, 36571, 25056, 8180, 33395, 44889, 38023, 40072, 40213,
	}

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

		lowestDriveTimeIndex := 0
		lowestDriveTime := order.Options.DeliveryOptions[0].DriveTime

		shortestDistanceIndex := 0
		shortestDistance := order.Options.DeliveryOptions[0].Distance

		lowDriveTime := order.Options.DeliveryOptions[0].DriveTime
		lowDistance := order.Options.DeliveryOptions[0].Distance

		var selectedDelivery *Delivery

		var sortedByDistance bool = true
		var sortedByDriveTime bool = true

		for j := 0; j < len(order.Options.DeliveryOptions); j++ {
			delivery := order.Options.DeliveryOptions[j]

			if j == 0 {
				selectedDelivery = order.Options.DeliveryOptions[j]
			}
			//if delivery.ETA.Low < lowestDriveTime && delivery.ETA.Method != "lower_bound"{
			//	lowestDriveTimeIndex = j
			//	lowestDriveTime = delivery.ETA.Low
			//}

			if delivery.DriveTime < lowDriveTime && sortedByDriveTime == true {
				sortedByDriveTime = false
			} else {
				lowDriveTime = delivery.DriveTime
			}

			if delivery.Distance < lowDistance && sortedByDistance == true {
				sortedByDistance = false
			} else {
				lowDistance = delivery.Distance
			}

			if delivery.DriveTime < lowestDriveTime {
				lowestDriveTimeIndex = j
				lowestDriveTime = delivery.DriveTime
			}
			if delivery.Distance < shortestDistance {
				shortestDistanceIndex = j
				shortestDistance = delivery.Distance
			}
		}

		lowestDriveTimeDelivery := order.Options.DeliveryOptions[lowestDriveTimeIndex]
		shortestDistanceDelivery := order.Options.DeliveryOptions[shortestDistanceIndex]

		shouldBeSortedBy := "distance"
		for _, v := range users {
			if v == order.UserId {
				shouldBeSortedBy = "drive_time"
			}
		}

		report := Report{
			order.OrderId,
			order.UserId,
			selectedDelivery.Vehicle,
			lowestDriveTimeDelivery.Vehicle,
			shortestDistanceDelivery.Vehicle,
			selectedDelivery,
			lowestDriveTimeDelivery,
			shortestDistanceDelivery,
			shouldBeSortedBy,
			sortedByDriveTime,
			sortedByDistance,
		}

		reports = append(reports, report)
		//}
	}

	return reports
}
