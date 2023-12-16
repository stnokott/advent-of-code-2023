package stringsx

import "testing"

func TestMustAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"single", args{"1"}, 1},
		{"complex", args{"123456789"}, 123456789},
		{"zero", args{"0"}, 0},
		{"with_space", args{" 16 "}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustAtoi(tt.args.s); got != tt.want {
				t.Errorf("MustAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
