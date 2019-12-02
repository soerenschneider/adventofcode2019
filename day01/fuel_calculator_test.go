package day01

import (
	"testing"
)

func TestProcessInput(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		want    int
		wantErr bool
	}{
		{
			name: "Basic",
			input: []int{12, 14, 1969, 100756},
			want: 34241,
			wantErr: false,
		},
		{
			name: "Extremes: nil",
			input: nil,
			want: -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestRead(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		want    []int
		wantErr bool
	}{
		{
			name: "readInput bonjour",
			path: "../resources/test/testfile.txt",
			want: []int{5},
			wantErr: false,
		},
		{
			name: "readInput non existent file",
			path: "../resources/test/non-existent.txt",
			want: nil,
			wantErr: true,
		},
		{
			name: "readInput invalid file",
			path: "../test/testfile-invalid.txt",
			want: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadModules(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("readInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !equal(got, tt.want) {
				t.Errorf("readInput() got = %v, want %v", got, tt.want)
			}
		})
	}
}
