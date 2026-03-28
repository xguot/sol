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

	var a, b, c int
	var s string
	fmt.Fscan(in, &a, &b, &c, &s)

	fmt.Fprintf(out, "%d %s\n", a+b+c, s)
}
