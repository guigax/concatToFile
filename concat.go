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

func getAmountOfFiles(lineCount float64, splitAt float64) int {
	return int(math.Ceil(lineCount / splitAt))
}

func exit(method string, msg string, err error) {
	// handle error
	log.Fatalf("%s: %s", method, err)
	fmt.Println(msg)
	os.Exit(1)
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		exit("readLines", "cannot open file: "+path, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		exit("readLines", "cannot parse data from file", err)
	}
	return lines
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, splitAt int, before string, beforeR string, after string, afterR string, remove bool, name string, format string, dir string) {
	lineCount := len(lines)
	fmt.Println("Total Lines: " + strconv.Itoa(lineCount))

	// get how many files will be created
	filesGenerated := getAmountOfFiles(float64(lineCount), float64(splitAt))

	currentLine := 0
	for i := 1; i <= filesGenerated; i++ {
		fmt.Println("Current file: " + strconv.Itoa(i))
		fmt.Println("Current line: " + strconv.Itoa(currentLine))

		file, _ := os.Create(dir + "/" + name + "_" + strconv.Itoa(i) + "." + format)
		defer file.Close()

		w := bufio.NewWriter(file)
		fmt.Fprintln(w, before)

		// if it is the last file (avoiding index out of bounds with "lineCount-1" instead of "(splitAt*i)-1")
		if i == filesGenerated {
			for currentLine <= lineCount-1 {
				repeatStr := beforeR + lines[currentLine] + afterR

				// check if it is the last line and remove is true
				if currentLine == lineCount-1 && remove {
					fmt.Fprintln(w, repeatStr[:len(repeatStr)-1])
				} else {
					fmt.Fprintln(w, repeatStr)
				}
				currentLine++
			}
		} else {
			for currentLine <= (splitAt*i)-1 {
				repeatStr := beforeR + lines[currentLine] + afterR

				// check if it is the last line and remove is true
				if currentLine == (splitAt*i)-1 && remove {
					fmt.Fprintln(w, repeatStr[:len(repeatStr)-1])
				} else {
					fmt.Fprintln(w, repeatStr)
				}
				currentLine++
			}
		}
		fmt.Fprintln(w, after)

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

	lines := readLines(*path)
	dir := filepath.Dir(*path)
	writeLines(lines, *splitAt, *before, *beforeR, *after, *afterR, *remove, *name, *format, dir)
}
