package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Order struct {
	DelayedFireAt   string      `json:"delayed_fire_at"`
	DeliveryOptions []*Delivery `json:"deliveryOptions"`
}

type Delivery struct {
	Type     uint8   `json:"delivery_type"`
	Vehicle  uint16  `json:"vehicle_id"`
	Items    []int   `json:"items"`
	Distance float64 `json:"distance"`
	ETA      *ETA    `json:"ETA"`
}

type ETA struct {
	Low                    uint8   `json:"low"`
	High                   uint8   `json:"high"`
	Method                 string  `json:"method"`
	OrderPlace             uint8   `json:"order_place_in_line"`
	PiesWaitingToCookDelay int8    `json:"pies_waiting_to_cook_delay"`
	PredictedTime          float32 `json:"predicted_time"`
	DriveTime              float32 `json:"drive_time"`
}



func main() {

	if len(os.Args[1:]) == 0 {
		fmt.Println("Provide path to file as a fist argument.")
		return
	}

	fileName := os.Args[1:][0]

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))

	lineCount := 0
	var orders []Order
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if line[0] == "data" {
			continue
		}

		fmt.Println("Record", lineCount, "has", len(line), "fields")
		lineCount += 1

		option := Order{}
		err = json.Unmarshal([]byte(line[0]), &option)
		if err != nil {
			log.Fatal(err)
		}
		orders = append(orders, option)
	}

	fmt.Println("Parsed: ", len(orders), "orders.")
	var ordersWithNoOptions []Order
	var ordersWithOneOption []Order
	var ordersWithOkOptions []Order
	var ordersWithWrongOptions []Order

	for i := 0; i < len(orders); i++ {
		order := orders[i]
		deliveryOptionsLength := len(order.DeliveryOptions)
		if deliveryOptionsLength == 0 {
			fmt.Println("No delivery option in ", i, "line!")
			ordersWithNoOptions = append(ordersWithNoOptions, order)
			continue
		}

		if deliveryOptionsLength == 1 {
			fmt.Println("Only 1 delivery option in ", i, "line!")
			ordersWithOneOption = append(ordersWithOneOption, order)
			continue
		}

		lowestEtaIndex := 0
		lowestEta := order.DeliveryOptions[lowestEtaIndex].ETA.Low


		for j := 0; j < len(order.DeliveryOptions); j++ {
			delivery := order.DeliveryOptions[j]
			if delivery.ETA.Low < lowestEta {
				lowestEtaIndex = j
				lowestEta = delivery.ETA.Low
			}
		}

		if lowestEtaIndex > 0 {
			ordersWithWrongOptions = append(ordersWithWrongOptions, order)
		} else {
			ordersWithOkOptions = append(ordersWithOkOptions, order)
		}
	}

	fmt.Println("ordersWithOkOptions: ", len(ordersWithOkOptions))
	fmt.Println("ordersWithWrongOptions: ", len(ordersWithWrongOptions))
	fmt.Println("ordersWithOneOption: ", len(ordersWithOneOption))
	fmt.Println("ordersWithNoOptions: ", len(ordersWithNoOptions))

	wrong, _ := json.Marshal(ordersWithWrongOptions)
	ok, _ := json.Marshal(ordersWithOkOptions)

	ioutil.WriteFile("wrongOrders.json", wrong, 0644)
	ioutil.WriteFile("okOrders.json", ok, 0644)
}
