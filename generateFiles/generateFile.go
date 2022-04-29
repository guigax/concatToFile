package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	file, err := os.Create(path + "/text_created.txt")
	if err != nil {
		log.Fatalf("writeLines: %s", err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for i := 1; i <= 753567; i++ {
		fmt.Fprintln(w, "line "+strconv.Itoa(i))
	}
	w.Flush()
}
