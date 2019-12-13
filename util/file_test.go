package util

import (
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		want    []int
		wantErr bool
	}{
		{
			name: "readInput",
			path: "../resources/test/testfile.txt",
			want: []int{5,15,20,-5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadIntLines(tt.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDimensions() = %v, want %v", got, tt.want)
			}
		})
	}
}
