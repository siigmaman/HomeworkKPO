package asmutils

import "testing"

func TestSumFood(t *testing.T) {
	if SumFood(3, 4) != 7 {
		t.Errorf("SumFood failed")
	}

	if SumFood(0, 0) != 0 {
		t.Errorf("SumFood zero failed")
	}
}

func TestCalcDailyFeedAverage(t *testing.T) {
	if CalcDailyFeedAverage(10, 2) != 5 {
		t.Errorf("Average failed")
	}

	if CalcDailyFeedAverage(10, 0) != 0 {
		t.Errorf("Division by zero should return 0")
	}
}

func TestSleepASM(t *testing.T) {
	SleepASM(1)
}
