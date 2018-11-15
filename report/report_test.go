package report

import "testing"

func CalcAverageTest(t *testing.T) {
	if CalcAverage(int64(100), 20) != float64(5) {
		t.Error("Expected 5")
	}
	if CalcAverage(int64(0), 0) != float64(0.0) {
		t.Error("Expected 0")
	}

}

func CalcTotalReceivedTest(t *testing.T) {
	if CalcTotalReceived(100, 100) != int64(200) {
		t.Error("Expected 200")
	}

}
