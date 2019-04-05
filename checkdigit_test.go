package checkdigit

import "testing"

func TestNewLuhnProvider(t *testing.T) {

	p := NewLuhnProvider()
	if p == nil {
		t.Error("failed to generate luhn provider")
	}
}

func TestNewDammProvider(t *testing.T) {

	p := NewDammProvider()
	if p == nil {
		t.Error("failed to generate damm provider")
	}
}

func TestNewVerhoeffProvider(t *testing.T) {

	p := NewVerhoeffProvider()
	if p == nil {
		t.Error("failed to generate verhoeff provider")
	}
}
