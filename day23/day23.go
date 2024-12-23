package day23

import (
	"slices"
	"strconv"
	"strings"

	"github.com/rindlow/aoc-utils"
)

type NodeSet = map[string]bool

func readGraph(filename string) (graph map[string]NodeSet) {
	graph = make(map[string]NodeSet)
	for _, line := range utils.ReadLines(filename) {
		a := line[:2]
		b := line[3:]
		_, ok := graph[a]
		if !ok {
			graph[a] = make(NodeSet)
		}
		graph[a][b] = true
		_, ok = graph[b]
		if !ok {
			graph[b] = make(NodeSet)
		}
		graph[b][a] = true
	}
	return
}

func interconnectedComputers(filename string) int {
	graph := readGraph(filename)
	triangles := make(map[[3]string]bool)
	for node, edges := range graph {
		for edge := range edges {
			for e := range edges {
				if graph[edge][e] && (node[0] == 't' || edge[0] == 't' || e[0] == 't') {
					triangle := []string{node, edge, e}
					slices.Sort(triangle)
					triangles[[3]string(triangle)] = true
				}
			}
		}
	}
	return len(triangles)
}

func password(filename string) string {
	graph := readGraph(filename)
	cliques := []NodeSet{}
ALLNODES:
	for node := range graph {
		for _, clique := range cliques {
			if clique[node] {
				continue ALLNODES
			}
		}

		clique := make(NodeSet)
		clique[node] = true
	ALLV:
		for v := range graph {
			for c := range clique {
				if !graph[c][v] {
					continue ALLV
				}
			}
			clique[v] = true
		}
		cliques = append(cliques, clique)
	}
	var maximal NodeSet
	maxLen := 0
	for _, clique := range cliques {
		if len(clique) > maxLen {
			maxLen = len(clique)
			maximal = clique
		}
	}
	nodes := []string{}
	for node := range maximal {
		nodes = append(nodes, node)
	}
	slices.Sort(nodes)
	return strings.Join(nodes, ",")
}

func Part1(filename string) string {
	return strconv.Itoa(interconnectedComputers(filename))
}

func Part2(filename string) string {
	return password(filename)
}
