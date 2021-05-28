package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	rowNum := 0
	fileA := "a.csv"
	fileB := "b.csv"

	f, err := os.Open(fileA)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rb := csv.NewReader(f)

	m1 := make(map[string]int)
	var lc1 int
	for {
		r, err := rb.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		lc1++
		m1[r[rowNum]]++
	}

	f2, err := os.Open(fileB)
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	rb2 := csv.NewReader(f2)
	m2 := make(map[string]int)
	var lc2 int
	for {
		r, err := rb2.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		lc2++
		m2[r[rowNum]]++
	}
	var c1 int
	for v := range m1 {
		n, ok := m2[v]
		if !ok {
			c1++
			fmt.Printf("B missing, %s\n", v)
		}
		if n > 1 {
			fmt.Printf("B contains %d entries of %s\n", n, v)
		}

	}

	var c2 int
	for v := range m2 {
		n, ok := m1[v]
		if !ok {
			c2++
			fmt.Printf("A missing, %s\n", v)
		}
		if n > 1 {
			fmt.Printf("A contains %d entries of %s\n", n, v)
		}

	}

	fmt.Printf("\nLines A(%d), B(%d), diff(%d)\n", lc1, lc2, lc1-lc2)
	fmt.Printf("Unique entries A(%d), B(%d), diff(%d)\n", len(m1), len(m2), len(m1)-len(m2))
	fmt.Printf("Missing entries A(%d), B(%d)\n", c2, c1)

}
