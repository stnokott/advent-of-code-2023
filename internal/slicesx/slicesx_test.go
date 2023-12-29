// Package slicesx provides high-level utility functions for slice/array operations

package slicesx

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
)

func TestIntersect(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"one_empty", args{[]int{26}, []int{}}, []int{}},
		{"both_empty", args{[]int{}, []int{}}, []int{}},
		{"one_element_equal", args{[]int{17}, []int{17}}, []int{17}},
		{"one_element_unequal", args{[]int{17}, []int{99}}, []int{}},
		{"equal_size", args{[]int{14, 7, 88, 95}, []int{96, 95, 14, 56}}, []int{14, 95}},
		{"unequal_size", args{[]int{83, 86, 6, 31, 17, 9, 48, 53}, []int{41, 48, 83, 86, 17}}, []int{17, 48, 83, 86}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
			if got := IntersectBrute(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectBrute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectSized(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"one_element_equal", args{[]int{17}, []int{17}}, []int{17}},
		{"one_element_unequal", args{[]int{17}, []int{99}}, []int{}},
		{"equal_size", args{[]int{96, 95, 14, 56}, []int{14, 7, 88, 95}}, []int{14, 95}},
		{"unequal_size", args{[]int{41, 48, 83, 86, 17}, []int{83, 86, 6, 31, 17, 9, 48, 53}}, []int{17, 48, 83, 86}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersectSized(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersectSized() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectSizedBrute(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"one_element_equal", args{[]int{17}, []int{17}}, []int{17}},
		{"one_element_unequal", args{[]int{17}, []int{99}}, []int{}},
		{"equal_size", args{[]int{96, 95, 14, 56}, []int{14, 7, 88, 95}}, []int{95, 14}},
		{"unequal_size", args{[]int{41, 48, 83, 86, 17}, []int{83, 86, 6, 31, 17, 9, 48, 53}}, []int{48, 83, 86, 17}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersectSizedBrute(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersectSizedBrute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkIntersect(numItems int, testFunc func(x []int, b []int) []int, b *testing.B) {
	x := make([]int, numItems)
	y := make([]int, numItems)
	for i := range x {
		x[i], y[i] = rand.Int(), rand.Int()
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		testFunc(x, y)
	}
}

func BenchmarkIntersect10(b *testing.B)       { benchmarkIntersect(10, Intersect, b) }
func BenchmarkIntersectBrute10(b *testing.B)  { benchmarkIntersect(10, IntersectBrute, b) }
func BenchmarkIntersect20(b *testing.B)       { benchmarkIntersect(20, Intersect, b) }
func BenchmarkIntersectBrute20(b *testing.B)  { benchmarkIntersect(20, IntersectBrute, b) }
func BenchmarkIntersect50(b *testing.B)       { benchmarkIntersect(50, Intersect, b) }
func BenchmarkIntersectBrute50(b *testing.B)  { benchmarkIntersect(50, IntersectBrute, b) }
func BenchmarkIntersect100(b *testing.B)      { benchmarkIntersect(100, Intersect, b) }
func BenchmarkIntersectBrute100(b *testing.B) { benchmarkIntersect(100, IntersectBrute, b) }
func BenchmarkIntersect250(b *testing.B)      { benchmarkIntersect(250, Intersect, b) }
func BenchmarkIntersectBrute250(b *testing.B) { benchmarkIntersect(250, IntersectBrute, b) }

func TestAll(t *testing.T) {
	type args struct {
		x             []int
		qualifierFunc func(int) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"all true", args{[]int{1, 2, 3}, func(i int) bool { return i > 0 }}, true},
		{"all false", args{[]int{1, 2, 3}, func(i int) bool { return i <= 0 }}, false},
		{"one true", args{[]int{1, 2, 3}, func(i int) bool { return i <= 1 }}, false},
		{"one false", args{[]int{1, 2, 3}, func(i int) bool { return i > 1 }}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := All(tt.args.x, tt.args.qualifierFunc); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduceString(t *testing.T) {
	type args struct {
		x          []int
		reduceFunc func(acc string, val int) string
		initial    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"append",
			args{
				[]int{1, 2, 3, 4, 5},
				func(acc string, val int) string { return acc + strconv.Itoa(val) },
				"x",
			},
			"x12345",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.x, tt.args.reduceFunc, tt.args.initial); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduceInt(t *testing.T) {
	type args struct {
		x          []int
		reduceFunc func(acc int, val int) int
		initial    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"2n",
			args{
				[]int{1, 2, 3, 4, 5},
				func(acc int, val int) int { return acc + 2*val },
				99,
			},
			99 + 2*1 + 2*2 + 2*3 + 2*4 + 2*5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.x, tt.args.reduceFunc, tt.args.initial); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
