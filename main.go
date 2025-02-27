package main

import (
	"fmt"
	"os"
	"bufio"
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
func parseInput(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	defer file.Close()
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