package day14

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBuildEquation(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name    string
		args    args
		want    Formulars
		wantErr bool
	}{
		{
			args: args{
				input: []string{
					"10 ORE => 10 A",
					"1 ORE => 1 B",
					"7 A, 1 B => 1 C",
					"7 A, 1 C => 1 D",
					"7 A, 1 D => 1 E",
					"7 A, 1 E => 1 FUEL",
				},
			},
			want: Formulars{
						"A":    Formular{producesUnits: 10, needs: map[string]int{"ORE": 10}},
						"B":    Formular{producesUnits: 1, needs: map[string]int{"ORE": 1}},
						"C":    Formular{producesUnits: 1, needs: map[string]int{"A": 7, "B": 1}},
						"D":    Formular{producesUnits: 1, needs: map[string]int{"A": 7, "C": 1}},
						"E":    Formular{producesUnits: 1, needs: map[string]int{"A": 7, "D": 1}},
						"FUEL": Formular{producesUnits: 1, needs: map[string]int{"A": 7, "E": 1}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildEquation(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildEquation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildEquation() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormulars_updateRequirements(t *testing.T) {
	type args struct {
		
	}
	tests := []struct {
		name     string
		formular Formulars
		args     map[string]int
		want 	 map[string]int
	}{
		{
			formular: Formulars{
				"A": Formular{producesUnits: 10, needs: map[string]int{"ORE": 10}},
				"B": Formular{producesUnits: 1, needs: map[string]int{"A": 2}},
				"FUEL": Formular{producesUnits: 1, needs: map[string]int{"A": 5, "B": 1}},
			},
			args: map[string]int{"FUEL": 2},
			want: map[string]int{
					"FUEL": 0,
					"ORE": 20,
					"A": -6,
					"B": 0,
			},
		},
		{
			formular: Formulars{
				"A": Formular{producesUnits: 10, needs: map[string]int{"ORE": 10}},
				"B": Formular{producesUnits: 1, needs: map[string]int{"A": 2}},
				"FUEL": Formular{producesUnits: 1, needs: map[string]int{"A": 5, "B": 1}},
			},
			args: map[string]int{"FUEL": 0},
			want: map[string]int{
				"FUEL": 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.formular.updateRequirements(tt.args)
			tt.formular.updateRequirements(tt.args)

			fmt.Println(tt.want)

			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("BuildEquation() got = %v, want %v", tt.args, tt.want)
			}
		})
	}
}

func TestFormulars_CalculateFuelRequirements(t *testing.T) {
	tests := []struct {
		name     string
		formular []string
		amount   int
		want     int
	}{
		{
			formular: []string{
				"171 ORE => 8 CNZTR",
				"7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL",
				"114 ORE => 4 BHXH",
				"14 VRPVC => 6 BMBT",
				"6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL",
				"6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT",
				"15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW",
				"13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW",
				"5 BMBT => 4 WPTQ",
				"189 ORE => 9 KTJDG",
				"1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP",
				"12 VRPVC, 27 CNZTR => 2 XDBXC",
				"15 KTJDG, 12 BHXH => 5 XCVML",
				"3 BHXH, 2 VRPVC => 7 MZWV",
				"121 ORE => 7 VRPVC",
				"7 XCVML => 6 RJRHP",
				"5 BHXH, 4 VRPVC => 5 LTCX",
			},
			amount: 1,
			want: 2210736,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formular, err := BuildEquation(tt.formular)
			if err != nil {
				t.Errorf("Error while building linear equation %s", err.Error())
			}
			if got := formular.CalculateFuelRequirements(tt.amount); got != tt.want {
				t.Errorf("CalculateFuelRequirements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isDemandMet(t *testing.T) {
	tests := []struct {
		name string
		args map[string]int
		want bool
	}{
		{
			args: map[string]int {
				"A": 0,
				"ORE": 2,
			},
			want: true,
		},
		{
			args: map[string]int {
				"A": 2,
				"ORE": 2,
			},
			want: false,
		},
		{
			args: map[string]int {
				"A": 0,
				"ORE": 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isProductionFinished(tt.args); got != tt.want {
				t.Errorf("isProductionFinished() = %v, want %v", got, tt.want)
			}
		})
	}
}