package main

import "testing"

func TestSort(t *testing.T) {
	tests := []struct {
		name          string
		w, h, l, mass int
		want          string
	}{
		// STANDARD cases
		{"AllJustBelowThresholds", 90, 80, 70, 19, standard},
		{"SmallDimensionsSmallMass", 10, 10, 10, 5, standard},

		// Bulky
		{"BulkyByVolumeBoundary", 100, 100, 100, 10, special}, // 1,000,000

		// REJECTED
		{"RejectedByBothDimensionAndHeavy", 150, 10, 10, 20, rejected},
		{"RejectedByVolumeAndHeavy", 100, 100, 100, 25, rejected},

		// Large values
		{"LargeAll", 1000, 1000, 1000, 100, rejected},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := sort(tc.w, tc.h, tc.l, tc.mass)
			if got != tc.want {
				t.Fatalf("sort(%d,%d,%d,%d) = %q; want %q",
					tc.w, tc.h, tc.l, tc.mass, got, tc.want)
			}
		})
	}
}
