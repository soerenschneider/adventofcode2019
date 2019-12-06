package day06

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
)

const inputSeparator = ")"

type Node struct {
	name	string
}

type Graph struct {
	nodes	map[string]*Node
	edges	map[Node][]*Node
}

func (g *Graph) AddNode(n *Node) {
	if g.nodes == nil {
		g.nodes = make(map[string]*Node)
	}
	_, contains := g.nodes[n.name]; if !contains {
		g.nodes[n.name] = n	
	}
}

func (g *Graph) GetNodeByValue(name string) *Node {
	return g.nodes[name]
}

func (g *Graph) AddEdge(n, m *Node) {
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}

	g.edges[*n] = append(g.edges[*n], m)
}

func (g *Graph) bfs(n *Node) map[Node]int {
	ret := make(map[Node]int)
	if n == nil {
		return ret
	}

	queue := make([]*Node, 0, len(g.edges[*n]))
	queue = append(queue, n)
	dist := 0

	for len(queue) > 0 {
		Node := queue[0]
		queue = queue[1:]
		dist++

		for _, edge := range g.edges[*Node] {
			if _, visited := ret[*edge]; !visited {
				ret[*edge] = dist
				queue = append(queue, edge)
			}
		}
	}

	return ret
}

func (g *Graph) MinDist(n, m *Node) int {
	// We're using a directed graph – if we're lucky, the nodes are directly connected
	a := g.bfs(n)
	val, ok := a[*m]; if ok {
		// don't count initial transition
		return val - 1
	}

	// Damn, they're not, find the cheapest intersection
	b := g.bfs(m)
	min := math.MaxInt32
	
	for node, costB := range b {
		costA, found := a[node]; if found {
			x := costA + costB
			if x < min {
				min = x
			}
		}
	}

	// don't count initial transition
	return min - 2
}

func Parse(input []string) ([][2]string, error) {
	ret := make([][2]string, 0)
	
	for _, orbit := range input {
		nodes := strings.Split(orbit, inputSeparator)
		if len(nodes) != 2 {
			return ret, fmt.Errorf("expected 2 nodes, got %d", len(nodes))
		}
		edge := [2]string{nodes[1], nodes[0]}
		ret = append(ret, edge)
	}
	return ret, nil
}

func (g *Graph) Add(input [][2]string) {
	for _, nodes := range input {
		for _, node := range nodes {
			n := g.GetNodeByValue(node)
			if n == nil {
				g.AddNode(&Node{node})
			}
		}
		g.AddEdge(g.GetNodeByValue(nodes[0]), g.GetNodeByValue(nodes[1]))
	}
}

func (g *Graph) Neighbours(n *Node) int {
	return len(g.bfs(n))
}

func (g *Graph) ConnectedNodes() int {
	sum := 0
	for _, node := range g.nodes {
		sum += g.Neighbours(node)
	}
	return sum
}

func Answer06() {
	rawInput := ReadInput("resources/day06/input.txt")
	input, _ := Parse(rawInput)

	g := Graph{}
	g.Add(input)

	x := g.ConnectedNodes()
	fmt.Printf(  "Day06: %d\n", x)

	minDist := g.MinDist(g.GetNodeByValue("YOU"), g.GetNodeByValue("SAN"))
	fmt.Printf("Min dist: %d\n", minDist)
}

func ReadInput(path string) []string {
	absPath, err := filepath.Abs(path); if err != nil {
		log.Fatalf("")
	}

	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal("could not open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var ret []string

	for scanner.Scan() {
		line := scanner.Text()
		ret = append(ret, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("couldn't read file: %s", err.Error())
	}

	return ret
}