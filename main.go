package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

/* Global Variables */
var (
	// List of colours to be used for the graph
	colours = []string{
        "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", 
        "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}

	// Adjacency list for the graph
	adjList = make(map[int][]int)
)

func main() {
	// Check if the correct number of arguments have been provided
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		os.Exit(1)
	}

	// Open the input file
	file := openFile(os.Args[1])

	// Parse the input file and create the adjacency list
	parseInput(file)

	// Sort the nodes in the adjacency list by the number of connections
	sortedKeys := sortNodes()

	// Colour the graph
	colourMap := colourGraph(sortedKeys)

	// Output the colours assigned to each node
	outputColours(colourMap, os.Args[2])
}

/* Output the colours assigned to each node accordingly */
func outputColours(colourMap map[int]string, output string) {
	
	file, err := os.Create(output)
	if err != nil {
		fmt.Println("Error creating output file: ", err)
		os.Exit(1)
	}

	// Write the colours to the output file

	for i := range len(colourMap) {
		_, err := file.WriteString(strconv.Itoa(i + 1) + colourMap[i + 1] + "\n")
		if err != nil {
			fmt.Println("Error writing to output file: ", err)
		}
	}

	// Close output file once done
	defer file.Close()
}

/* 
Function to run the graph colouring algorithm:
	1. Assume an ordered list of colours (eg, red, black, blue, etc, here denoted by A, B, C, …)
	2. Assume an interference graph, where nodes are numbered: 1, 2, 3, …
	3. Rank nodes (that is, live ranges) of the interference graph according to the number of 
	   neighbours in descending order. In case of a tie (that is, nodes with the same number of neighbours) the node with the lowest id takes priority.
	4. Follow the ranking to assign colours from the list of colours. For each node, select the first colour from the list that is not used by the node’s neighbours.
	5. Keep following the ranking and repeating step 4 until all nodes are coloured.

Parameters:
	sortedKeys: A slice (list) of the node numbers sorted by the number of connections
Returns: 
	A map of node numbers to the colour assigned to each node
*/
func colourGraph(sortedKeys []int) (map[int]string) {
	// Create a map to hold the colours assigned to each node
	// Node Number --> Colour
	colourMap := make(map[int]string)

	for _, node := range sortedKeys {
		neighbours := adjList[node]
		neighbourColours := make(map[string]bool)

		// Get the colours of the neighbours
		for _, neighbour := range neighbours {
			colour := colourMap[neighbour]
			if colour != "" {
				neighbourColours[colour] = true
			}
		}

		// Assign the first available colour to the node
		for _, colour := range colours {
			if !neighbourColours[colour] {
				colourMap[node] = colour
				break
			}
		}
	}

	return colourMap
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
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")		   		// Get the current line
		if len(line) < 1 {
			fmt.Println("Error: Invalid input file format")		// We should never get here but just in case
			os.Exit(1)
		}
		node, err := strconv.Atoi(line[0])						// Get the node number
		if err != nil {
			fmt.Println("Error parsing first node: ", err)
		}
		
		connections := make([]int, len(line[1:]))				// Create a slice to hold the connections
		for i, conn := range line[1:] {
			connections[i], err = strconv.Atoi(conn)			// Convert the connections to integers
			if err != nil {
				fmt.Println("Error parsing connections: ", err)
			}
		}
		
		adjList[node] = connections
	}

	defer file.Close()
	return adjList
}

/* Open a file, return a pointer to it and exit gracefully if there is an error */
func openFile(path string) (*os.File) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return file
}

/* 
Sort the nodes in the adjacency list by the number of connections
If two nodes have the same number of connections, sort them by node number with the smaller node number first
Returns a slice of the sorted node numbers
*/
func sortNodes() ([]int) {
	sortedKeys := make([]int, 0, len(adjList))
	for k, _ := range adjList {
		sortedKeys = append(sortedKeys, k)
	}

	slices.SortFunc(sortedKeys,
		func(a, b int) int {
			if len(adjList[a]) == len(adjList[b]) {
				return cmp.Compare(a, b)
			}
			return -cmp.Compare(len(adjList[a]), len(adjList[b]))
		})

	return sortedKeys
}