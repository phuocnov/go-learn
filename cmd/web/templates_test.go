package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	// Create a sllice of anonymost struct to hold the test cases
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2024, time.June, 15, 14, 30, 0, 0, time.UTC),
			want: "15 Jun 2024 at 14:30",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2024, time.December, 25, 9, 0, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "25 Dec 2024 at 09:00",
		},
	}

	// Loop through each test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := humanDate(tt.tm)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
