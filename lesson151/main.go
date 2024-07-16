package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the input file
	inputFile, err := os.Open("in.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Create the output file
	outputFile, err := os.Create("out.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Create a scanner to read the input file
	scanner := bufio.NewScanner(inputFile)

	// Read input file line by line and write to output file with line numbers
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		outputFile.WriteString(line + "\n")
	}

	// if err := scanner.Err(); err != nil {
	// 	fmt.Println("Error reading input file:", err)
	// 	return
	// }

	fmt.Println("Data successfully written to output.txt")
}
