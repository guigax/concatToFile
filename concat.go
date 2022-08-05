package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

type parameters struct {
	source      string
	destination string
	before      string
	beforeR     string
	after       string
	afterR      string
	name        string
	format      string
	splitAt     int
	remove      bool
}

/*
	always round up
	it uses math.Ceil and it needs two float64 args
*/
func getAmountOfFiles(lineCount, splitAt int) int {
	return int(math.Ceil(float64(lineCount) / float64(splitAt)))
}

func exit(method, msg string, err error) {
	// handle error
	log.Fatalf("exit method={%s}, msg={%s}, err={%s}", method, msg, err.Error())
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func openAndReadFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		exit("openAndReadFile", "cannot open file: "+fileName, err)
	}
	defer file.Close()

	lines, err := readFile(file)
	if err != nil {
		fmt.Printf("cailed to read file: %s", fileName)
	}

	return lines
}

func readFile(reader io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		exit("openAndReadFile", "cannot parse data from file", scanner.Err())
	}

	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(fileLines []string, params parameters) {
	totalLineCount := len(fileLines)
	fmt.Printf("Total Lines: %d\n", totalLineCount)

	// get how many files will be created
	amountOfFiles := getAmountOfFiles(totalLineCount, params.splitAt)

	iTotalLines := 0
	// it will generate files based on the split at parameter
	// foreach section, it will generate a file
	for iFile := 1; iFile <= amountOfFiles; iFile++ {
		fmt.Printf("The file %d starts at %d lines\n", iFile, iTotalLines)

		filepath := params.destination + "/" + params.name + "_" + strconv.Itoa(iFile) + "." + params.format
		file, err := os.Create(filepath)
		if err != nil {
			exit("writeLines", "unable to create a file named: "+filepath, err)
		}

		defer file.Close()

		w := bufio.NewWriter(file)
		fmt.Fprintln(w, params.before)

		// creates a bool, avoiding duplicate code
		var endOfFile int
		if iFile == amountOfFiles {
			// if it is the last file (avoiding index out of bounds with "lineCount-1" instead of "(splitAt*i)-1")
			endOfFile = totalLineCount - 1
		} else {
			endOfFile = (params.splitAt * iFile) - 1
		}

		for iTotalLines <= endOfFile {
			repeatStr := params.beforeR + fileLines[iTotalLines] + params.afterR

			// check if it is the last line and "remove" is true, than inserts one less character to the generated file
			if iTotalLines == endOfFile && params.remove {
				// TODO: it would be fancier to remove a character from "w", instead to insert with one less character into it
				fmt.Fprintln(w, repeatStr[:len(repeatStr)-1])
			} else {
				fmt.Fprintln(w, repeatStr)
			}
			iTotalLines++
		}

		fmt.Fprintln(w, params.after)
		w.Flush()
	}
}

func main() {
	parameters := parseFlags()
	fileContent := openAndReadFile(parameters.source)
	writeLines(fileContent, parameters)
}
