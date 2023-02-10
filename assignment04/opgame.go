// Opgame Assignemt 04
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var input []string
var puzzle []string
var ops = []string{"+", "-", "*", "/"}

// var sol [][]string
var solution []string

func calc(ele1, ele2 int, op string) any {
	switch op {
	case "+":
		return ele1 + ele2
	case "-":
		return ele1 - ele2
	case "*":
		return ele1 * ele2
	case "/":
		if ele2 != 0 {
			if ele1%ele2 == 0 {
				return ele1 / ele2
			}
		}
	}
	return nil
}

func dp(result int, index int, value int, s string) (bool, string) {
	if index == len(puzzle)-1 {
		if value == result {
			return true, s
		}
	}
	if index <= len(puzzle)-2 {
		var ele1, ele2 int
		if index == 0 {
			ele1, _ = strconv.Atoi(puzzle[0])
			ele2, _ = strconv.Atoi(puzzle[1])
			index = 1
		} else {
			ele1 = value
			ele2, _ = strconv.Atoi(puzzle[index])
		}
		for _, op := range ops {
			res := calc(ele1, ele2, op)
			if res != nil {
				val := res.(int)
				// index = index + 1
				b, s := dp(result, index+1, val, s+op+strconv.Itoa(ele2))
				if b {
					// fmt.Println(s + "=" + strconv.Itoa(result))
					solution = append(solution, s+"="+strconv.Itoa(result))
					// fmt.Println(solution)
					// break
				}
			}
		}
	}
	return false, ""
}

func main() {
	allFlag := flag.Bool("all", false, "all flag")
	flag.Parse()
	// fmt.Println(*allFlag)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ele := scanner.Text()
		input = append(input, ele)
	}
	if err := scanner.Err(); err != nil {
		// fmt.Fprintln(os.Stderr, "reading standard input:", err)
		log.Fatal(err)
		// fmt.Fprintln(os.Stderr, err)
	}

	for _, ele := range input {
		puzzle = strings.Split(ele, " ")
		// var solution []string

		for _, element := range puzzle {
			_, err := strconv.Atoi(element)
			if err != nil {
				log.Fatal(err)
				// fmt.Fprintln(os.Stderr, err)
			}
		}
		solution = []string{}
		res, err := strconv.Atoi(puzzle[len(puzzle)-1])
		if err == nil {
			dp(res, 0, 0, puzzle[0])
		}
		if *allFlag {
			fmt.Println(strings.Join(solution, ","))
		} else {
			if len(solution) == 0 {
				fmt.Println("")
			} else {
				fmt.Println(solution[0])
			}
		}
	}
}
