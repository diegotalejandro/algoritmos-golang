package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	ID                  int
	visited             bool
	edges, reverseEdges []int
}

type graph struct {
	nodes map[int]*node
}

func newGraph() *graph {
	var g graph
	g.nodes = make(map[int]*node)
	return &g
}

func newNode() *node {
	var n node
	n.ID = -1
	return &n
}

func (g *graph) addEdge(t, h int) {
	if _, ok := g.nodes[t]; !ok {
		panic("No edge tail")
	}
	if _, ok := g.nodes[h]; !ok {
		panic("No edge head")
	}
	g.nodes[t].edges = append(g.nodes[t].edges, h)
	g.nodes[h].reverseEdges = append(g.nodes[h].reverseEdges, t)
}

func (g *graph) addNode(label int) bool {
	if _, ok := g.nodes[label]; !ok {
		n := newNode()
		g.nodes[label] = n
		return true
	}
	return false
}

func (g *graph) resetVisited() {
	for _, n := range g.nodes {
		n.visited = false
	}
}
func (g *graph) createFinishingOrder() []*node {
	g.resetVisited()
	t := make([]*node, 0, len(g.nodes))
	for _, v := range g.nodes {
		if v.visited == false {
			dfsAssignFinishingNumber(v, g, &t)
		}
	}
	return t
}

func dfsAssignFinishingNumber(n *node, g *graph, t *[]*node) {
	n.visited = true
	for _, neighborLabel := range n.reverseEdges {
		if g.nodes[neighborLabel].visited == false {
			dfsAssignFinishingNumber(g.nodes[neighborLabel], g, t)
		}
	}
	(*t) = append(*t, n)
}

func dfsMarkScc(n *node, g *graph, s int) {
	n.visited = true
	n.ID = s
	for _, neighborLabel := range n.edges {
		if g.nodes[neighborLabel].visited == false {
			dfsMarkScc(g.nodes[neighborLabel], g, s)
		}
	}
}

func (g *graph) generateScc() { //
	p := 0
	fo := g.createFinishingOrder()
	g.resetVisited()
	for i := len(fo) - 1; i >= 0; i-- {
		n := fo[i]
		if n.visited == false {
			s := p
			p++
			dfsMarkScc(n, g, s)
		}
	}
}

func (g *graph) getSccs() {
	a := make(map[int][]int)
	for k, n := range g.nodes {
		if n.ID == -1 {
			panic(fmt.Sprintf("Error for %d.", n.ID))
		}
		temp := a[n.ID]
		temp = append(temp, k)
		a[n.ID] = temp
	}
	var slice []int
	for i := 0; i < N; i++ {
		slice = append(slice, i+1)
	}
	for _, i := range a {
		for _, j := range i {
			for k := 0; k < N; k++ {
				if slice[k] == j {
					slice[k] = 0
				}
			}
			fmt.Print(j, " ")
		}
		fmt.Println("")
	}
	for i := 0; i < N; i++ {
		if slice[i] != 0 {
			fmt.Println(slice[i])
		}
	}
}

var N int

func main() {
	g := newGraph()
	f, err := os.Open("texto_sampled/Texto.txt") //<---------Ubicacion Texto
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		iS := strings.Fields(line)
		if len(iS) != 2 && i != 0 {
			panic(fmt.Sprintf("Error line : %v", iS))
		}
		if i != 0 {
			t, err1 := strconv.Atoi(iS[0])
			if err1 != nil {
				panic(err1)
			}
			h, err2 := strconv.Atoi(iS[1])
			if err2 != nil {
				panic(err2)
			}

			g.addNode(t)
			g.addNode(h)
			g.addEdge(t, h)
		}
		if i == 0 {
			n, err3 := strconv.Atoi(iS[0])
			if err3 != nil {
				panic(err3)
			}
			N = n
			i++
		}
	}
	g.generateScc()
	g.getSccs()
}
