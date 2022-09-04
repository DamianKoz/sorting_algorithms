package main

import "testing"

func TestArrayInitilation(t *testing.T) {
	got := initializeArray(100)

	t.Run("array length", func(t *testing.T) {
		want := 5

		if len(got) != want {
			t.Errorf("Got length of %v, Want length of %v", len(got), want)
		}
	})
}
