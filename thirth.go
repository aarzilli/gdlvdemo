package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// This program implements an interpreter for a simple stack based language.
// An input containing a single number will push that number on the stack.
//
// The following arithmetic operators are available:
//     + - * / %
// their behavior is popping two variables from the stack a1, a2 (in this
// order) and then pushing "a2 op a1" on the stack
//
// The program:
//    1 2 +
// Will push 1 on the stack, push 2 on the stack, then pop both values and
// push 3 (1+2 on the stack)
//    4 2 /
// Will push 4 on the stack, push 2 on the stack, then pop both values,
// compute 4/2 and push '2' (the result) on thestack.
//
// The following stack operators are available:
//  DROP deletes one level from the stack
//  n DROPN deletes n levels from the stack
//  PRINT deletes one value from the stack, prints it
//  SWAP swaps levels 1 and 2
//  ROT moves level 3 to level 1
//  n ROLL moves level n to level 1
//  n ROLLD moves level 1 to level n
//  n PICK copies level n to level 1
//  DEPTH counts the current depth of the stack, pushes it on the stack
//  DUP duplicates level 1
//  n DUPN duplicates n levels
//  SHOW prints the contents of the stack

var Stack []int

func main() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		v := strings.Split(strings.TrimSpace(in.Text()), " ")
		execList(v)
	}
}

func execList(v []string) {
	var i int
	var cmd string
	defer func() {
		if ierr := recover(); ierr != nil {
			fmt.Printf("error executing word %s at %d: %v\n", cmd, i, ierr)
		}
	}()
	for _, cmd = range v {
		execOne(cmd)
	}
}

func execOne(cmd string) {
	switch strings.ToUpper(cmd) {
	case "+", "-", "/", "*", "%":
		arithOp(cmd)
	case "DROP": // ok
		// deletes one level from the stack
		pop()
	case "DROPN": // ok
		// deletes n levels from the stack
		dropnOp(pop())
	case "PRINT":
		// deletes one value from the stack, prints it
		fmt.Printf("%d\n", pop())
	case "SWAP": // ok
		// swaps levels 1 and 2
		swapOp()
	case "ROT": // ok
		// moves level 3 to level 1
		rollOp(3)
	case "ROLL":
		// moves level n to level 1
		rollOp(pop())
	case "ROLLD":
		// moves level 1 to level n
		rolldOp(pop())
	case "PICK": // ok
		// copies level n to level 1
		pickOp(pop())
	case "DEPTH": // ok
		// counts the current depth of the stack, pushes it on the stack
		depthOp()
	case "DUP": // ok
		// duplicates level 1
		dupnOp(1)
	case "DUPN": // ok
		// duplicates n levels
		dupnOp(pop())
	case "SHOW": // ok
		// prints the contents of the stack
		for i := range Stack {
			fmt.Printf("%d: %d\n", len(Stack)-i, Stack[i])
		}
	default:
		n, err := strconv.ParseInt(cmd, 10, 64)
		if err != nil {
			panic(fmt.Errorf("unknown command %s", cmd))
		}
		push(int(n))
	}
}

func pop() int {
	if len(Stack) <= 0 {
		panic("not enough stack levels")
	}
	r := Stack[len(Stack)-1]
	Stack = Stack[:len(Stack)-1]
	return r
}

func push(x int) {
	Stack = append(Stack, x)
}

func arithOp(op string) {
	a1 := pop()
	a2 := pop()
	var r int
	switch op {
	case "+":
		r = a2 + a1
	case "-":
		r = a2 - a1
	case "*":
		r = a2 * a1
	case "/":
		r = a2 / a1
	case "%":
		r = a2 % a1
	}
	push(r)
}

func dropnOp(n int) {
	if n > len(Stack) {
		panic("not enough stack levels")
	}
	if n <= 0 {
		panic(fmt.Errorf("%d not a stack level", n))
	}
	Stack = Stack[:len(Stack)-n]
}

func swapOp() {
	a1 := pop()
	a2 := pop()
	push(a1)
	push(a2)
}

func rollOp(n int) {
	if n > len(Stack) {
		panic("not enough stack levels")
	}
	if n <= 0 {
		panic(fmt.Errorf("%d not a stack level", n))
	}
	x := Stack[len(Stack)-n]
	copy(Stack[len(Stack)-n:], Stack[len(Stack)-n+1:])
	Stack[len(Stack)-1] = x
}

func rolldOp(n int) {
	if n > len(Stack) {
		panic("not enough stack levels")
	}
	if n <= 0 {
		panic(fmt.Errorf("%d not a stack level", n))
	}
	x := Stack[len(Stack)-1]
	copy(Stack[len(Stack)-n+1:], Stack[len(Stack)-n:len(Stack)-1])
	Stack[len(Stack)-n] = x
}

func pickOp(n int) {
	if n > len(Stack) {
		panic("not enough stack levels")
	}
	if n <= 0 {
		panic(fmt.Errorf("%d not a stack level", n))
	}
	x := Stack[len(Stack)-n]
	push(x)
}

func depthOp() {
	push(len(Stack))
}

func dupnOp(n int) {
	if n > len(Stack) {
		panic("not enough stack levels")
	}
	if n <= 0 {
		panic(fmt.Errorf("invalid argument %d", n))
	}
	v := make([]int, n)
	copy(v, Stack[len(Stack)-n:])
	Stack = append(Stack, v...)
}
