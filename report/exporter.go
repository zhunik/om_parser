package report

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func Export(reports []Report) {

	date := time.Now()
	file, err := os.Create(fmt.Sprintf("result_%d-%d-%d.csv", date.Year(), date.Month(), date.Day()))
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	head := []string{
		"order_id",
		"user_id",
		"hub_id",
		"hub id with shortest distance",
		"hub_id with lowest drive_time",
		"hub distance",
		"shortest distance",
		"distance with lowest drive time",
		"hub drive_time",
		"drive_time for shortest distance",
		"lowest drive_time",
		"should be sorted by",
		"sorting by distance",
		"sorting by drive_time",
		"status",
	}
	var reportData [][]string
	reportData = append(reportData, head)

	for _, value := range reports {

		status := "OK"

		if value.shouldBeSortedBy == "drive_time" && value.sortedByDriveTime == false {
			status = "ERROR"
		}
		row := []string{
			value.OrderID,
			strconv.Itoa(value.UserID),
			strconv.Itoa(value.SelectedHub),
			strconv.Itoa(value.HubWithShortestDistance),
			strconv.Itoa(value.HubWithLowestEta),
			//value.SelectedDelivery.ETA.Method,
			//value.LowestEtaDelivery.ETA.Method,
			//strconv.Itoa(value.SelectedDelivery.ETA.Low),
			//strconv.Itoa(value.LowestEtaDelivery.ETA.Low),
			//strconv.Itoa(value.SelectedDelivery.ETA.High),
			//strconv.Itoa(value.LowestEtaDelivery.ETA.High),
			//fmt.Sprintf("%f", value.SelectedDelivery.ETA.PredictedTime),
			//fmt.Sprintf("%f", value.LowestEtaDelivery.ETA.PredictedTime),
			//fmt.Sprintf("%f", value.SelectedDelivery.ETA.DriveTime),
			//fmt.Sprintf("%f", value.LowestEtaDelivery.ETA.DriveTime),
			fmt.Sprintf("%f", value.SelectedDelivery.Distance),
			fmt.Sprintf("%f", value.ShortestDistanceDelivery.Distance),
			fmt.Sprintf("%f", value.LowestEtaDelivery.Distance),
			strconv.Itoa(value.SelectedDelivery.DriveTime),
			strconv.Itoa(value.ShortestDistanceDelivery.DriveTime),
			strconv.Itoa(value.LowestEtaDelivery.DriveTime),
			value.shouldBeSortedBy,
			strconv.FormatBool(value.sortedByDistance),
			strconv.FormatBool(value.sortedByDriveTime),
			status,
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
