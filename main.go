package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	step := 0
	path := 1
	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//		var b strings.Builder
		maps := scanner.Text()
		for len(maps) < path {
			maps += maps
		}

		patern := maps[step:path]
		if patern == "." {
			patern = replaceAtIndex(maps, 'O', step)
		} else {
			patern = replaceAtIndex(maps, 'X', step)
			result += 1
		}

		step += 3
		path = step + 1
		fmt.Println(patern)
	}
	fmt.Printf("Total lewatin pepuhunan : %d kali", result)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
