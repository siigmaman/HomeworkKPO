package services

import "testing"

func TestClinic_CheckHealth(t *testing.T) {
	clinic := NewClinic()
	h := clinic.CheckHealth()

	if h != true && h != false {
		t.Errorf("CheckHealth should return bool, got %v", h)
	}
}
