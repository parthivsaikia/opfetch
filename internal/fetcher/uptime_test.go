package fetcher_test

import (
	"os"
	"testing"

	"github.com/parthivsaikia/opfetch/internal/fetcher"
)

func TestGetUptime(t *testing.T) {

	tests := []struct {
		name         string
		uptimeData   string
		expectedDays int
		expecteHours int
		expectedMin  int
		expectedSec  int
		expectedStr  string
	}{
		{
			name:         "uptime is less than a minute",
			uptimeData:   "45.73 30.12",
			expectedDays: 0,
			expecteHours: 0,
			expectedMin:  0,
			expectedSec:  45,
			expectedStr:  "45s",
		},
		{
			name:         "uptime is less than an hour",
			uptimeData:   "2345.89 1987.44",
			expectedDays: 0,
			expecteHours: 0,
			expectedMin:  39,
			expectedSec:  5,
			expectedStr:  "39m",
		},
		{
			name:         "uptime is more than an hour",
			uptimeData:   "19876.12 17654.33",
			expectedDays: 0,
			expecteHours: 5,
			expectedMin:  31,
			expectedSec:  16,
			expectedStr:  "5h 31m",
		},
		{
			name:         "uptime is more than a day",
			uptimeData:   "176543.78 165432.10",
			expectedDays: 2,
			expecteHours: 1,
			expectedMin:  2,
			expectedSec:  23,
			expectedStr:  "2d 1h 2m",
		},
	}

	for _, testCases := range tests {
		t.Run(testCases.name, func(t *testing.T) {

			fileName := "test-uptime"
			file, err := os.CreateTemp("", fileName)
			if err != nil {
				t.Fatal(err)
			}

			defer os.Remove(file.Name())

			file.Write([]byte(testCases.uptimeData))
			file.Close()
			uptime, err := fetcher.GetUptime(file.Name())
			if err != nil {
				t.Fatal(err)
			}

			if testCases.expectedDays != uptime.Days {
				t.Errorf("Expected days to be %d but %d", testCases.expectedDays, uptime.Days)
			}

			if testCases.expecteHours != uptime.Hours {
				t.Errorf("Expected hour to be %d but got %d", testCases.expecteHours, uptime.Hours)
			}

			if testCases.expectedMin != uptime.Minutes {
				t.Errorf("Expected minutes to be %d but got %d", testCases.expectedMin, uptime.Minutes)
			}

			if testCases.expectedSec != uptime.Seconds {
				t.Errorf("Expected seconds to be %d but got %d", testCases.expectedSec, uptime.Seconds)
			}

			if testCases.expectedStr != uptime.String() {
				t.Errorf("Expected string to be %s but got %s", testCases.expectedStr, uptime.String())
			}

		})
	}

}
