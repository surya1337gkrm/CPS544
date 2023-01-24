package main

//Code to display cards based on the command line arguments
//Author: Surya Venkatesh Vijjana
//Course: CPS544
import (
    "fmt"
    "math/rand"
    "os"
    "strings"
    "time"
)

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
|              |
|       /|     |
|      / |     |
|     /  |     |
|    /   |     |
|   /____|__   |
|        |     |
|        |     |
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
    if len(os.Args) <=2{
        fmt.Println("Too Few Arguments.")
    }else{
        alignment:=os.Args[1]
        if alignment=="v"{
            for i:=2;i<len(os.Args);i++{
                if os.Args[i]=="r"{
                    card:=keys[rand.Intn(len(keys))]
                    fmt.Println(stringMap[card])
                }else{
                    _,keyFound:=stringMap[os.Args[i]]
                    if keyFound{
                        fmt.Println(stringMap[os.Args[i]])
                    }else{
                        fmt.Println("Invalid Commandline Argument. Choose from 2-10/a,k,q,j or r.\nTerminating the Program.")
                        break
                    }    
                }
            }
        }else if alignment=="h"{
            args:=os.Args[2:]
            // cards:=make([][]string,len(args))
            var cards [][]string
            for _,v:=range args{
                if v=="r"{
                    cards=append(cards,strings.Split(stringMap[keys[rand.Intn(len(keys))]],"\n"))
                }else{
                    _,keyFound:=stringMap[v]
                    if keyFound{
                        cards=append(cards,strings.Split(stringMap[v],"\n"))
                    }else{
                        fmt.Println("Invalid Commandline Argument Found. Choose from 2-10/a,k,q,j or r.\nTerminating the code.")
                        break
                    }
                }
            }
            for j:=range [11]int{}{
                line:=""
                for i:=range cards{
                    line=line+" "+cards[i][j]
                }
                fmt.Println(line)
            }
            
        }else{
            fmt.Println("Invalid Positional Argument. Choose h for Horizantal/v for Vertical.")
        }
    }
}