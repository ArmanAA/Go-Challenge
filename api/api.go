package api

import (
	"log"
	"sync"
	"time"

	omise "github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

//Customer is a struct used to store each doner
type Customer struct {
	Name           string
	Number         string
	ExpMonth       time.Month
	ExpYear        int
	DonationAmount int64
}

//UserData stores each customers name and the amount they donated
type UserData struct {
	Name   string
	Amount int64
}

//ReportData keeps track of everything that will be reported at the end
type ReportData struct {
	SuccessfulSum int64
	FailedSum     int64
	DonationCount int
	First         UserData
	Second        UserData
	Third         UserData
}

const (
	//OmisePublicKey from environment variables or configuration files!
	OmisePublicKey = "pkey_test_521w1g1t7w4x4rd22z0"
	//OmiseSecretKey from environment variables or configuration files!
	OmiseSecretKey = "skey_test_521w1g1t6yh7sx4pu8n"
)

var data = ReportData{
	SuccessfulSum: -1,
	FailedSum:     -1,
	DonationCount: -1,
	First: UserData{
		Name:   "",
		Amount: -1,
	},
	Second: UserData{
		Name:   "",
		Amount: -1,
	},
	Third: UserData{
		Name:   "",
		Amount: -1,
	},
}
var mutex = &sync.Mutex{}

//ChargeCustomer charges the customer based on the donation amount that they have
func ChargeCustomer(token *omise.Token, newClient *omise.Client, customer Customer, data *ReportData) {
	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   int64(customer.DonationAmount),
		Currency: "thb",
		Card:     token.ID,
	}

	if e := newClient.Do(charge, createCharge); e != nil {

		log.Println(e)
	}
	mutex.Lock()
	if charge.Status == "successful" {
		CalculateTopThree(data, UserData{Name: customer.Name, Amount: customer.DonationAmount})

		data.SuccessfulSum += customer.DonationAmount

	} else if charge.Status == "failed" {

		data.FailedSum += customer.DonationAmount

	}

	data.DonationCount++
	mutex.Unlock()

}

//CreateNewClient creates a new client using the omise public and private keys
func CreateNewClient() *omise.Client {
	client, e := omise.NewClient(OmisePublicKey, OmiseSecretKey)
	if e != nil {
		log.Println(e)
	}
	return client
}

//CalculateTopThree calculates the top three donations
func CalculateTopThree(data *ReportData, user UserData) {

	if user.Amount > data.First.Amount {
		data.Third = data.Second
		data.Second = data.First
		data.First = user
	} else if user.Amount > data.Second.Amount {
		data.Third = data.Second
		data.Second = user
	} else if user.Amount > data.Third.Amount {
		data.Third = user

	}
}

//CreateNewToken creates a new token for each customer
func CreateNewToken(client *omise.Client, customer Customer) *omise.Token {
	token, createToken := &omise.Token{}, &operations.CreateToken{
		Name:            customer.Name,
		Number:          customer.Number,
		ExpirationMonth: customer.ExpMonth,
		ExpirationYear:  customer.ExpYear,
	}

	if e := client.Do(token, createToken); e != nil {
		log.Println(e)
	}

	return token
}
