//Code to display cards based on the command line arguments
//Author: Surya Venkatesh Vijjana
//Course: CPS544

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var card1 string
var card2 string

// Function to assign values to the card1 and card2
func cards(arr [13]string, s map[string]string) {
	if os.Args[2] == "r" {
		card1 = arr[rand.Intn(len(arr))]
	} else {
		_, keyFound := s[os.Args[2]]
		if keyFound {
			card1 = os.Args[2]
		} else {
			fmt.Println("Invalid Argument for Card1.Choose from 2-10/a,k,q,j or r.")
			return
		}
	}
	if os.Args[3] == "r" {
		card2 = arr[rand.Intn(len(arr))]
	} else {
		_, keyFound := s[os.Args[3]]
		if keyFound {
			card2 = os.Args[3]
		} else {
			fmt.Println("Invalid Argument for card2.Choose from 2-10/a,k,q,j or r.")
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	a := `----------------
|              |
|      /\      |
|     /  \     |
|    /    \    |
|   /------\   |
|  /        \  |
| /          \ |
|              |
----------------
`
	j := `----------------
|              |
|  ----------  |
|       |      |
|       |      |
|       |      |
|       |      |
|  _ _ _|      |
|              |
----------------
`
	q := `----------------
|              |
|  ----------  |
|  |        |  |
|  |        |  |
|  |   \    |  |
|  -----\----  |
|        \     |
|         \    |
----------------
`

	k := `----------------
|              |
|     |  /     |
|     | /      |
|     |/       |
|     |\       |
|     | \      |
|     |  \     |
|              |
----------------
`
	char2 := `----------------
|    ________  |
|           |  |
|           |  |
|   ________|  |
|   |          |
|   |          |
|   |_______   |
|              |
----------------
`
	char3 := `----------------
|   ________   |
|          |   |
|          |   |
|   _______|   |
|          |   |
|          |   |
|   _______|   |
|              |
----------------
`
	char4 := `----------------
|       /|     |
|      / |     |
|     /  |     |
|    /   |     |
|   /____|__   |
|        |     |
|        |     |
|              |
----------------
`
	char5 := `----------------
|   ________   |
|   |          |
|   |          |
|   |_______   |
|          |   |
|          |   |
|   _______|   |
|              |
----------------
`
	char6 := `----------------
|   ________   |
|   |          |
|   |          |
|   |_______   |
|   |       |  |
|   |       |  |
|   |_______|  |
|              |
----------------
`
	char7 := `----------------
|   ________   |
|          /   |
|         /    |
|        /     |
|       /      |
|      /       |
|     /        |
|              |
----------------
`
	char8 := `----------------
|   ________   |
|   |       |  |
|   |       |  |
|   |_______|  |
|   |       |  |
|   |       |  |
|   |_______|  |
|              |
----------------
`
	char9 := `----------------
|   ________   |
|  |       |   |
|  |       |   |
|  |_______|   |
|          |   |
|          |   |
|   _______|   |
|              |
----------------
`
	char10 := `----------------
|              |
|  |  -------  |
|  |  |     |  |
|  |  |     |  |
|  |  |     |  |
|  |  |     |  |
|  |  |_____|  |
|              |
----------------
`

	stringMap := map[string]string{
		"2":  char2,
		"3":  char3,
		"4":  char4,
		"5":  char5,
		"6":  char6,
		"7":  char7,
		"8":  char8,
		"9":  char9,
		"10": char10,
		"a":  a,
		"j":  j,
		"k":  k,
		"q":  q,
	}
	keys := [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "a", "j", "q", "k"}
	if len(os.Args) >= 4 {
		alignment := os.Args[1]
		if alignment == "v" {
			cards(keys, stringMap)
			fmt.Println(stringMap[card1])
			fmt.Println(stringMap[card2])
		} else if alignment == "h" {
			cards(keys, stringMap)
			card1Arr := strings.Split(stringMap[card1], "\n")
			card2Arr := strings.Split(stringMap[card2], "\n")
			for i := range [11]int{} {
				fmt.Println(card1Arr[i] + "  " + card2Arr[i])
			}
		} else {
			fmt.Println("Invalid Positional Argument. Enter v for vertical/h for horizantal.")
		}
	} else {
		fmt.Println("Too Few Commandline Arguments.")
	}

}
