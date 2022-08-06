package main

import "fmt"

type Node struct{}
type Edge struct {
	weight int
}

type Graph struct {
	Nodes       map[string]Node
	Edges       map[string]map[string]Edge
	edgesAmount int
}

func NewGraph() *Graph {
	return &Graph{
		Nodes:       make(map[string]Node),
		Edges:       make(map[string]map[string]Edge),
		edgesAmount: 0,
	}
}

func (g *Graph) AddGetNode(value string) string {
	g.Nodes[value] = Node{}
	return value
}

func (g *Graph) AddEdge(u, v string) {

	if _, ok := g.Edges[u][v]; ok {
		g.Edges[u][v] = Edge{g.Edges[u][v].weight + 1}
	} else {
		if _, ok := g.Edges[u]; !ok {
			g.Edges[u] = make(map[string]Edge)
		}
		g.Edges[u][v] = Edge{weight: 1}
		g.edgesAmount++
	}

}
func (g *Graph) PrintEdges() {
	sum := 0
	for _, edg := range g.Edges {
		sum += len(edg)
	}
	fmt.Println(sum)
	for u, m := range g.Edges {
		for v := range m {
			fmt.Printf("%s %s %d\n", u, v, g.Edges[u][v].weight)
		}
	}
}
func main() {
	mygraph := NewGraph()
	var T int
	fmt.Scanf("%d\n", &T)
	for i := 0; i < T; i++ {
		var str string
		fmt.Scanf("%d\n", &str)
		lastNode := mygraph.AddGetNode(str[0:3])
		for i := 1; (i + 2) < len(str); i++ {
			currNode := mygraph.AddGetNode(str[i : i+3])
			mygraph.AddEdge(lastNode, currNode)
			lastNode = currNode
		}
	}
	fmt.Println(len(mygraph.Nodes))
	mygraph.PrintEdges()

}
