package executor

import (
	"bufio"
	"fmt"
	"omise/challenges/api"
	"omise/challenges/cipher"
	"omise/challenges/commandline"
	"omise/challenges/parser"
	"omise/challenges/report"
	"os"
	"sync"
	"time"

	"github.com/urfave/cli"
)

//Execute ...
func Execute(c *cli.Context) error {

	filePath, requestNum, timeLimit := commandline.ValidateArgs(c)

	file, _ := os.Open(filePath)

	reader, _ := cipher.NewRot128Reader(file)
	limiter := time.Tick(time.Duration(timeLimit) * time.Millisecond)
	reportData := api.ReportData{}
	lines := make(chan string)
	wg := new(sync.WaitGroup)
	fmt.Println("performing donations...")
	go ReadLinesFromCSV(lines, reader)

	for w := 1; w <= requestNum; w++ {
		wg.Add(1)

		go ProcessLine(lines, wg, limiter, &reportData)
	}

	wg.Wait()

	fmt.Println("done.")
	fmt.Println(report.Builder(reportData))
	return nil
}

//ReadLinesFromCSV scans and reads in each line into the lines channel
func ReadLinesFromCSV(lines chan string, file *cipher.Rot128Reader) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		lines <- scanner.Text()
	}
	close(lines)
}

//ProcessLine parses each line and calculates the charge
func ProcessLine(lines <-chan string, wg *sync.WaitGroup, limiter <-chan time.Time, reportData *api.ReportData) {
	defer wg.Done()

	for line := range lines {
		<-limiter

		customer := parser.ParseCustomerFromCSV(line)
		newClient := api.CreateNewClient()
		if customer != (api.Customer{}) {
			newToken := api.CreateNewToken(newClient, customer)

			api.ChargeCustomer(newToken, newClient, customer, reportData)

		}
	}
}
