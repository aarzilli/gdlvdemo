package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Thing struct {
	X, Y int
}

func aComplexFunction(idx int) (*Thing, *Thing) {
	if idx%2 == 0 {
		return theBaroness(idx)
	} else {
		p := theWealthyCapitalist(idx + 1)
		var q *Thing
		if P(idx + 1) {
			p, q = mrReadHarrings(p, idx)
			return p, q
		} else if P(idx - 1) {
			q = thePainter(p)
			return p, q
		}
		return p, theButler()
	}
}

func main() {
	if len(os.Args) > 2 {
		fmt.Fprintf(os.Stderr, "Wrong number of arguments, specify a number\n")
		os.Exit(1)
	}
	var n int
	if len(os.Args) > 1 {
		n, _ = strconv.Atoi(os.Args[1])
	} else {
		rand.Seed(time.Now().Unix())
		n = rand.Intn(30)
		n += 20
	}
	p, q := aComplexFunction(n)
	execute(p.X, q.Y)
	fmt.Printf("success %d %d %d\n", p.X, q.X, n)
}
