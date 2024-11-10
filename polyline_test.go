package polyline

import (
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		points   [][]float64
		expected string
	}{
		{
			name: "Basic case",
			points: [][]float64{
				{38.5, -120.2},
				{40.7, -120.95},
				{43.252, -126.453},
			},
			expected: "_p~iF~ps|U_ulLnnqC_mqNvxq`@",
		},
		{
			name:     "Empty input",
			points:   [][]float64{},
			expected: "",
		},
	}

	// Run test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			encoded := Encode(tc.points)
			if encoded != tc.expected {
				t.Errorf("Encode(%v) = %v; want %v", tc.points, encoded, tc.expected)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		polyline string
		expected [][]float64
	}{
		{
			name:     "Basic case",
			polyline: "_p~iF~ps|U_ulLnnqC_mqNvxq`@",
			expected: [][]float64{
				{38.5, -120.2},
				{40.7, -120.95},
				{43.252, -126.453},
			},
		},
		{
			name:     "Empty input",
			polyline: "",
			expected: [][]float64{},
		},
	}

	// Run test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			decoded := Decode(tc.polyline)

			// Explicitly check for empty slices
			if len(tc.expected) == 0 && len(decoded) == 0 {
				return // Both are empty, so the test passes
			}

			// Use reflect.DeepEqual for non-empty cases
			if !reflect.DeepEqual(decoded, tc.expected) {
				t.Errorf("Decode(%v) = %v; want %v", tc.polyline, decoded, tc.expected)
			}
		})
	}
}
