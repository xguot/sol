package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	defer out.Flush()

	var n, x, y int

	fmt.Fscan(in, &n)
	m := make([][4]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &m[i][0], &m[i][1], &m[i][2], &m[i][3])
	}

	fmt.Fscan(in, &x, &y)

	for i := n - 1; i >= 0; i-- {
		a, b, g, k := m[i][0], m[i][1], m[i][2], m[i][3]
		if x >= a && x <= a+g && y >= b && y <= b+k {
			fmt.Fprintln(out, i+1)
			return
		}
	}

	fmt.Fprint(out, -1)
}
