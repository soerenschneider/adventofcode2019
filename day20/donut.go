package day20

import (
	"errors"
	"fmt"
	"github.com/soerenschneider/adventofcode2019/util"
)

const (
	inwards     direction = 1
	outwards    direction = -1
	openPassage           = '.'
	wall                  = '#'
	blank                 = ' '
	start                 = "AA"
	target                = "ZZ"
	// while the grid should be a grid with dimensions nxn, due to having two portals,
	// they can vary in size up to two portals
	offset				  = 2 * len(target)
)

var (
	North    = util.Coordinate{X: 0, Y: -1}
	South    = util.Coordinate{X: 0, Y: 1}
	West     = util.Coordinate{X: -1, Y: 0}
	East     = util.Coordinate{X: 1, Y: 0}
	Adjacent = []util.Coordinate{North, West, South, East}
)

type direction int
type void struct{}

type node struct {
	value     rune
	neighbors []*edge
}

type edge struct {
	name      string
	direction direction
	node      *node
}

type donut struct {
	grid       map[util.Coordinate]*node
	boardSizeX int
	boardSizeY int
}

func (f *donut) parseInput(input []string) (*node, *node) {
	f.grid = make(map[util.Coordinate]*node)
	f.collectNodes(input)
	start, end := f.findPortals(input)
	f.detectPath()
	return start, end
}

func canMove(r rune) bool {
	return r != blank && r != wall
}

func (f *donut) collectNodes(input []string) {
	for y, row := range input {
		f.boardSizeY = util.MaxInt(f.boardSizeY, len(input[y]))
		f.boardSizeX = util.MaxInt(f.boardSizeX, len(row))

		for x, rune := range row {
			if canMove(rune) {
				f.grid[util.Coordinate{X: x, Y: y}] = &node{value: rune}
			}
		}
	}
}

func (f *donut) detectPath() {
	for position, node := range f.grid {
		for _, direction := range Adjacent {
			n := position.Move(direction)

			if x, ok := f.grid[n]; ok {
				if x.value == openPassage {
					node.neighbors = append(node.neighbors, &edge{node: x})
				}
			}
		}
	}
}

func isPortal(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func (f *donut) findPortals(input []string) (*node, *node) {
	portals := make(map[string]*node)

	for curPos, node := range f.grid {
		if node.value != openPassage {
			continue
		}

		// Move around neighbouring coords 
		for _, direction := range Adjacent {

			// Move twice in same direction
			lookaheadOne := curPos.Move(direction)
			lookaheadTwo := lookaheadOne.Move(direction)

			// Check if the coords actually exist
			nodeOneStep, foundA := f.grid[lookaheadOne]
			nodeTwoSteps, foundB := f.grid[lookaheadTwo]

			// Coords must exist and be a label
			if foundA && isPortal(nodeOneStep.value) && foundB && isPortal(nodeTwoSteps.value) {
				portalName := f.determinePortalName(nodeTwoSteps, nodeOneStep, lookaheadOne, lookaheadTwo)

				if target, foundPortal := portals[portalName]; !foundPortal {
					portals[portalName] = node
				} else {
					f.wirePortals(portalName, curPos, target, node)
				}
			}
		}
	}

	return portals[start], portals[target]
}

func (f *donut) wirePortals(portalName string, curPos util.Coordinate, target *node, node *node) {
	portalOut := &edge{name: portalName, direction: outwards}
	portalIn := &edge{name: portalName, direction: inwards}

	// Determine whether we're moving in or out of the fucking donut
	if f.boardSizeX-curPos.X <= offset || curPos.X <= offset || f.boardSizeY-curPos.Y <= offset || curPos.Y <= offset {
		portalOut.node = target
		portalIn.node = node
		node.neighbors = append(node.neighbors, portalOut)
		target.neighbors = append(target.neighbors, portalIn)
	} else {
		portalOut.node = node
		portalIn.node = target
		target.neighbors = append(target.neighbors, portalOut)
		node.neighbors = append(node.neighbors, portalIn)
	}
}

func (f *donut) determinePortalName(nodeB *node, nodeA *node, lookaheadOne util.Coordinate, lookaheadTwo util.Coordinate) string {
	// Read the full label name consisting of two runes either vertically or horizontally
	if lookaheadOne.Y < lookaheadTwo.Y || lookaheadOne.X < lookaheadTwo.X {
		return fmt.Sprintf("%c%c", nodeA.value, nodeB.value)
	}
	return fmt.Sprintf("%c%c", nodeB.value, nodeA.value)
}

type QueueItem struct {
	Node  *node
	Level int
}

// do iterative bfs, let the queue in the go std library do the hard lifting. oh, wait...
func bfs(start *node, target *node, recursive bool) int {
	visited := map[QueueItem]void{QueueItem{Node: start, Level: 0}: void{}}
	queue := Init()
	queue.Append(QueueItem{Node: start, Level: 0}, 0)

	for !queue.IsEmpty() {
		item, distance, _ := queue.Pop()
		if item.Node == target && item.Level == 0 {
			return distance
		}

		for _, neighbor := range item.Node.neighbors {
			nc := QueueItem{Node: neighbor.node, Level: item.Level}
			if recursive { // Part two...
				nc.Level += int(neighbor.direction)
			}

			if _, found := visited[nc]; !found && nc.Level >= 0 {
				visited[nc] = void{}
				queue.Append(nc, distance+1)
			}
		}
	}

	return -1
}

func Answer20() int {
	input := util.ReadStringLinesFromFile("resources/day20/input.txt")
	f := donut{}
	start, end := f.parseInput(input)
	return bfs(start, end, false)
}

func Answer20b() int {
	input := util.ReadStringLinesFromFile("resources/day20/input.txt")
	f := donut{}
	start, end := f.parseInput(input)
	return bfs(start, end, true)
}

type Queue struct {
	items []QueueItem
	dist  map[QueueItem]int
}

func Init() *Queue {
	return &Queue{
		dist: make(map[QueueItem]int),
	}
}

func (q *Queue) Append(item QueueItem, dist int) {
	if _, found := q.dist[item]; !found {
		q.items = append(q.items, item)
	}

	q.dist[item] = dist
}

func (q *Queue) Pop() (QueueItem, int, error) {
	if q.IsEmpty() {
		return QueueItem{}, -1, errors.New(("queue is empty"))
	}

	item := q.items[0]
	q.items = q.items[1:]

	priority := q.dist[item]
	delete(q.dist, item)

	return item, priority, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}