package add

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		desc     string
		a, b     int
		expected int
	}{
		{"TestBothZero", 0, 0, 0},
		{"TestBothPositive", 2, 5, 7},
		{"TestBothNegative", -2, -5, -7},
		{"TestAZero", 0, 3, 3},
		{"TestBZero", 3, 0, 3},
		{"TestANegativeBPositive", -1, 5, 4},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			actual := add(tc.a, tc.b)
			if actual != tc.expected {
				t.Fatalf("expected: %d got: %d for a: %d and b %d",
					tc.expected, actual, tc.a, tc.b)
			}
		})
	}

}
