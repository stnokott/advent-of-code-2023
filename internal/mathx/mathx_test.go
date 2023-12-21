// Package mathx provides utility functions for mathematics
package mathx

import (
	"math"
	"testing"
)

func TestMinInt(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"simple", args{64, 33}, 33},
		{"equal", args{78, 78}, 78},
		{"negative", args{-87, -556}, -556},
		{"min_int", args{math.MinInt + 1, math.MinInt}, math.MinInt},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MinInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"simple", args{64, 33}, 64},
		{"equal", args{78, 78}, 78},
		{"negative", args{-87, -556}, -87},
		{"max_int", args{math.MaxInt32 - 1, math.MaxInt32}, math.MaxInt32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MaxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
