package eqnparse

import (
	"fmt"
	"testing"
)

// STUDENTS: YOU MUST IMPLEMENT ALL OF THE BELOW TESTS IN THIS FILE
// ExampleParseEquation
// TestParseEquation
// BenchmarkParseEquation
// FuzzParseEquation

/*
	Example equations. Note equations do have to be mathematically valid

3-1=2
5*4+2=22
3-1=2
9*2=18
6*2/3=4
3 - 1 = 2
9 + 0 = 9
2 * 3 = 6
4 + 5 = 9
4 + 5 + 6 = 15
2 + 3 + 1 = 6
7+3-3=7* 1
*/
func ExampleParseEquation() {
	eqn, _ := ParseEquation("3-2 =1")

	fmt.Println(eqn.String())
	// Output: 3-2=1
}

func TestParseEquation(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"3-1=2", "3-1=2"},
		{"5*4+2=22", "5*4+2=22"},
		{"3-1=2", "3-1=2"},
		{"9*2=18", "9*2=18"},
		{"6*2/3=4", "6*2/3=4"},
		{"3 - 1 = 2", "3-1=2"},
		{"9 + 0 = 9", "9+0=9"},
		{"2 * 3 = 6", "2*3=6"},
		{"4 + 5 = 9", "4+5=9"},
		{"4 + 5 + 6 = 15", "4+5+6=15"},
		{"2 + 3 + 1 = 6", "2+3+1=6"},
		{"7+3-3=7* 1", "7+3-3=7*1"},
	}
	for _, tc := range tests {
		t.Run("case: "+tc.input, func(t *testing.T) {
			res, _ := ParseEquation(tc.input)
			if res.String() != tc.want {
				t.Errorf("Expected %s,got %s", tc.want, res.String())
			}
		})
	}
	// eqn, _ := ParseEquation("3-2=1")
	// // a := 1
	// // if a == 1 {
	// // 	fmt.Printf("here %T", eqn)
	// // } else {
	// // 	t.Errorf("not equal")
	// // }
	// if eqn.String() != "3-2=1" {
	// 	t.Errorf("Not Equal")
	// }
}
func BenchmarkParseEquation(b *testing.B) {
	b.StopTimer()
	var s = "7-   2+    5*1=   8+2/1"
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		eqn, _ := ParseEquation(s)
		b.StopTimer()
		expectedString := "7-2+5*1=8+2/1"
		if eqn.String() != expectedString {
			b.Errorf("Expected %s, got %s", expectedString, eqn.String())
		}
		b.SetBytes(int64(len(eqn.String())))
	}
}

func FuzzParseEquation(f *testing.F) {
	f.Add("'3+2=1")
	f.Add("3a-2=1bc+2c")
	f.Add("2.2*1=1+1.2")
	f.Add("Input 2+1")
	f.Add("2+*1")
	f.Add("1%2/3+4")
	f.Add("  3+   2=   1")
	f.Add("  ")
	f.Add("3+1=4")
	f.Add("1*2+3=5")
	f.Add("")

	f.Fuzz(func(t *testing.T, a string) {
		eqn, err := ParseEquation(a)
		if err != nil && eqn != nil {
			fmt.Println("string", a, err)
			t.Errorf("error: %s", err)
		}
	})
}
