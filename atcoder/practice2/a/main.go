package main

import (
	"bufio"
	"fmt"
	"os"
)

// Fast IO
var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	p := make([]int, n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
		s[i] = 1
	}
	return &DSU{parent: p, size: s}
}

func (d *DSU) Find(i int) int {
	if d.parent[i] == i {
		return i
	}
	d.parent[i] = d.Find(d.parent[i])
	return d.parent[i]
}

func (d *DSU) Unite(i, j int) {
	rootI, rootJ := d.Find(i), d.Find(j)
	if rootI != rootJ {
		if d.size[rootI] < d.size[rootJ] {
			rootI, rootJ = rootJ, rootI
		}
		d.parent[rootJ] = rootI
		d.size[rootI] += d.size[rootJ]
	}
}

func (d *DSU) Same(i, j int) bool {
	return d.Find(i) == d.Find(j)
}

func main() {
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &q)

	dsu := NewDSU(n)

	for i := 0; i < q; i++ {
		var t, u, v int
		fmt.Fscan(in, &t, &u, &v)

		if t == 0 {
			dsu.Unite(u, v)
		} else {
			a, b := dsu.Find(u), dsu.Find(v)
			if a == b {
				fmt.Fprintln(out, 1)
			} else {
				fmt.Fprintln(out, 0)
			}
		}
	}
}
