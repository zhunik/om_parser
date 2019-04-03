package report

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Order struct {
	OrderId   string   `json:"order_id"`
	CreatedAt string   `json:"event_at"`
	UserId    int      `json:"user_id"`
	Options   *Options `json:"options"`
}

type Options struct {
	DelayedFireAt   string      `json:"delayed_fire_at"`
	DeliveryOptions []*Delivery `json:"deliveryOptions"`
}

type Delivery struct {
	Type      uint8   `json:"delivery_type"`
	Vehicle   int     `json:"vehicle_id"`
	Items     []int   `json:"items"`
	Distance  float64 `json:"distance"`
	DriveTime int     `json:"drive_time"`
	ETA       *ETA    `json:"ETA"`
}

type ETA struct {
	Low                    int     `json:"low"`
	High                   int     `json:"high"`
	Method                 string  `json:"method"`
	OrderPlace             int     `json:"order_place_in_line"`
	PiesWaitingToCookDelay int     `json:"pies_waiting_to_cook_delay"`
	PredictedTime          float32 `json:"predicted_time"`
	DriveTime              float32 `json:"drive_time"`
}

func Parse(fileName string) ([]Order, error) {

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))

	var lineCount uint = 1
	var orders []Order
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if line[0] == "order_id" || line[0] == "" {
			continue
		}

		fmt.Println("Record", lineCount, "has", len(line), "fields")
		lineCount += 1

		options := Options{}
		err = json.Unmarshal([]byte(line[2]), &options)
		if err != nil {
			log.Fatal(err)
		}

		userId, _ := strconv.Atoi(line[3])

		order := Order{
			line[0], line[1], userId, &options,
		}

		orders = append(orders, order)
	}

	fmt.Println("Parsed: ", len(orders), "orders.")

	return orders, nil
}
