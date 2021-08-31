package strategy

import "testing"

func TestPayStrategy(t *testing.T) {
	payment1 := NewPayment("Ada", "", 123, &Cash{})
	payment1.Pay()

	payment2 := NewPayment("Bob", "0002", 888, &Bank{})
	payment2.Pay()
}
