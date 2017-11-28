package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	pathToCsv string
)

func main() {
	flag.StringVar(&pathToCsv, "csv", "problems.csv", "Path to csv file with questions and answers")
	flag.Parse()

	file, err := os.Open(pathToCsv)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)

	answersCorrect := 0

	for i, record := range records {
		question, answer := record[0], record[1]

		fmt.Fprintf(os.Stdout, "Question #%d: %s = ", i+1, question)
		// Scan os.Stdin Reader for tokens separated by delim "\n"
		scanner.Scan()
		input := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
			return
		}

		if answer == strings.TrimSpace(input) {
			answersCorrect++
		}
	}

	percentage := 100 * (float32(answersCorrect) / float32(len(records)))
	fmt.Printf("You answered %d/%d (%.2f%%) questions correctly!\n", answersCorrect, len(records), percentage)
}
