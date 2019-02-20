package report

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Export(reports []Report) {

	file, err := os.Create("data/result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	head := []string{"order_id",
		"hub_id",
		"suggested hub_id",
		"hub ETA method",
		"suggested hub ETA method",
		"hub ETA low",
		"suggested hub ETA low",
		"hub ETA high",
		"suggested hub ETA high",
		"hub ETA predicted time inc. drive time",
		"suggested hub ETA predicted time inc. drive time",
		"hub ETA drive time",
		"suggested hub ETA drive time",
	}
	var reportData [][]string
	reportData = append(reportData, head)

	for _, value := range reports {
		row := []string { value.OrderID,
			strconv.Itoa(value.SelectedHub),
			strconv.Itoa(value.HubWithLowestEta),
			value.SelectedDelivery.ETA.Method,
			value.LowestEtaDelivery.ETA.Method,
			strconv.Itoa(value.SelectedDelivery.ETA.Low),
			strconv.Itoa(value.LowestEtaDelivery.ETA.Low),
			strconv.Itoa(value.SelectedDelivery.ETA.High),
			strconv.Itoa(value.LowestEtaDelivery.ETA.High),
			fmt.Sprintf("%f", value.SelectedDelivery.ETA.PredictedTime),
			fmt.Sprintf("%f", value.LowestEtaDelivery.ETA.PredictedTime),
			fmt.Sprintf("%f", value.SelectedDelivery.ETA.DriveTime),
			fmt.Sprintf("%f", value.LowestEtaDelivery.ETA.DriveTime),
		}
		reportData = append(reportData, row)
	}

	err = writer.WriteAll(reportData)
	checkError("Cannot write to file", err)
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}