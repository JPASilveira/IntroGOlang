package main

import "testing"

func TestIterators(t *testing.T) {
	verify := func(t *testing.T, resulted, expected string) {
		t.Helper()
		if resulted != expected {
			t.Errorf("expected %q, got %q", expected, resulted)
		}
	}

	t.Run("iteratores test", func(t *testing.T) {
		resulted := iter(3, "*")
		expected := "***"
		verify(t, resulted, expected)
	})
}
func BenchmarkInterators(b *testing.B) {
	for n := 0; n < b.N; n++ {
		iter(3, "*")
	}
}
