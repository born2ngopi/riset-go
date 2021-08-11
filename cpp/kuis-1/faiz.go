package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func readLines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	defer f.Close()

	if err != nil {
		return nil, errors.New("File not found!")
	}

	lines := []string{}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func makeMapSki(lines []string, loop uint64) []string {
	var newLines []string
	for _, line := range lines {
		i := uint64(0)
		var joinLine string
		for i < loop {
			joinLine += line
			i++
		}

		newLines = append(newLines, joinLine)
	}

	return newLines
}

type result struct {
	ID int
}

func (r *result) AddStep() {
	r.ID++
}

func (r *result) Val() int {
	return r.ID
}

func main() {
	lines, err := readLines("mapSki.csv")
	if err != nil {
		panic(err)
	}

	newLines := makeMapSki(lines, 1)

	var resData result

	j := 0
	for i := 0; i < len(newLines); i++ {
		for j < len(newLines[i]) {
			newChar := fmt.Sprintf("%c", newLines[i][j])
			if newChar == "#" {
				resData.AddStep()
				log.Printf("Baris %d, Kolom %d", i+1, j+1)
				log.Printf("Pohon yang sedang Joni lewati : %d \n", resData.Val())

				j += 3
				break
			} else {
				j += 3
				break
			}
		}
	}

	fmt.Println()
	log.Printf("Total pohon yang sudah Joni lewati : %d", resData.ID)
}
