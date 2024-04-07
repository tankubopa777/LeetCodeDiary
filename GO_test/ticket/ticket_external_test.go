package ticket_test

import (
	"testing"
	"ticket"
)

func TestTicket(t *testing.T) {
	t.Run("should return 0 when age is 3", func(t *testing.T) {
		age := 3
		want := 0.0
		got := ticket.Price(age)
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})
}