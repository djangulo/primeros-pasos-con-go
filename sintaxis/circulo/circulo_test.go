package main

import (
	"fmt"
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	const epsilon = 0.05
	for _, tt := range []struct {
		in, want float64
	}{
		{0.0, 0},
		{0.1, 0.031416},
		{1.0, 3.1416},
		{2.0, 12.56},
		{3.0, 28.26},
	} {
		t.Run(fmt.Sprintf("area(%.4f)==%.4f", tt.in, tt.want), func(t *testing.T) {
			got := area(tt.in)

			if math.Abs(tt.want-got) > epsilon {
				t.Errorf(
					"expected abs(%.4f - %.4f) < %.4f, got diff %.4f",
					tt.want,
					got,
					epsilon,
					math.Abs(tt.want-got),
				)
			}
		})

	}
}
