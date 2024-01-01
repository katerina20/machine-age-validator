package utils

import (
	"testing"
)

func TestParseToDays(t *testing.T) {
	testCases := []struct {
		name     string
		age      string
		expected int
		wantErr  bool
	}{
		{
			name:     "ValidYears",
			age:      "2 years",
			expected: 730,
			wantErr:  false,
		},
		{
			name:     "ValidMonths",
			age:      "3 months",
			expected: 90,
			wantErr:  false,
		},
		{
			name:     "ValidMonths",
			age:      "3 days",
			expected: 3,
			wantErr:  false,
		},
		{
			name:     "ValidMonths",
			age:      "5 weeks",
			expected: 35,
			wantErr:  false,
		},
		{
			name:     "InvalidFormat",
			age:      "one day",
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "InvalidFormat",
			age:      "15",
			expected: 0,
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseToDays(tc.age)
			if (err != nil) != tc.wantErr {
				t.Errorf("ParseToDays(%s) error = %v, wantErr %v", tc.age, err, tc.wantErr)
				return
			}
			if result != tc.expected {
				t.Errorf("ParseToDays(%s) = %v, want %v", tc.age, result, tc.expected)
			}
		})
	}
}
