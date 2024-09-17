package main

import "testing"

func TestGeometricShapes(t *testing.T) {
	t.Run("area", func(t *testing.T) {
		t.Helper()
		areaTests := []struct {
			form Form
			want float64
		}{
			{form: Rectangle{Width: 12, Height: 6}, want: 72},
			{form: Circle{radius: 10}, want: 314.1592653589793},
			{form: Triangle{base: 12, height: 6}, want: 36.0},
		}
		for _, tt := range areaTests {
			result := tt.form.Area()
			if result != tt.want {
				t.Errorf("%#v area = %.2f; want %.2f", tt.form, result, tt.want)
			}
		}
	})
}
