package services

/*
float CheckTempC();
*/
import "C"

func CheckTemp() float32 {
	return float32(C.CheckTempC())
}
