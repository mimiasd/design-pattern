package simplefactory

import "testing"

func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("mimiasd")
	if s != "Hi, mimiasd" {
		t.Fatal("Type1 test fail")
	}
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("mimiasd")
	if s != "Hello, mimiasd" {
		t.Fatal("Type2 test fail")
	}
}
