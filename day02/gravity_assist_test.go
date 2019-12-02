package day02

import (
	"reflect"
	"testing"
)

func TestProcessOpcode(t *testing.T) {
	tests := []struct {
		name    string
		opcode  []int
		want    []int
		wantErr bool
	}{
		{
			name: "Example 1",
			opcode: []int{1,0,0,0,99},
			want: []int{2,0,0,0,99},
			wantErr: false,
		},
		{
			name: "Example 2",
			opcode: []int{2,3,0,3,99},
			want: []int{2,3,0,6,99},
			wantErr: false,
		},
		{
			name: "Example 3",
			opcode: []int{2,4,4,5,99,0},
			want: []int{2,4,4,5,99,9801},
			wantErr: false,
		},
		{
			name: "Example 4",
			opcode: []int{1,1,1,4,99,5,6,0,99},
			want: []int{30,1,1,4,2,5,6,0,99},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ProcessInput(tt.opcode, tt.opcode[1], tt.opcode[2])
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessInput() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessOpcodeExtremes(t *testing.T) {
	tests := []struct {
		name    string
		opcode  []int
		noun 	int
		verb	int
		want    []int
		wantErr bool
	}{
		{
			name: "Extremes",
			opcode: []int{0,0,0,0,0},
			want: []int{0,0,0,0,0},
			noun: 0,
			verb: 0,
			wantErr: false,
		},
		{
			name: "Extremes",
			opcode: []int{0},
			want: nil,
			noun: 0,
			verb: 0,
			wantErr: true,
		},
		{
			name: "Extremes",
			opcode: nil,
			want: nil,
			noun: 0,
			verb: 0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ProcessInput(tt.opcode, tt.noun, tt.verb)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessInput() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_halt(t *testing.T) {
	type fields struct {
		instruction int
		param1      int
		param2      int
		pos         int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Positive",
			fields: fields{instruction: 99},
			want: true,
		},
		{
			name: "Negative",
			fields: fields{instruction: 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Operation{
				instruction: tt.fields.instruction,
				param1:      tt.fields.param1,
				param2:      tt.fields.param2,
				param3:      tt.fields.pos,
			}
			if got := o.halt(); got != tt.want {
				t.Errorf("halt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_apply(t *testing.T) {
	type fields struct {
		instruction int
		param1      int
		param2      int
		pos         int
	}

	tests := []struct {
		name    string
		fields  fields
		input    []int
		want     []int
		wantErr bool
	}{
		{
			name: "Input array too small",
			fields: fields{
					instruction: InstrAdd,
					param1: 1,
					param2: 2,
					pos: 3,
				},

				input: []int{1},
				want: []int{1},
				wantErr: true,
		},
		{
			name: "Simple pass add",
			fields: fields{
				instruction: InstrAdd,
				param1: 1,
				param2: 2,
				pos: 3,
			},

			input: []int{1, 1, 3, 2, 5},
			want:  []int{1, 1, 3, 4, 5},
			wantErr: false,
		},
		{
			name: "Simple pass mult",
			fields: fields{
				instruction: InstrMult,
				param1: 1,
				param2: 2,
				pos: 3,
			},

			input: []int{2, 1, 3, 2, 5},
			want:  []int{2, 1, 3, 3, 5},
			wantErr: false,
		},
		{
			name: "Invalid instruction",
			fields: fields{
				instruction: 0,
				param1: 1,
				param2: 2,
				pos: 3,
			},

			input: []int{1, 1, 3, 2, 5},
			want: []int{1, 1, 3, 2, 5},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Operation{
				instruction: tt.fields.instruction,
				param1:      tt.fields.param1,
				param2:      tt.fields.param2,
				param3:      tt.fields.pos,
			}
			if err := o.apply(tt.input); (err != nil) != tt.wantErr {
				t.Errorf("apply() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.input, tt.want) {
				t.Errorf("ProcessInput() got = %v, want %v", tt.input, tt.want)
			}
		})
	}
}

func Test_genOpcode(t *testing.T) {
	type args struct {
		input []int
		index int
	}
	tests := []struct {
		name string
		args args
		want Operation
	}{
		{
			name: "Extremes",
			args: args{input: nil, index: 5},
			want: Operation{},
		},
		{
			name: "Extremes",
			args: args{input: []int{}, index: 5},
			want: Operation{},
		},
		{
			name: "Simple",
			args: args{input: []int{1, 2, 3, 4, 5}, index: 0},
			want: Operation{instruction:1, param1: 2, param2: 3, param3: 4},
		},
		{
			name: "Operation with no parameters",
			args: args{input: []int{1}, index: 0},
			want: Operation{instruction:1},
		},
		{
			name: "Operation with no parameters",
			args: args{input: []int{1, 2, 3, 4, 5, 6, 7, 8}, index: 4},
			want: Operation{instruction:5, param1: 6, param2: 7, param3:8 },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractOperation(tt.args.input, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}