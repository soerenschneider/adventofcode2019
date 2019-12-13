package util

import (
	"reflect"
	"testing"
)

func Test_interpreter_Execute(t *testing.T) {
	type fields struct {
		alphabet  []int64
		processed int
		pointer   int
		shutdown  bool
	}
	tests := []struct {
		fields fields
		want   []int64
	}{
		{
			fields: fields{
				alphabet: []int64{1002, 4, 3, 4, 33},
			},
			want: []int64{1002, 4, 3, 4, 99},
		},
		{
			fields: fields{
				alphabet: []int64{1101, 100, -1, 4, 0},
			},
			want: []int64{1101, 100, -1, 4, 99},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			i := &interpreter{
				alphabet:  tt.fields.alphabet,
				pointer:   tt.fields.pointer,
				shutdown:  tt.fields.shutdown,
			}
			i.Execute()
			if !reflect.DeepEqual(i.alphabet, tt.want) {
				t.Errorf("parse() got = %v, want %v", i.alphabet, tt.want)
			}
		})
	}
}

func Test_parseopcode(t *testing.T) {
	tests := []struct {
		instruction int64
		want        opcode
	}{
		{
			instruction: 2,
			want: opcode{
				opcode: 2,
				modes: []mode{0,0,0,0},
				parameters: 4,
			},
		},
		{
			instruction: 1002,
			want: opcode{
				opcode: 2,
				modes: []mode{0,1,0, 0},
				parameters: 4,
			},
		},
		{
			instruction: 12399,
			want:        opcode{opcode: 99, modes: []mode{3}, parameters: 1},
		},
		{
			instruction: 1002,
			want:        opcode{opcode: 2, modes: []mode{0, 1, 0, 0},parameters: 4},
		},
	}
	for _, tt := range tests {
		t.Run("",func(t *testing.T) {
			if got := parseOpcode(tt.instruction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseopcode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_interpreter_get(t *testing.T) {
	type fields struct {
		alphabet  []int64
		pointer   int
		shutdown  bool
	}
	type args struct {
		n int
		m opcode
	}
	tests := []struct {
		name 	string
		fields fields
		args   args
		want   int64
		wantErr bool
	}{
		
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			i := &interpreter{
				alphabet:  tt.fields.alphabet,
				pointer:   tt.fields.pointer,
				shutdown:  tt.fields.shutdown,
			}
			got := i.get(tt.args.n, tt.args.m)
			if got != tt.want {
				t.Errorf("get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_interpreter_halt(t *testing.T) {
	type fields struct {
		alphabet  []int64
		pointer   int
		shutdown  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "False",
			fields: fields{
				alphabet: []int64{1,2,3,4},
				pointer: 0,
				shutdown: false,
			},
			want: false,
		},
		{
			name: "True, shutdown set to true",
			fields: fields{
				alphabet: []int64{1,2,3,4},
				pointer: 0,
				shutdown: true,
			},
			want: true,
		},
		{
			name: "True, pointer points to end of array",
			fields: fields{
				alphabet: []int64{1,2,3,4},
				pointer: 3,
				shutdown: false,
			},
			want: true,
		},
		{
			name: "True, shutdown set to true and pointer points to end of array",
			fields: fields{
				alphabet: []int64{1,2,3,4},
				pointer: 3,
				shutdown: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &interpreter{
				alphabet:  tt.fields.alphabet,
				pointer:   tt.fields.pointer,
				shutdown:  tt.fields.shutdown,
			}
			if got := i.halt(); got != tt.want {
				t.Errorf("halt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_interpreter_c1(t *testing.T) {
	type fields struct {
		alphabet  []int64
		pointer   int
		shutdown  bool
	}
	tests := []struct {
		name   string
		fields fields
		args   opcode
		alphabet  []int64
	}{
		{
			fields: fields{
				alphabet: []int64{1,2,3,1},
				pointer: 0,
				shutdown: false,
			},
			args: opcode{
				parameters: 3,
				modes: []mode{IMM, IMM, IMM},
				opcode: 1,
			},
			alphabet: []int64{1,2,3,5},
		},
		{
			fields: fields{
				alphabet: []int64{1,2,3,1},
				pointer: 0,
				shutdown: false,
			},
			args: opcode{
				parameters: 3,
				modes: []mode{POS, POS, POS},
				opcode: 1,
			},
			alphabet: []int64{1,4,3,1},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			i := &interpreter{
				alphabet:  tt.fields.alphabet,
				pointer:   tt.fields.pointer,
				shutdown:  tt.fields.shutdown,
			}
			c1(i, &tt.args)
			if !reflect.DeepEqual(tt.alphabet, i.alphabet) {
				t.Errorf("c1() = %v, want %v", i.alphabet, tt.alphabet)
			}
		})
	}
}

func Test_interpreter_c2(t *testing.T) {
	type fields struct {
		alphabet  []int64
		pointer   int
		shutdown  bool
	}
	tests := []struct {
		name   string
		fields fields
		args   opcode
		alphabet  []int64
	}{
		{
			fields: fields{
				alphabet: []int64{1,2,3,1},
				pointer: 0,
				shutdown: false,
			},
			args: opcode{
				parameters: 3,
				modes: []mode{IMM, IMM, IMM},
				opcode: 1,
			},
			alphabet: []int64{1,2,3,6},
		},
		{
			fields: fields{
				alphabet: []int64{1,2,3,1},
				pointer: 0,
				shutdown: false,
			},
			args: opcode{
				parameters: 3,
				modes: []mode{POS, POS, POS},
				opcode: 1,
			},
			alphabet: []int64{1,3,3,1},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			i := &interpreter{
				alphabet:  tt.fields.alphabet,
				pointer:   tt.fields.pointer,
				shutdown:  tt.fields.shutdown,
			}
			c2(i, &tt.args)
			if !reflect.DeepEqual(tt.alphabet, i.alphabet) {
				t.Errorf("c2() = %v, want %v", i.alphabet, tt.alphabet)
			}
		})
	}
}

func Test_interpreter_c5(t *testing.T) {
	type fields struct {
		alphabet  []int64
		pointer   int
		shutdown  bool
	}
	tests := []struct {
		name   string
		fields fields
		args   opcode
		alphabet  []int64
		pointer int
	}{
		{
			fields: fields{
				alphabet: []int64{1,2,3,1},
				pointer: 0,
				shutdown: false,
			},
			args: opcode{
				parameters: 3,
				modes: []mode{IMM, IMM, IMM},
				opcode: 1,
			},
			alphabet: []int64{1,2,3,1},
			pointer: 3,
		},
		{
			fields: fields{
				alphabet: []int64{1,2,3,1},
				pointer: 0,
				shutdown: false,
			},
			args: opcode{
				parameters: 3,
				modes: []mode{POS, POS, POS},
				opcode: 1,
			},
			alphabet: []int64{1,2,3,1},
			pointer: 1,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			i := &interpreter{
				alphabet:  tt.fields.alphabet,
				pointer:   tt.fields.pointer,
				shutdown:  tt.fields.shutdown,
			}
			c5(i, &tt.args)
			if !reflect.DeepEqual(tt.alphabet, i.alphabet) {
				t.Errorf("c5() = %v, want %v", i.alphabet, tt.alphabet)
			}
			if !reflect.DeepEqual(tt.pointer, i.pointer) {
				t.Errorf("c5() = %v, want %v", i.pointer, tt.pointer)
			}
		})
	}
}

func Test_interpreter_c6(t *testing.T) {
	type fields struct {
		alphabet  []int64
		pointer   int
		shutdown  bool
	}
	tests := []struct {
		name   string
		fields fields
		args   opcode
		alphabet  []int64
		pointer int
	}{
		{
			fields: fields{
				alphabet: []int64{0,0,3,0},
				pointer: 0,
				shutdown: false,
			},
			args: opcode{
				parameters: 3,
				modes: []mode{IMM, IMM, IMM},
				opcode: 1,
			},
			alphabet: []int64{0,0,3,0},
			pointer: 3,
		},
		{
			fields: fields{
				alphabet: []int64{0,0,3,4},
				pointer: 0,
				shutdown: false,
			},
			args: opcode{
				parameters: 3,
				modes: []mode{POS, POS, POS},
				opcode: 1,
			},
			alphabet: []int64{0,0,3,4},
			pointer: 4,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			i := &interpreter{
				alphabet:  tt.fields.alphabet,
				pointer:   tt.fields.pointer,
				shutdown:  tt.fields.shutdown,
			}
			c6(i, &tt.args)
			if !reflect.DeepEqual(tt.alphabet, i.alphabet) {
				t.Errorf("c6() = %v, want %v", i.alphabet, tt.alphabet)
			}
			if !reflect.DeepEqual(tt.pointer, i.pointer) {
				t.Errorf("c6() = %v, want %v", i.pointer, tt.pointer)
			}
		})
	}
}

func Test_interpreter_c7(t *testing.T) {
	type fields struct {
		alphabet  []int64
		pointer   int
		shutdown  bool
	}
	tests := []struct {
		name   string
		fields fields
		args   opcode
		alphabet  []int64
		pointer int
	}{
		{
			fields: fields{
				alphabet: []int64{1,2,0,1},
				pointer: 0,
				shutdown: false,
			},
			args: opcode{
				parameters: 3,
				modes: []mode{IMM, IMM, IMM},
				opcode: 1,
			},
			alphabet: []int64{1,2,0,0},
			pointer: 0,
		},
		{
			fields: fields{
				alphabet: []int64{1,1,1,1},
				pointer: 0,
				shutdown: false,
			},
			args: opcode{
				parameters: 3,
				modes: []mode{POS, POS, POS},
				opcode: 1,
			},
			alphabet: []int64{1,0,1,1},
			pointer: 0,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			i := &interpreter{
				alphabet:  tt.fields.alphabet,
				pointer:   tt.fields.pointer,
				shutdown:  tt.fields.shutdown,
			}
			c7(i, &tt.args)
			if !reflect.DeepEqual(tt.alphabet, i.alphabet) {
				t.Errorf("c7() = %v, want %v", i.alphabet, tt.alphabet)
			}
			if !reflect.DeepEqual(tt.pointer, i.pointer) {
				t.Errorf("c7() = %v, want %v", i.pointer, tt.pointer)
			}
		})
	}
}

func Test_interpreter_c8(t *testing.T) {
	type fields struct {
		alphabet  []int64
		pointer   int
		shutdown  bool
	}
	tests := []struct {
		name   string
		fields fields
		args   opcode
		alphabet  []int64
		pointer int
	}{
		{
			fields: fields{
				alphabet: []int64{1,2,0,1},
				pointer: 0,
				shutdown: false,
			},
			args: opcode{
				parameters: 3,
				modes: []mode{IMM, IMM, IMM},
				opcode: 1,
			},
			alphabet: []int64{1,2,0,0},
			pointer: 0,
		},
		{
			fields: fields{
				alphabet: []int64{1,1,1,1},
				pointer: 0,
				shutdown: false,
			},
			args: opcode{
				parameters: 3,
				modes: []mode{POS, POS, POS},
				opcode: 1,
			},
			alphabet: []int64{1,1,1,1},
			pointer: 0,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			i := &interpreter{
				alphabet:  tt.fields.alphabet,
				pointer:   tt.fields.pointer,
				shutdown:  tt.fields.shutdown,
			}
			c8(i, &tt.args)
			if !reflect.DeepEqual(tt.alphabet, i.alphabet) {
				t.Errorf("c8() = %v, want %v", i.alphabet, tt.alphabet)
			}
			if !reflect.DeepEqual(tt.pointer, i.pointer) {
				t.Errorf("c8() = %v, want %v", i.pointer, tt.pointer)
			}
		})
	}
}
