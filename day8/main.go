package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	name  string
	left  string
	right string
}

type Network map[string][2]string
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

	var network Network = make(Network)
	var navigation Navigation
	line_num := 0
	for scanner.Scan() {
		line := scanner.Text()
		line_num++
		if line_num == 1 {
			navigation = Navigation(line)
		}
		if line_num > 2 {
			node := parseLineAsNode(line)
			network[node.name] = [2]string{node.left, node.right}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, "", err
	}
	return network, navigation, nil
}

func findStartNodes(nodes Network) []string {
	var startNodes []string
	for node := range nodes {
		if strings.HasSuffix(node, "A") {
			startNodes = append(startNodes, node)
		}
	}
	return startNodes
}

func solve(net Network, nav Navigation) int {
	var current_nodes []string
	current_nodes = findStartNodes(net)
	fmt.Println(current_nodes)
	return 0
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
		fmt.Println(navigation)
		fmt.Println("Part2: ", solve(network, navigation))

	}
}
