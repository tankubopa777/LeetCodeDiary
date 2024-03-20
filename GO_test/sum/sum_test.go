package sum

import "testing"

func TestSum(t *testing.T){
	t.Run("should return 2+1", func(t *testing.T){
		// Arrange
		want := 3

		// Act
		got := sum(1,2)

		// Assert
		if got != want {
			t.Errorf("sum(1,2) = %d; want 3", got)
		}
	})

	// t.Run("should return 5+1", func(t *testing.T){
	// 	if sum(1,2) != 3 {
	// 		t.Log("Dammit it's error")
	// 	}
	// })

	// t.Run("should return 7+1", func(t *testing.T){
	// 	if sum(1,2) != 3 {
	// 		t.Log("Dammit it's error")
	// 	}
	// })
	
}
