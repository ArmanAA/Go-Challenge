package report

import (
	"omise/challenges/api"

	"github.com/leekchan/accounting"
)

//Builder creates the final report string
func Builder(report api.ReportData) string {
	ac := accounting.Accounting{Symbol: "THB ", Precision: 2}
	return "total received: " + ac.FormatMoney(CalcTotalReceived(report.SuccessfulSum, report.FailedSum)) +
		"\nsuccessfully donated: " + ac.FormatMoney(report.SuccessfulSum) +
		"\nfaulty donation: " + ac.FormatMoney(report.FailedSum) + "\n \n" +
		"\naverage per person: " + ac.FormatMoney((CalcAverage(report.SuccessfulSum, report.DonationCount))) + "\n" +
		"\t \t " + report.First.Name + "\n" +
		"\t \t " + report.Second.Name + "\n" +
		"\t \t " + report.Third.Name + "\n"

}

//CalcAverage calculates the total average donations
func CalcAverage(totalAmount int64, totalCount int) float64 {
	if totalAmount == 0 {
		return 0.0
	}
	return float64(totalAmount / int64(totalCount))

}

// CalcTotalReceived sums the failed and successful charges
func CalcTotalReceived(successfulDonation int64, failedDonation int64) int64 {
	return failedDonation + successfulDonation
}
