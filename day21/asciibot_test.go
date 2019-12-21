package day21

import (
	"reflect"
	"strings"
	"testing"
)

func Test_isAscii(t *testing.T) {
	tests := []struct {
		name string
		args int64
		want bool
	}{
		{
			args: -1,
			want: false,
		},
		{
			args: 0,
			want: true,
		},
		{
			args: 32,
			want: true,
		},
		{
			args: 128,
			want: true,
		},
		{
			args: 129,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAscii(tt.args); got != tt.want {
				t.Errorf("isAscii() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_asciiBot_format(t *testing.T) {
	type fields struct {
		mode   mode
		script []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{
				mode:   "TestMode",
				script: []string{"A", "B", "C"},
			},
			want: "A\nB\nC\nTestMode\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &asciiBot{
				mode:   tt.fields.mode,
				script: tt.fields.script,
			}
			if got := b.format(); got != tt.want {
				t.Errorf("format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAsciiBot(t *testing.T) {
	type args struct {
		mode   mode
		script []string
	}
	tests := []struct {
		name    string
		args    args
		want    *asciiBot
		wantErr bool
	}{
		{
			args: args{
				mode:   "mode",
				script: []string{"A", "B", "C"},
			},
			want: &asciiBot{mode: "mode", script: []string{"A", "B", "C"}},
		},
		{
			args: args{
				mode:   "mode",
				script: []string{"A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAsciiBot(tt.args.mode, tt.args.script)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAsciiBot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAsciiBot() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_asciiBot_processOutput(t *testing.T) {
	type fields struct {
		mode         mode
		script       []string
		outputBuffer strings.Builder
	}
	type args struct {
		botOutput int64
		botInput  chan int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			fields: fields{
				mode: "mode",
				script: []string{"A","B","C"},
				outputBuffer: strings.Builder{},
			},
			args: args {
				botOutput: 12,
				botInput: make(chan int64),
			},
			want: -1,
		},
		{
			fields: fields{
				mode: "mode",
				script: []string{"A","B","C"},
				outputBuffer: strings.Builder{},
			},
			args: args {
				botOutput: 129,
				botInput: make(chan int64),
			},
			want: 129,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &asciiBot{
				mode:         tt.fields.mode,
				script:       tt.fields.script,
				outputBuffer: tt.fields.outputBuffer,
			}
			got := b.processOutput(tt.args.botOutput, tt.args.botInput)
			if got != tt.want {
				t.Errorf("processOutput() got = %v, want %v", got, tt.want)
			}
		})
	}
}