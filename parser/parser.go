package parser

import (
	"fmt"
	"omise/challenges/api"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//ValidateUserInputLine makes sure each line from the decrypted file has the correct format
func ValidateUserInputLine(line string) bool {
	r, err := regexp.Compile(`^(.+),(\d+),(\d+),(\d+),(\d+),(\d+)$`)

	if err != nil {
		fmt.Printf("There is a problem with your regexp.\n")
		return false
	}

	if r.MatchString(line) == true {
		return true
	}
	return false

}

//ParseCustomerFromCSV reads each line of the decrypted file and parses the data by ',' and returns a customer struct using the data
func ParseCustomerFromCSV(line string) api.Customer {
	if ValidateUserInputLine(line) == true {
		data := strings.Split(line, ",")

		return api.Customer{
			Name:           data[0],
			Number:         data[2],
			ExpMonth:       CreditCardMonth(data[4]),
			ExpYear:        CreditCardYear(data[5]),
			DonationAmount: DonationAmount(data[1]),
		}
	}
	return api.Customer{}

}

//CreditCardMonth converts the string type month to time.Month
func CreditCardMonth(month string) time.Month {
	integerMonth, _ := strconv.Atoi(month)
	return time.Month(integerMonth)
}

//CreditCardYear converts the string type year to int
func CreditCardYear(year string) int {
	integerYear, _ := strconv.Atoi(year)
	return integerYear
}

//DonationAmount converts the string type donation amount to int64
func DonationAmount(donation string) int64 {
	amount, _ := strconv.Atoi(donation)

	return int64(amount)
}
