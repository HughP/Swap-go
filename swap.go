package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"flag"
	"log"
)

func main() {
	var translator string
	    flag.StringVar(&translator, "translator", "", "Sereilazed translation of strings")

	var corpus string
	    flag.StringVar(&corpus, "corpus", "", "The corpus we are translating")

	var delim string
	    flag.StringVar(&delim, "delim", ":", "The delimiter we are using in the translator")

			flag.Parse()

	// Set up the mapping of pairs
  conffile,err  := os.Open(translator)
	//Translator file not found
	if err !=nil{
		log.Fatal("Translator file not found")
	}
	conf := bufio.NewReader(conffile)
	pairs := make(map[string]string)
	scanner := bufio.NewScanner(conf)
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), delim)
		pairs[pair[0]] = pair[1]
	}

	// Use the map to make the swaps
inputfile,err  := os.Open(corpus)
if err !=nil{
	log.Fatal("Corpus file not found")
}
	file := bufio.NewReader(inputfile)
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		for _, r1 := range scanner.Text() {
			r2, prs := pairs[string(r1)]
			if prs {
				fmt.Print(string(r2))
				continue
			}
			fmt.Print(string(r1))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
