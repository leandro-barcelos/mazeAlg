package maze

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/leandro-barcelos/mazeAlg/tree"
)

type Direction int

const (
	N Direction = 1
	S Direction = 2
	E Direction = 4
	W Direction = 8
)

func (d Direction) dx() int {
	switch d {
	case E:
		return 1
	case W:
		return -1
	}

	return 0
}

func (d Direction) dy() int {
	switch d {
	case S:
		return 1
	case N:
		return -1
	}

	return 0
}

func (d Direction) opposite() Direction {
	switch d {
	case E:
		return W
	case W:
		return E
	case N:
		return S
	}

	return N
}

type Edge struct {
	x   int
	y   int
	dir Direction
}

type Maze struct {
	Grid  [][]int
	sets  [][]*tree.Tree
	edges []Edge
}

func (m *Maze) Init(rows int, cols int) {
	m.Grid = make([][]int, rows)
	for i := range m.Grid {
		m.Grid[i] = make([]int, cols)
	}

	m.sets = make([][]*tree.Tree, rows)
	for i := range m.sets {
		m.sets[i] = make([]*tree.Tree, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			m.Grid[i][j] = 0
			m.sets[i][j] = new(tree.Tree)

			if i > 0 {
				m.edges = append(m.edges, Edge{
					x:   j,
					y:   i,
					dir: N,
				})
			}

			if j > 0 {
				m.edges = append(m.edges, Edge{
					x:   j,
					y:   i,
					dir: W,
				})
			}
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(m.edges), func(i, j int) {
		m.edges[i], m.edges[j] = m.edges[j], m.edges[i]
	})

	m.kruskalAlg()
}

func (m *Maze) kruskalAlg() {
	for !(len(m.edges) == 0) {
		e := m.edges[len(m.edges)-1]
		m.edges = m.edges[:len(m.edges)-1]

		nx, ny := e.x+e.dir.dx(), e.y+e.dir.dy()

		set1, set2 := m.sets[e.y][e.x], m.sets[ny][nx]

		if !set1.Is_connected(set2) {
			m.Display()
			time.Sleep(10 * time.Millisecond)

			set1.Connect(set2)
			m.Grid[e.y][e.x] |= int(e.dir)
			m.Grid[ny][nx] |= int(e.dir.opposite())
		}
	}
}

func (m Maze) Display() {
	fmt.Println(" " + strings.Repeat("_", len(m.Grid[0])))

	for _, row := range m.Grid {
		fmt.Print("|")
		for x, cell := range row {
			if cell == 0 {
				fmt.Print("\033[47m")
			}
			if cell&int(S) != 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("_")
			}

			if cell&int(E) != 0 {
				if (cell|row[x+1])&int(S) != 0 {
					fmt.Print(" ")
				} else {
					fmt.Print("_")
				}
			} else {
				fmt.Print("|")
			}
			if cell == 0 {
				fmt.Print("\033[m]")
			}
		}
		fmt.Println()
	}
}
