package main

import (
	"encoding/json"
	"fmt"
	"github.com/zhunik/om_parser/report"
	"log"
	"os"
)


func main() {

	if len(os.Args[1:]) == 0 {
		fmt.Println("Provide path to file as a fist argument.")
	}
	fileName := os.Args[1:][0]

	fmt.Println(fileName)

	orders, err := report.Parse(fileName)

	if err != nil {
		log.Fatal(err)
	}

	reports := report.Process(orders)


	report.Export(reports)

	printReports, _ := json.Marshal(reports)
	fmt.Println(string(printReports))
}
