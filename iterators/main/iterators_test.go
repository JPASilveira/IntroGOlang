package main

import "testing"

func TestIterators(t *testing.T) {
	verify := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("expected %q, got %q", expected, result)
		}
	}

	t.Run("iteratores test", func(t *testing.T) {
		result := iter(3, "*")
		expected := "***"
		verify(t, result, expected)
	})
}
func BenchmarkInterators(b *testing.B) {
	for n := 0; n < b.N; n++ {
		iter(3, "*")
	}
}
