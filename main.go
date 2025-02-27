package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {

}

/* 
Take the input file and return an adjacency list (map) of the graph
For a graph with n nodes, the input file has n lines.
First number in each line is the node number, followed by the nodes it is connected to
	Example:
	1 2 3 4
	2 4 1
	3 1
	4 1 2
Returns:
	1: [2, 3, 4]
	2: [4, 1]
	3: [1]
	4: [1, 2]
*/
func parseInput(file *os.File) (map[int][]int) {
	// The map that will hold the adjacency list
	adjList := make(map[int][]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")		   // Get the current line
		node, err := strconv.Atoi(line[0])									   // Get the node number
		if err != nil {
			fmt.Println("Error parsing first node: ", err)
		}
		
		connections := make([]int, len(line[1:]))						   // Create a slice to hold the connections
		for i, conn := range line[1:] {
			connections[i], err = strconv.Atoi(conn)					   // Convert the connections to integers
			if err != nil {
				fmt.Println("Error parsing connections: ", err)
			}
		}
		
		adjList[node] = connections
	}

	defer file.Close()
	return adjList
}

// Open a file, return a pointer to it and exit gracefully if there is an error
func openFile(path string) (*os.File) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return file
}