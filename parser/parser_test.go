package parser

import (
	"omise/challenges/api"
	"strconv"
	"testing"
	"time"
)

func TestValidateUserInput(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"Ms. Daisy C Tûk,2213937,4556712499523363,814,5,2022", true},
		{"123AAaa@#$%^mv,2213937,4556712499523363,814,5,2022", true},
		{"0,0,4,14,5,2", true},
		{"Ms. Daisy C Tûk,22139$37,4556712499523363,814,5,2022", false},
		{"Ms. Daisy C Tûk,2213937,455^^^6712499523363,814,5,2022", false},
		{"Ms. Daisy C Tûk,2213937,4556712499523363,81##4,5,2022", false},
		{"Ms. Daisy C Tûk,2213937,4556712499523363,814,5*@!,2022", false},
		{"Ms. Daisy C Tûk,2213937,4556712499523363,814,5,202a2", false},
		{"Ms. Daisy C Tûk,221", false},
		{" ", false},
	}
	for _, test := range tests {
		if result := ValidateUserInputLine(test.input); result != test.expected {
			t.Error("Test failed: {} inputed, {} expected, received: {}", test.input, test.expected, result)
		}
	}
}
func TestParseCustomerFromCSV(t *testing.T) {
	var tests = []struct {
		input    string
		expected api.Customer
	}{
		{"hello", api.Customer{}},
		{"Ms. Daisy C Tûk,2213937,4556712499523363,814,5,2022", api.Customer{
			Name:           "Ms. Daisy C Tûk",
			Number:         "4556712499523363",
			ExpMonth:       CreditCardMonth("5"),
			ExpYear:        CreditCardYear("2022"),
			DonationAmount: DonationAmount("2213937")}},
	}
	for _, test := range tests {
		if result := ParseCustomerFromCSV(test.input); result != test.expected {
			t.Error("Test failed: {} inputed, {} expected, received: {}", test.input, test.expected, result)
		}
	}

}
func TestCreditCardMonth(t *testing.T) {

	integerMonth, _ := strconv.Atoi("9")

	if CreditCardMonth("9") != time.Month(integerMonth) {
		t.Error("Expected September")
	}
}
func TestCreditCardYear(t *testing.T) {

	integerYear, _ := strconv.Atoi("2022")

	if CreditCardYear("2022") != integerYear {
		t.Error("Expected ", 2022)
	}
}
func TestDonationAmount(t *testing.T) {

	integer, _ := strconv.Atoi("2022")

	if DonationAmount("2022") != int64(integer) {
		t.Error("Expected ", 2022)
	}
}
