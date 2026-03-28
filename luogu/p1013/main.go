package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	solve()
}

func solve() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		return
	}

	table := make([][]string, n)
	mapping := make(map[string]int)

	for i := 0; i < n; i++ {
		table[i] = make([]string, n)
		carries := 0

		for j := 0; j < n; j++ {
			fmt.Fscan(in, &table[i][j])
			if len(table[i][j]) >= 2 {
				carries++
			}
		}

		if i > 0 {
			mapping[table[i][0]] = carries
		}
	}

	letters := table[0][1:]
	base := n - 1

	if !isUnique(mapping, base) {
		fmt.Fprintln(out, "ERROR!")
		return
	}

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			targetSum := mapping[table[i][0]] + mapping[table[0][j]]
			cellStr := table[i][j]
			actualVal := 0

			if len(cellStr) == 1 {
				actualVal = mapping[cellStr]
			} else if len(cellStr) == 2 {
				actualVal = mapping[string(cellStr[0])]*base + mapping[string(cellStr[1])]
			} else {
				fmt.Fprintln(out, "ERROR!")
				return
			}

			if targetSum != actualVal {
				fmt.Fprintln(out, "ERROR!")
				return
			}
		}
	}

	for i, char := range letters {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprintf(out, "%s=%d", char, mapping[char])
	}
	fmt.Fprintln(out)
	fmt.Fprintln(out, base)
}

func isUnique(mapping map[string]int, base int) bool {
	used := make(map[int]bool)
	for _, val := range mapping {
		if val < 0 || val >= base || used[val] {
			return false
		}
		used[val] = true
	}
	// Bijection
	return len(used) == base
}
