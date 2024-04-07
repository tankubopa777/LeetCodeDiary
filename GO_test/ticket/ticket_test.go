package ticket

import "testing"


func TestTicketPrice(t *testing.T){
	
	test := []struct{
		name string
		age int 
		want float64
	}{
		{"Free ticket when age less than 3", 0, 0.0},
		{"Free ticket when age under 3", 3, 0.0},
		{"Ticket 15 when 4 year old", 4, 15.0},
		{"Ticket 15 when 15 year old", 15, 15.0},
		{"Ticket 30 when over 15 year old", 16, 30.0},
		{"Ticket 30 when age is 50", 50, 30.0},
		{"Ticket 5 when age is 50", 51, 5.0},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T){
			got := Price(tt.age) 
			if got != tt.want{
				t.Errorf("Dont need that")
			}
		})
	}


	}