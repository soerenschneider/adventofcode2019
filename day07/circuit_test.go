package day07

import (
	"reflect"
	"testing"
)

func TestInterpreter_execute(t *testing.T) {
	type fields struct {
		alphabet  []int
		processed int
		pointer   int
		shutdown  bool
	}
	tests := []struct {
		fields fields
		want   []int
	}{
		{
			fields: fields{
				alphabet: []int{1002, 4, 3, 4, 33},
			},
			want: []int{1002, 4, 3, 4, 99},
		},
		{
			fields: fields{
				alphabet: []int{1101, 100, -1, 4, 0},
			},
			want: []int{1101, 100, -1, 4, 99},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			i := &Interpreter{
				alphabet:  tt.fields.alphabet,
				processed: tt.fields.processed,
				pointer:   tt.fields.pointer,
				shutdown:  tt.fields.shutdown,
			}
			i.execute()
			if !reflect.DeepEqual(i.alphabet, tt.want) {
				t.Errorf("parse() got = %v, want %v", i.alphabet, tt.want)
			}
		})
	}
}

func Test_parseOpcode(t *testing.T) {
	tests := []struct {
		instruction int
		want        Opcode
	}{
		{
			instruction: 2,
			want: Opcode{
				opcode: 2,
				modes: []mode{0,0,0},
				parameters: 3,
			},
		},
		{
			instruction: 1002,
			want: Opcode{
				opcode: 2,
				modes: []mode{0,1,0},
				parameters: 3,
			},
		},
		{
			instruction: 12399,
			want:        Opcode{opcode: 99, modes: nil, parameters: 0},
		},
		{
			instruction: 1002,
			want:        Opcode{opcode: 2, modes: []mode{0, 1, 0},parameters: 3},
		},
	}
	for _, tt := range tests {
		t.Run("",func(t *testing.T) {
			if got := parseOpcode(tt.instruction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseOpcode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterpreter_get(t *testing.T) {
	type fields struct {
		alphabet  []int
		processed int
		pointer   int
		shutdown  bool
	}
	type args struct {
		n int
		m mode
	}
	tests := []struct {
		name 	string
		fields fields
		args   args
		want   int
		wantErr bool
	}{
		{
			name: "Test args 2 and IMM",
			fields: fields{
				alphabet: []int{1,2,3,4},
				processed: 0,
				pointer: 0,
				shutdown: false,
			},
			args: args{
					n: 1,
					m: IMM,
			}, 
			want: 2,
			wantErr: false,
		},
		{
			name: "Test args 2 and POS",
			fields: fields{
				alphabet: []int{1,2,3,4},
				processed: 0,
				pointer: 0,
				shutdown: false,
			},
			args: args{
				n: 1,
				m: POS,
			},
			want: 3,
			wantErr: false,
		},
		{
			name: "Boundary",
			fields: fields{
				alphabet: []int{1,2,3,4},
				processed: 0,
				pointer: 0,
				shutdown: false,
			},
			args: args{
				n: 5,
				m: POS,
			},
			want: -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			i := &Interpreter{
				alphabet:  tt.fields.alphabet,
				processed: tt.fields.processed,
				pointer:   tt.fields.pointer,
				shutdown:  tt.fields.shutdown,
			}
			got, err := i.get(tt.args.n, tt.args.m) 
			if got != tt.want {
				t.Errorf("get() = %v, want %v", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestInterpreter_halt(t *testing.T) {
	type fields struct {
		alphabet  []int
		processed int
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
				alphabet: []int{1,2,3,4},
				processed: 0,
				pointer: 0,
				shutdown: false,
			},
			want: false,
		},
		{
			name: "True, shutdown set to true",
			fields: fields{
				alphabet: []int{1,2,3,4},
				processed: 0,
				pointer: 0,
				shutdown: true,
			},
			want: true,
		},
		{
			name: "True, pointer points to end of array",
			fields: fields{
				alphabet: []int{1,2,3,4},
				processed: 0,
				pointer: 3,
				shutdown: false,
			},
			want: true,
		},
		{
			name: "True, shutdown set to true and pointer points to end of array",
			fields: fields{
				alphabet: []int{1,2,3,4},
				processed: 0,
				pointer: 3,
				shutdown: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interpreter{
				alphabet:  tt.fields.alphabet,
				processed: tt.fields.processed,
				pointer:   tt.fields.pointer,
				shutdown:  tt.fields.shutdown,
			}
			if got := i.halt(); got != tt.want {
				t.Errorf("halt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterpreter_c1(t *testing.T) {
	type fields struct {
		alphabet  []int
		processed int
		pointer   int
		shutdown  bool
	}
	tests := []struct {
		name   string
		fields fields
		args   Opcode
		alphabet []int
	}{
		{
			fields: fields{
				alphabet: []int{1,2,3,1},
				processed: 0,
				pointer: 0,
				shutdown: false,
			},
			args: Opcode{
				parameters: 3,
				modes: []mode{IMM, IMM, IMM},
				opcode: 1,
			},
			alphabet: []int{1,5,3,1},
		},
		{
			fields: fields{
				alphabet: []int{1,2,3,1},
				processed: 0,
				pointer: 0,
				shutdown: false,
			},
			args: Opcode{
				parameters: 3,
				modes: []mode{POS, POS, POS},
				opcode: 1,
			},
			alphabet: []int{1,4,3,1},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			i := &Interpreter{
				alphabet:  tt.fields.alphabet,
				processed: tt.fields.processed,
				pointer:   tt.fields.pointer,
				shutdown:  tt.fields.shutdown,
			}
			i.c1(tt.args)
			if !reflect.DeepEqual(tt.alphabet, i.alphabet) {
				t.Errorf("c1() = %v, want %v", i.alphabet, tt.alphabet)
			}
		})
	}
}

func TestInterpreter_c2(t *testing.T) {
	type fields struct {
		alphabet  []int
		processed int
		pointer   int
		shutdown  bool
	}
	tests := []struct {
		name   string
		fields fields
		args   Opcode
		alphabet []int
	}{
		{
			fields: fields{
				alphabet: []int{1,2,3,1},
				processed: 0,
				pointer: 0,
				shutdown: false,
			},
			args: Opcode{
				parameters: 3,
				modes: []mode{IMM, IMM, IMM},
				opcode: 1,
			},
			alphabet: []int{1,6,3,1},
		},
		{
			fields: fields{
				alphabet: []int{1,2,3,1},
				processed: 0,
				pointer: 0,
				shutdown: false,
			},
			args: Opcode{
				parameters: 3,
				modes: []mode{POS, POS, POS},
				opcode: 1,
			},
			alphabet: []int{1,3,3,1},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			i := &Interpreter{
				alphabet:  tt.fields.alphabet,
				processed: tt.fields.processed,
				pointer:   tt.fields.pointer,
				shutdown:  tt.fields.shutdown,
			}
			i.c2(tt.args)
			if !reflect.DeepEqual(tt.alphabet, i.alphabet) {
				t.Errorf("c2() = %v, want %v", i.alphabet, tt.alphabet)
			}
		})
	}
}

func TestInterpreter_c5(t *testing.T) {
	type fields struct {
		alphabet  []int
		processed int
		pointer   int
		shutdown  bool
	}
	tests := []struct {
		name   string
		fields fields
		args   Opcode
		alphabet []int
		pointer int
	}{
		{
			fields: fields{
				alphabet: []int{1,2,3,1},
				processed: 0,
				pointer: 0,
				shutdown: false,
			},
			args: Opcode{
				parameters: 3,
				modes: []mode{IMM, IMM, IMM},
				opcode: 1,
			},
			alphabet: []int{1,2,3,1},
			pointer: 3,
		},
		{
			fields: fields{
				alphabet: []int{1,2,3,1},
				processed: 0,
				pointer: 0,
				shutdown: false,
			},
			args: Opcode{
				parameters: 3,
				modes: []mode{POS, POS, POS},
				opcode: 1,
			},
			alphabet: []int{1,2,3,1},
			pointer: 1,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			i := &Interpreter{
				alphabet:  tt.fields.alphabet,
				processed: tt.fields.processed,
				pointer:   tt.fields.pointer,
				shutdown:  tt.fields.shutdown,
			}
			i.c5(tt.args)
			if !reflect.DeepEqual(tt.alphabet, i.alphabet) {
				t.Errorf("c5() = %v, want %v", i.alphabet, tt.alphabet)
			}
			if !reflect.DeepEqual(tt.pointer, i.pointer) {
				t.Errorf("c5() = %v, want %v", i.pointer, tt.pointer)
			}
		})
	}
}

func TestInterpreter_c6(t *testing.T) {
	type fields struct {
		alphabet  []int
		processed int
		pointer   int
		shutdown  bool
	}
	tests := []struct {
		name   string
		fields fields
		args   Opcode
		alphabet []int
		pointer int
	}{
		{
			fields: fields{
				alphabet: []int{0,0,3,0},
				processed: 0,
				pointer: 0,
				shutdown: false,
			},
			args: Opcode{
				parameters: 3,
				modes: []mode{IMM, IMM, IMM},
				opcode: 1,
			},
			alphabet: []int{0,0,3,0},
			pointer: 3,
		},
		{
			fields: fields{
				alphabet: []int{0,0,3,4},
				processed: 0,
				pointer: 0,
				shutdown: false,
			},
			args: Opcode{
				parameters: 3,
				modes: []mode{POS, POS, POS},
				opcode: 1,
			},
			alphabet: []int{0,0,3,4},
			pointer: 4,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			i := &Interpreter{
				alphabet:  tt.fields.alphabet,
				processed: tt.fields.processed,
				pointer:   tt.fields.pointer,
				shutdown:  tt.fields.shutdown,
			}
			i.c6(tt.args)
			if !reflect.DeepEqual(tt.alphabet, i.alphabet) {
				t.Errorf("c6() = %v, want %v", i.alphabet, tt.alphabet)
			}
			if !reflect.DeepEqual(tt.pointer, i.pointer) {
				t.Errorf("c6() = %v, want %v", i.pointer, tt.pointer)
			}
		})
	}
}

func TestInterpreter_c7(t *testing.T) {
	type fields struct {
		alphabet  []int
		processed int
		pointer   int
		shutdown  bool
	}
	tests := []struct {
		name   string
		fields fields
		args   Opcode
		alphabet []int
		pointer int
	}{
		{
			fields: fields{
				alphabet: []int{1,2,0,1},
				processed: 0,
				pointer: 0,
				shutdown: false,
			},
			args: Opcode{
				parameters: 3,
				modes: []mode{IMM, IMM, IMM},
				opcode: 1,
			},
			alphabet: []int{1,0,0,1},
			pointer: 4,
		},
		{
			fields: fields{
				alphabet: []int{1,1,1,1},
				processed: 0,
				pointer: 0,
				shutdown: false,
			},
			args: Opcode{
				parameters: 3,
				modes: []mode{POS, POS, POS},
				opcode: 1,
			},
			alphabet: []int{1,0,1,1},
			pointer: 4,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			i := &Interpreter{
				alphabet:  tt.fields.alphabet,
				processed: tt.fields.processed,
				pointer:   tt.fields.pointer,
				shutdown:  tt.fields.shutdown,
			}
			i.c7(tt.args)
			if !reflect.DeepEqual(tt.alphabet, i.alphabet) {
				t.Errorf("c7() = %v, want %v", i.alphabet, tt.alphabet)
			}
			if !reflect.DeepEqual(tt.pointer, i.pointer) {
				t.Errorf("c7() = %v, want %v", i.pointer, tt.pointer)
			}
		})
	}
}

func TestInterpreter_c8(t *testing.T) {
	type fields struct {
		alphabet  []int
		processed int
		pointer   int
		shutdown  bool
	}
	tests := []struct {
		name   string
		fields fields
		args   Opcode
		alphabet []int
		pointer int
	}{
		{
			fields: fields{
				alphabet: []int{1,2,1,1},
				processed: 0,
				pointer: 0,
				shutdown: false,
			},
			args: Opcode{
				parameters: 3,
				modes: []mode{IMM, IMM, IMM},
				opcode: 1,
			},
			alphabet: []int{1,0,1,1},
			pointer: 4,
		},
		{
			fields: fields{
				alphabet: []int{1,1,1,1},
				processed: 0,
				pointer: 0,
				shutdown: false,
			},
			args: Opcode{
				parameters: 3,
				modes: []mode{POS, POS, POS},
				opcode: 1,
			},
			alphabet: []int{1,1,1,1},
			pointer: 4,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			i := &Interpreter{
				alphabet:  tt.fields.alphabet,
				processed: tt.fields.processed,
				pointer:   tt.fields.pointer,
				shutdown:  tt.fields.shutdown,
			}
			i.c8(tt.args)
			if !reflect.DeepEqual(tt.alphabet, i.alphabet) {
				t.Errorf("c8() = %v, want %v", i.alphabet, tt.alphabet)
			}
			if !reflect.DeepEqual(tt.pointer, i.pointer) {
				t.Errorf("c8() = %v, want %v", i.pointer, tt.pointer)
			}
		})
	}
}

func TestPermutate(t *testing.T) {
	tests := []struct {
		name string
		alphabet []int
		want [][]int
	}{
		{
			name: "Enumerate",
			alphabet: []int{1,2,3},
			want: [][]int{
				{1,2,3},
				{2,3,1},
				{3,1,2},
				{1,3,2},
				{2,1,3},
				{3,2,1},
			},
		},
		{
			name: "Enumerate",
			alphabet: []int{1,2},
			want: [][]int{
				{1,2},
				{2,1},
			},
		},
		{
			name: "Enumerate",
			alphabet: []int{1},
			want: [][]int{
				{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Permutate(tt.alphabet); 
			if len(got) != len(tt.want) {
				t.Errorf("Permutate() = %v, want %v", got, tt.want)
			}
			
			for _, o := range(tt.want) {
				exists := false
				for _, i := range(got) {
					if reflect.DeepEqual(i, o) {
						exists = true
					}
				}
				if !exists {
					t.Errorf("Permutate() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestPhases1(t *testing.T) {
	type args struct {
		input    []int
		alphabet []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				input: []int{3,1,2,4,0},
				alphabet: []int{3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0},
			},
			want: 43210,
		},
		{
			args: args{
				input: []int{0,1,2,3,4},
				alphabet: []int{3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0},
			},
			want: 54321,
		},
		{
			args: args{
				input: []int{1,0,4,3,2},
				alphabet: []int{3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0},
			},
			want: 65210,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Phases(tt.args.input, tt.args.alphabet); got != tt.want {
				t.Errorf("Phases() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestPhases2(t *testing.T) {
	type args struct {
		input    []int
		alphabet []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				input: []int{9,8,7,6,5},
				alphabet: []int{3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5},
			},
			want: 139629729,
		},
		{
			args: args{
				input: []int{5,6,7,8,9},
				alphabet: []int{3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10},
			},
			want: 18216,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Phases(tt.args.input, tt.args.alphabet); got != tt.want {
				t.Errorf("Phases() = %v, want %v", got, tt.want)
			}
		})
	}
}
