package eqnparse

import (
	// "bytes"
	"strconv"
	"strings"
)

// STUDENTS: MODIFY THE IMPLEMENTATIONS TO IMPROVE THE BENCHMARK

func (expr Expression) String() (s string) {
	var buffer strings.Builder
	buffer.Grow(len(expr.Numbers) + len(expr.Operators))
	var err error
	for i := 0; i < len(expr.Numbers)+len(expr.Operators); i++ {
		if i%2 == 0 {
			// even
			// s += strconv.Itoa(expr.Numbers[i/2])
			_, err = buffer.WriteString(strconv.Itoa(expr.Numbers[i/2]))

		} else {
			// odd
			// s += string(expr.Operators[(i-1)/2])
			_, err = buffer.WriteString(string(expr.Operators[(i-1)/2]))
		}
		if err != nil {
			panic(err)
		}

	}
	s = buffer.String()
	buffer.Reset()
	return s
}

func (eqn Equation) String() string {
	var buffer strings.Builder
	buffer.Grow(len(eqn.lhs.Numbers) + len(eqn.lhs.Operators) + len(eqn.rhs.Numbers) + len(eqn.rhs.Operators) + 1)
	s := ""
	_, err1 := buffer.WriteString(s)
	// return eqn.lhs.String() + "=" + eqn.rhs.String()
	_, err2 := buffer.WriteString(eqn.lhs.String())
	_, err3 := buffer.WriteString("=")
	_, err4 := buffer.WriteString(eqn.rhs.String())
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		panic("error occurred")
	}
	s = buffer.String()
	buffer.Reset()
	return s
}
