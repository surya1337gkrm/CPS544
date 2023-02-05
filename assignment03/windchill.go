package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)

// twc = 35.74 + 0.6215ta - 35.75v^0.16 + 0.4275tav^0.16

func windChill(t, v float64, table bool) string {
	w := 35.74 + (0.6215 * t) - (35.75 * math.Pow(v, 0.16)) + (0.4275 * t * math.Pow(v, 0.16))
	if table {
		res := fmt.Sprintf("%.1f", w)
		return res
	} else {
		res := fmt.Sprintf("%.2f", w)
		return res
	}
}

func confirmWind() float64 {
	var value string
	//var num float64
	for {
		fmt.Print("Enter the wind speed in miles per hour: ")
		fmt.Scan(&value)

		num, err := strconv.ParseFloat(value, 64)
		if err == nil {
			if num >= 2 {
				return num
			} else {
				fmt.Printf("%.6f is out of range.  Valid wind speeds are above 2 mph.\n", num)
			}
		} else {
			fmt.Println(err)
		}
	}
}

func confirmTemp() {
	var value string
	//var num float64
	for {
		fmt.Print("Enter the temperature in °F: ")
		fmt.Scan(&value)

		temp, err := strconv.ParseFloat(value, 64)
		if err == nil {
			if temp >= -58 && temp <= 41 {
				wind := confirmWind()
				res := windChill(temp, wind, false)
				fmt.Printf("The wind chill index is %s°F.\n", res)
				break
			} else {
				fmt.Printf("%.6f is out of range. Valid temperatures are between -58°F and 41°F.\n", temp)
			}
		} else {
			fmt.Println(err)
		}
	}
}

func main() {
	flagPtr := flag.Bool("table", false, "Table Flag")
	flag.Parse()
	if *flagPtr {
		fmt.Println("\nWind Chill Chart.")
		for j := 5.0; j <= 60.0; j = j + 5 {
			line := ""
			for i := 40.0; i >= -45.0; i = i - 5 {
				res := windChill(i, j, true)
				if len(res) <= 4 {
					if len(res) == 3 {
						line = line + "    " + res
					} else {
						line = line + "   " + res
					}
				} else {
					line = line + "  " + res
				}
			}
			fmt.Println(line)
		}
	} else {
		confirmTemp()
	}
}
