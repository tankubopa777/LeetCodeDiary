package teardown

import "testing"

func setup(t *testing.T) {
	t.Log("Before all tests")

	return func() {
		t.Log("After all tests")
	}
}

func TestTeardown(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	t.Run("subset", func(t *testing.T) {
		t.Log("This is a subset")
	}
	)
}