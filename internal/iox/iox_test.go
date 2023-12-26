package iox

import (
	"reflect"
	"testing"
)

func TestReadLines(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"unix", args{"testdata/input_unix.txt"}, []string{"I have", "multiple lines", "with", "unix", "line terminator"}, false},
		{"windows", args{"testdata/input_windows.txt"}, []string{"I have", "multiple lines", "with", "windows", "line terminator"}, false},
		{"additional newline", args{"testdata/input_with_newline.txt"}, []string{"I have", "multiple lines", "and an", "additional", "line terminator", "at the end"}, false},
		{"empty lines", args{"testdata/input_empty_lines.txt"}, []string{"I have", "multiple lines", "with", "", "empty", "", "lines", "", "inbetween"}, false},
		{"not found", args{"testdata/i_dont_exist.txt"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadLines(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
