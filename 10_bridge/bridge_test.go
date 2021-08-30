package bridge

import "testing"

func ExampleCommonSMS() {
	m := NewCommMessage(ViaSMS())
	m.SendMessage("have a drink?", "bob")
}

func ExampleCommonEmail() {
	m := NewCommMessage(ViaEmail())
	m.SendMessage("have a drink?", "bob")
}

func ExampleUrgencySMS() {
	m := NewUrgencyMessage(ViaSMS())
	m.SendMessage("have a drink?", "bob")
}

func ExampleUrgencyEmail() {
	m := NewUrgencyMessage(ViaEmail())
	m.SendMessage("have a drink?", "bob")
}

func TestALL(t *testing.T) {
	ExampleCommonSMS()
	ExampleCommonEmail()
	ExampleUrgencySMS()
	ExampleUrgencyEmail()
}
