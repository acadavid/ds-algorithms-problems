package main

import ("fmt")

type Graph struct {
	vertices []Vertex
}

type Vertex struct {
	data string
	visited bool
	adj []Vertex
}

func (origV *Vertex) AddEdge(v Vertex) {
	origV.adj = append(origV.adj, v)
}


func RouteExists(graph Graph, a Vertex, b Vertex) bool {
	queue := make([]Vertex, 0)

	for _,v := range graph.vertices {
		v.visited = false
	}

	if a.data == b.data {
		return true
	}

	a.visited = true
	queue = append(queue, a)

	for len(queue) != 0 {
		s := queue[0]
		queue = queue[1:]

		for _,v := range s.adj {
			if ! (v.visited) {
				if v.data == b.data {
					return true
				}
				v.visited = true
				queue = append(queue, v)
			}
		}
		s.visited = true
	}

	return false
}

func main() {
	graph := Graph{}
	v4 := Vertex{data: "D"}
	v6 := Vertex{data: "F",adj: []Vertex{v4}}
	v2 := Vertex{data: "B", adj: []Vertex{v6}}
	v1 := Vertex{data: "A", adj: []Vertex{v2}}
	v3 := Vertex{data: "C", adj: []Vertex{v2, v4}}
	v5 := Vertex{data: "E", adj: []Vertex{v3}}
	
	graph.vertices = append(graph.vertices, v1)
	graph.vertices = append(graph.vertices, v2)
	graph.vertices = append(graph.vertices, v3)
	graph.vertices = append(graph.vertices, v4)
	graph.vertices = append(graph.vertices, v5)
	graph.vertices = append(graph.vertices, v6)

	fmt.Println(RouteExists(graph, v1, v4))
}
