package services

/*
int CheckHealthC();
*/
import "C"

type Clinic struct{}

func NewClinic() *Clinic {
	return &Clinic{}
}

func (c *Clinic) CheckHealth() bool {
	return C.CheckHealthC() == 1
}
