package entities

import (
	"testing"
)

func TestCalcTotalCapitalization(t *testing.T) {
	tests := []struct {
		name        string
		quote       Quote
		expected    string
		expectPanic bool
	}{
		{
			name: "No Pref Capitalization",
			quote: Quote{
				Ticker:         "AAPL",
				Price:          150.00,
				Time:           "2023-10-01T10:00:00Z",
				SeqNum:         "1",
				Capitalization: "1000000000",
			},
			expected:    "1000000000",
			expectPanic: false,
		},
		{
			name: "With Pref Capitalization",
			quote: Quote{
				Ticker:         "AAPL",
				Price:          150.00,
				Time:           "2023-10-01T10:00:00Z",
				SeqNum:         "1",
				Capitalization: "1000000000",
				Pref: &Quote{
					Ticker:         "AAPL.P",
					Price:          100.00,
					Time:           "2023-10-01T10:00:00Z",
					SeqNum:         "2",
					Capitalization: "500000000",
				},
			},
			expected:    "1500000000",
			expectPanic: false,
		},
		{
			name: "Invalid Capitalization",
			quote: Quote{
				Ticker:         "AAPL",
				Price:          150.00,
				Time:           "2023-10-01T10:00:00Z",
				SeqNum:         "1",
				Capitalization: "not-a-number",
			},
			expected:    "",
			expectPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("expected panic but did not get one")
					}
				}()
			}

			result := tt.quote.CalcTotalCapitalization().String()
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}
