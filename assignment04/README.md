# Assignment: OpGame

![Math operators](https://upload.wikimedia.org/wikipedia/commons/thumb/1/12/Basic_arithmetic_operators.svg/330px-Basic_arithmetic_operators.svg.png)

## Overview

In this assignment you will write a Go program to solve a simple mathematical puzzle.  The program must efficiently find the binary operators that complete the equation.  For this problem the only operators you need to consider are addition, subtraction, multiplication, and division (when the divisor is non-zero and the remainder is zero).  For example, given `1 O 2 = 3` the only operator that works to put in the circle is "+".  So `1 + 2 = 3` solves the problem.

For simplicity we will not use the normal operator precedence.  Instead there will be no precedence thus expressions are evaluated strictly from left to right.  For example, given `1 O 2 O 3 = 1`,  your goal is to find two operators that satisfy the equation.  The solution would be `1 + 2 / 3 = 1`.  Notice that this is computed as `(1+2) / 3 = 1` and not `1 + (2/3)`.

Generating your own example problems is easy.  Write an equation (being careful to evaluate it strictly left to right) and then remove the operators and supply that to your program.

## Learning Objectives

- Writing functions
- Tree searching
- Recursion

## Requirements

Write a Go program to do the following:

- Reads problems from STDIN as a list of positive integers separated by whitespace, one problem per line.
- Output a solution to the problem to STDOUT that makes the equation valid.  Each +, -, *, /, and = must be passed by one space on each side.
- If no solution is found simply output a blank line to STDOUT.
- Errors in parsig STDIN exit with a non-zero status code after writing the error to STDERR.
- **Your program must use recursion to search for solutions.**  Specifically your program should break the problem down into a set of smaller problems to solve.  This is known as [Dynamic Programming](https://en.wikipedia.org/wiki/Dynamic_programming).
- Do not use any library other than the Go standard library.
- The source code must compile with the most recent version of the Go compiler.
- The program must not panic under any circumstances.
- Make sure your code is "gofmt'd".  See "gofmt" or better use "goimports" or better yet configure IDE to do this formatting on file save.
- Commit and push your working code to your GIT repository.

## Additional Requirements (CPS 544 only)

CPS 444 are welcome to do this part of the assignment as well.

- If you program is provided with the flag `-all` then find all solutions to the problem.
- When outputting multiple solutions, separate them by a comma followed by a space.

## Example Output

```shell
$ echo "3 1 2" | go run .
3-1=2

$ echo "5 4 2 22" | go run .
5*4+2=22

$ (echo "3 1 2" && echo "9 2 18") | go run .
3-1=2
9*2=18

$ echo "6 2 3 4" | go run .
6*2/3=4

$ go run . < testdata/basic.txt
3 - 1 = 2

9 + 0 = 9
2 * 3 = 6
4 + 5 = 9
4 + 5 + 6 = 15
2 + 3 + 1 = 6
```

With multiple solutions it looks like:

```shell
$ echo "7 3 3 7" | go run . -all
7+3-3=7, 7-3+3=7, 7*3/3=7

$ go run . -all < testdata/basic.txt
3 - 1 = 2

9 + 0 = 9, 9 - 0 = 9
2 * 3 = 6
4 + 5 = 9
4 + 5 + 6 = 15
2 + 3 + 1 = 6, 2 * 3 * 1 = 6, 2 * 3 / 1 = 6
```

## Hints

- Commit and push often so you do not loose any of your work.
- Use the `Makefile`.  For example, you can run tests with `make test`.  You can make sure you code is properly formatted with `make fix`.
