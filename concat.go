package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

type parameters struct {
	path    string
	before  string
	beforeR string
	after   string
	afterR  string
	name    string
	format  string
	remove  bool
	splitAt int
}

func getAmountOfFiles(lineCount float64, splitAt float64) int {
	// always round up
	return int(math.Ceil(lineCount / splitAt))
}

func exit(method string, msg string, err error) {
	// handle error
	log.Fatalf("exit method={%s}, msg={%s}, err={%s}", method, msg, err.Error())
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		exit("readLines", "Cannot open file: "+path, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		exit("readLines", "Cannot parse data from file", err)
	}
	return lines
}

// writeLines writes the lines to the given file.
func writeLines(fileLines []string, parameters parameters, dir string) {
	totalLineCount := len(fileLines)
	fmt.Printf("Total Lines: %d\n", totalLineCount)

	// get how many files will be created
	amountOfFiles := getAmountOfFiles(float64(totalLineCount), float64(parameters.splitAt))

	iTotalLines := 0
	// it will generate files based on the split at parameter
	// foreach section, it will generate a file
	for i := 1; i <= amountOfFiles; i++ {
		fmt.Printf("Current file: %d\n", i)
		fmt.Printf("Current line: %d\n", iTotalLines)

		filepath := dir + "/" + parameters.name + "_" + strconv.Itoa(i) + "." + parameters.format
		file, err := os.Create(filepath)
		if err != nil {
			exit("writeLines", "Unable to create a file named: "+filepath, err)
		}

		defer file.Close()

		w := bufio.NewWriter(file)
		fmt.Fprintln(w, parameters.before)

		// creates a bool, avoiding duplicate code
		var endOfFile int
		if i == amountOfFiles {
			// if it is the last file (avoiding index out of bounds with "lineCount-1" instead of "(splitAt*i)-1")
			endOfFile = totalLineCount - 1
		} else {
			endOfFile = (parameters.splitAt * i) - 1
		}

		for iTotalLines <= endOfFile {
			repeatStr := parameters.beforeR + fileLines[iTotalLines] + parameters.afterR

			// check if it is the last line and remove is true
			if iTotalLines == endOfFile && parameters.remove {
				// TODO: it would be fancier to remove a character from "w", instead to insert with one less character into it
				// inserts one less character, because of the "remove" flag
				fmt.Fprintln(w, repeatStr[:len(repeatStr)-1])
			} else {
				fmt.Fprintln(w, repeatStr)
			}
			iTotalLines++
		}

		fmt.Fprintln(w, parameters.after)
		w.Flush()
	}
}

func main() {
	path := flag.String("path", "", "path of file that will be processed")
	splitAt := flag.Int("split", 100000, "at which line it will split the resulted files")
	before := flag.String("before", "", "a string that will be concatenated before all of the repetitions")
	beforeR := flag.String("beforeR", "", "a string that will be concatenated before the start of each repetition")
	after := flag.String("after", "", "a string that will be concatenated after all of the repetitions")
	afterR := flag.String("afterR", "", "a string that will be concatenated after the end of each repetition")
	remove := flag.Bool("remove", false, "if true, it removes the last character o the file, before the contents of the \"after\" flag")
	name := flag.String("name", "generatedFile", "resulting file name")
	format := flag.String("format", "sql", "resulting file format")
	flag.Parse()

	parameters := parameters{
		path:    *path,
		before:  *before,
		beforeR: *beforeR,
		after:   *after,
		afterR:  *afterR,
		name:    *name,
		format:  *format,
		remove:  *remove,
		splitAt: *splitAt,
	}

	fileLines := readLines(*path)
	dir := filepath.Dir(*path)
	writeLines(fileLines, parameters, dir)
}
