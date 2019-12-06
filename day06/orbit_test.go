package day06

import (
	"fmt"
	"reflect"
	"testing"
)

var test = [][2]string{
	{"B", "COM"},
	{"C", "B"},
	{"D", "C"},
	{"E", "D"},
	{"F", "E"},
	{"G", "B"},
	{"H", "G"},
	{"I", "D"},
	{"J", "E"},
	{"K", "J"},
	{"L", "K"},
}

func TestGraph_Bfs_Positive(t *testing.T) {
	g := &Graph{
	}
	a := Node{"A"}
	b := Node{"B"}
	c := Node{"C"}
	g.AddNode(&a)
	g.AddNode(&b)
	g.AddNode(&c)

	g.AddEdge(&a, &b)
	g.AddEdge(&a, &c)

	ar := g.bfs(&a)
	wantAr := map[Node]int{b: 1, c: 1}
	if !reflect.DeepEqual(ar, wantAr) {
		t.Errorf("wa = %v, want %v", ar, wantAr)
	}
}

func TestGraph_Bfs_Deep(t *testing.T) {
	g := &Graph{
	}
	a := Node{"A"}
	b := Node{"B"}
	c := Node{"C"}
	g.AddNode(&a)
	g.AddNode(&b)
	g.AddNode(&c)

	g.AddEdge(&a, &b)
	g.AddEdge(&b, &c)

	ar := g.bfs(&a)
	wantAr := map[Node]int{b: 1, c: 2}
	if !reflect.DeepEqual(ar, wantAr) {
		t.Errorf("wa = %v, want %v", ar, wantAr)
	}
}

func TestGraph_Bfs_Negative(t *testing.T) {
	g := &Graph{
	}
	a := Node{"A"}
	b := Node{"B"}
	c := Node{"C"}
	g.AddNode(&a)
	g.AddNode(&b)
	g.AddNode(&c)

	g.AddEdge(&a, &b)
	g.AddEdge(&a, &c)

	ar := g.bfs(&b)
	wantAr := make(map[Node]int)
	if !reflect.DeepEqual(ar, wantAr) {
		t.Errorf("wa = %v, want %v", ar, wantAr)
	}
}

func TestGraph_Bfs_NonExistingNode(t *testing.T) {
	g := &Graph{
	}
	a := Node{"A"}
	b := Node{"B"}
	c := Node{"C"}
	g.AddNode(&a)
	g.AddNode(&b)

	g.AddEdge(&a, &b)

	ar := g.bfs(&c)
	wantAr := make(map[Node]int)
	if !reflect.DeepEqual(ar, wantAr) {
		t.Errorf("wa = %v, want %v", ar, wantAr)
	}
}

func TestGraph_Bfs_Nil(t *testing.T) {
	g := &Graph{
	}
	ar := g.bfs(nil)
	wantAr := make(map[Node]int)
	if !reflect.DeepEqual(ar, wantAr) {
		t.Errorf("wa = %v, want %v", ar, wantAr)
	}
}

func TestGraph_Add_1(t *testing.T) {
	g := Graph{}
	input := test
	want := 7
	g.Add(input)
	nodes := g.bfs(g.GetNodeByValue("L"))
	if len(nodes) != want {
		t.Errorf("expected %d nodes, got = %d", want, len(nodes))
	}
}

func TestGraph_Add_2(t *testing.T) {
	g := Graph{}
	input := test
	want := 3
	g.Add(input)
	nodes := g.bfs(g.GetNodeByValue("D"))
	if len(nodes) != want {
		t.Errorf("expected %d nodes, got = %d", len(nodes), want)
	}

	fmt.Println(g.ConnectedNodes())
}

func TestGraph_Add_3(t *testing.T) {
	g := Graph{}
	input := test
	g.Add(input)
	want := 0
	nodes := g.bfs(g.GetNodeByValue("COM"))
	if len(nodes) != want {
		t.Errorf("expected %d nodes, got = %d", len(nodes), want)
	}
}

func TestGraph_Complete(t *testing.T) {
	g := Graph{}
	input := test
	wantSum := 42
	g.Add(input)
	connectedNodes := g.ConnectedNodes()
	if connectedNodes != wantSum {
		t.Errorf("expected %d connected nodes, got = %d", wantSum, connectedNodes)
	}
}

func TestGraph_MinDist_Indirect(t *testing.T) {
	g := Graph{}
	rawInput := ReadInput("../resources/day06/test2.txt")
	expected := 4
	nodes, _ := Parse(rawInput)
	g.Add(nodes)
	actual := g.MinDist(g.GetNodeByValue("YOU"), g.GetNodeByValue("SAN"))
	if expected != actual {
		t.Errorf("expected cost of %d, got = %d", expected, actual)
	}
}

func TestGraph_MinDist_Direct(t *testing.T) {
	g := Graph{}
	rawInput := ReadInput("../resources/day06/test2.txt")
	expected := 3
	nodes, _ := Parse(rawInput)
	g.Add(nodes)
	actual := g.MinDist(g.GetNodeByValue("YOU"), g.GetNodeByValue("D"))
	if expected != actual {
		t.Errorf("expected cost of %d, got = %d", expected, actual)
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    [][2]string
		wantErr bool
	}{
		{
			name: "Positive",
			args:  ReadInput("../resources/day06/test.txt"),
			want: [][2]string{
				{"B", "COM"},
				{"C", "B"},
				{"D", "C"},
				{"E", "D"},
				{"F", "E"},
				{"G", "B"},
				{"H", "G"},
				{"I", "D"},
				{"J", "E"},
				{"K", "J"},
				{"L", "K"},
			},
			wantErr: false,
		},
		{
			name: "Invalid input",
			args: []string{
				"COM)B)C",
			},
			want: [][2]string{
			},
			wantErr: true,
		},
		{
			name: "Empty input",
			args: []string{
			},
			want: [][2]string{
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
