package service

import (
	"errors"
	"fin_data_processing/internal/entities"
	"testing"
)

func TestTargetsAchievementCheck(t *testing.T) {
	tests := []struct {
		name             string
		target           entities.TargetUser
		fundamental      entities.Fundamental
		quote            entities.Quote
		expectedAchieved bool
		expectedResult   float64
		expectedError    error
	}{
		{
			name: "Target PBV achieved",
			target: entities.TargetUser{
				Target: entities.Target{
					ValuationRatio: TargetPbv,
					Value:          1.5,
				},
			},
			fundamental: entities.Fundamental{
				BookValue: 100,
			},
			quote: entities.Quote{
				Price:          150,
				Capitalization: "150",
			},
			expectedAchieved: false,
			expectedResult:   1.5,
			expectedError:    nil,
		},
		{
			name: "Target PE achieved",
			target: entities.TargetUser{
				Target: entities.Target{
					ValuationRatio: TargetPe,
					Value:          1.5,
				},
			},
			fundamental: entities.Fundamental{
				NetIncome: 100,
			},
			quote: entities.Quote{
				Price:          150,
				Capitalization: "150",
			},
			expectedAchieved: true,
			expectedResult:   1.5,
			expectedError:    nil,
		},
		{
			name: "Target PS achieved",
			target: entities.TargetUser{
				Target: entities.Target{
					ValuationRatio: TargetPs,
					Value:          10,
				},
			},
			fundamental: entities.Fundamental{
				Revenue: 50,
			},
			quote: entities.Quote{
				Price:          200,
				Capitalization: "200",
			},
			expectedAchieved: true,
			expectedResult:   4,
			expectedError:    nil,
		},
		{
			name: "Target Price achieved",
			target: entities.TargetUser{
				Target: entities.Target{
					ValuationRatio: TargetPrice,
					Value:          100,
				},
			},
			fundamental: entities.Fundamental{},
			quote: entities.Quote{
				Price:          100,
				Capitalization: "100",
			},
			expectedAchieved: true,
			expectedResult:   100,
			expectedError:    nil,
		},
		{
			name: "Zero Book Value",
			target: entities.TargetUser{
				Target: entities.Target{
					ValuationRatio: TargetPbv,
					Value:          1.5,
				},
			},
			fundamental: entities.Fundamental{
				BookValue: 0,
			},
			quote: entities.Quote{
				Price:          150,
				Capitalization: "150",
			},
			expectedAchieved: false,
			expectedResult:   0,
			expectedError:    errors.New("book value is zero"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isAchieved, currResult, err := TargetsAchievementCheck(tt.target, tt.fundamental, tt.quote)

			if isAchieved != tt.expectedAchieved {
				t.Errorf("expected isAchieved %v, got %v", tt.expectedAchieved, isAchieved)
			}
			if currResult != tt.expectedResult {
				t.Errorf("expected currResult %v, got %v", tt.expectedResult, currResult)
			}
			if (err != nil && tt.expectedError == nil) || (err == nil && tt.expectedError != nil) {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			} else if err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error message %v, got %v", tt.expectedError, err)
			}
		})
	}
}
