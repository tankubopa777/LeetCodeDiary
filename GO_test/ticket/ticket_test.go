package ticket

import "testing"


func TestTicketPrice(t *testing.T){
	
	test := []struct{
		name string
		age int 
		want float64
	}{
		{"Free ticket when age less than 3", 0, 0.0},
		{"Free ticket when age less than 3", 0, 15.0},
		{"Free ticket when age less than 3", 0, 0.0},
		{"Free ticket when age less than 3", 0, 0.0},
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