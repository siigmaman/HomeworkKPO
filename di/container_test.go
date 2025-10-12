package di

import "testing"

func TestNewContainer(t *testing.T) {
	c := NewContainer()
	if c.Clinic == nil || c.Zoo == nil || c.Menu == nil {
		t.Errorf("Container fields should not be nil")
	}
}
