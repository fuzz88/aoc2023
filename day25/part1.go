package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput(filename string) (map[string][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	graph := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		component := parts[0]
		connections := strings.Split(parts[1], " ")
		graph[component] = connections
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return graph, nil
}

func disconnectWires(graph map[string][]string, disconnectCount int) (int, error) {
	resetButton := "reset"
	groupSizes := make([]int, 0)

	var dfs func(component string) int
	dfs = func(component string) int {
		size := 1
		for _, neighbor := range graph[component] {
			if neighbor != resetButton {
				size += dfs(neighbor)
			}
		}
		return size
	}

	for disconnectCount > 0 {
		// Find the component with the largest group
		maxSize := 0
		maxComponent := ""
		for component := range graph {
			size := dfs(component)
			if size > maxSize {
				maxSize = size
				maxComponent = component
			}
		}

		// Disconnect the largest group by removing the component
		delete(graph, maxComponent)

		// Recalculate group sizes
		groupSizes = append(groupSizes, dfs(resetButton))

		disconnectCount--
	}

	// Multiply the sizes of the disconnected groups
	result := 1
	for _, size := range groupSizes {
		result *= size
	}

	return result, nil
}

func main() {
	filename := "test0.txt" // Change this to your actual input file name
	graph, err := readInput(filename)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println(graph)

	disconnectCount := 3
	result, err := disconnectWires(graph, disconnectCount)
	if err != nil {
		fmt.Println("Error disconnecting wires:", err)
		return
	}

	fmt.Println("Result:", result)
}
