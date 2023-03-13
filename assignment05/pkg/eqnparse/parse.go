package eqnparse

import (
	"fmt"
	"strconv"
	"strings"
)

// IMPLEMENT THIS

func fn(arr []string) (*Expression, error) {
	var exp Expression
	for idx := 0; idx < len(arr); idx++ {
		n := ""
		ele := arr[idx]
		if ele == " " {
			continue
		} else if _, err := strconv.Atoi(ele); err == nil {
			n += ele
			if idx < len(arr)-1 {
				nextEle := arr[idx+1]
				if _, err := strconv.Atoi(nextEle); err == nil {
					n += nextEle
					idx++
				}
			}

			val, _ := strconv.Atoi(n)
			exp.Numbers = append(exp.Numbers, val)

		} else {
			var op Operator
			switch ele {
			case "+":
				op = Addition
			case "-":
				op = Subtraction
			case "*":
				op = Multiplication
			case "/":
				op = Division
			default:
				return nil, fmt.Errorf("unexpected Char %s", ele)
			}

			exp.Operators = append(exp.Operators, op)

		}
	}
	if len(exp.Operators) != len(exp.Numbers)-1 {
		return nil, fmt.Errorf("invalid expression")
	}
	return &exp, nil
}

// ParseEquation parses the given equation and return it otherwise it returns an error
func ParseEquation(equation string) (*Equation, error) {
	// STUDENTS: IMPLEMENT THIS FUNCTION
	if !strings.Contains(equation, "=") {
		return nil, fmt.Errorf("invalid equation. doesn't contain =")
	}
	splitString := strings.Split(equation, "=")
	left := strings.Split(splitString[0], "")
	right := strings.Split(splitString[1], "")
	exp1, err1 := fn(left)
	if err1 != nil {
		return nil, fmt.Errorf("error: %w", err1)
	}
	exp2, err2 := fn(right)
	if err2 != nil {
		return nil, fmt.Errorf("error: %w", err2)
	}

	var eq Equation
	eq.lhs = *exp1
	eq.rhs = *exp2
	return &eq, nil

}
