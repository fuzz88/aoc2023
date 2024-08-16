package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Node struct {
	name  string
	left  string
	right string
}

type Network []Node
type Navigation string

func parseLineAsNode(line string) Node {
	first_split := strings.Split(line, "=")
	second_split := strings.Split(first_split[1], ",")
	return Node{
		name:  strings.Trim(first_split[0], " "),
		left:  strings.Trim(second_split[0], " ")[1:],
		right: strings.Trim(second_split[1][:len(second_split[1])-1], " "),
	}
}

func readNetworkFromFile(filePath string) (Network, Navigation, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", err
	}
	scanner := bufio.NewScanner(file)

	var network Network
	var navigation Navigation
	line_num := 0
	for scanner.Scan() {
		line := scanner.Text()
		line_num++
		if line_num == 1 {
			navigation = Navigation(line)
		}
		if line_num > 2 {
			network = append(network, parseLineAsNode(line))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, "", err
	}
	return network, navigation, nil
}

func nextNodeIndex(current_instr string, network Network, current_node Node) int {
	if current_instr == "R" {
		next_index, found := slices.BinarySearchFunc(network, current_node, func(a Node, b Node) int {
			if a.name > b.right{
				return 1
			}
			if a.name < b.right{
				return -1
			}
			if a.name == b.right {
				return 0
			}
			return 0
		})
		if !found {
			panic("not_found")
		}
		return next_index
	}
	if current_instr == "L" {
		next_index, found := slices.BinarySearchFunc(network, current_node, func(a Node, b Node) int {
			if a.name > b.left {
				return 1
			}
			if a.name < b.left {
				return -1
			}
			if a.name == b.left {
				return 0
			}
			return 0
		})
		if !found {
			panic("not_found")
		}
		return next_index
	}
	return -1
}

func solvePart1(network Network, nav Navigation) int {
	start := "AAA"
	end := "ZZZ"
	start_index := slices.IndexFunc(network, func(node Node) bool {
		return node.name == start
	})

	current_node := network[start_index]
	steps := 0

	for i := 0; ; i++ {
		steps++
		current_instr := string(nav[i%len(nav)])
		next_index := nextNodeIndex(current_instr, network, current_node)
		current_node = network[next_index]
		if current_node.name == end {
			return steps
		}
	}
}

// Define a generic filter function
func filter[T any](slice []T, fn func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

func solvePart2(net Network, nav Navigation) int {
	var current_nodes []Node
	current_nodes = filter(net, func(node Node) bool {
		return string(node.name[len(node.name)-1]) == "A"
	})
	fmt.Println(current_nodes)
	steps := 0
	for i := 0; ; i++ {
		current_instr := string(nav[i%len(nav)])
		steps++
		for j := 0; j < len(current_nodes); j++ {
			next_index := nextNodeIndex(current_instr, net, current_nodes[j])
			current_nodes[j] = net[next_index]
		}
		if !slices.ContainsFunc(current_nodes, func(node Node) bool {
			end_of_name := node.name[len(node.name)-1]
			return string(end_of_name) != "Z"
		}) {
			return steps
		} 
	}
}

func main() {
	fmt.Printf("\nAOC-2023 Day8 Solution\n\n")
	args := os.Args[1:]
	for _, filePath := range args {
		fmt.Println(filePath)
		network, navigation, err := readNetworkFromFile(filePath)
		if err != nil {
			panic(err)
		}
		slices.SortFunc(network, func(a Node, b Node) int {
			if a.name > b.name {
				return 1
			}
			if a.name < b.name {
				return -1
			}
			if a.name == b.name {
				return 0
			}
			return 0
		})
		fmt.Println(navigation)
		//fmt.Println("Part1: ", solvePart1(network, navigation))

		fmt.Println("Part2: ", solvePart2(network, navigation))

	}
}
