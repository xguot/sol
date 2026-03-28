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

	var s string
	if _, err := fmt.Fscan(in, &s); err != nil {
		return
	}

	if len(s) == 0 {
		fmt.Fprintln(out, 0)
		return
	}

	// Single char
	ans := 0
	runLen := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			runLen++
		} else {
			ans += runLen * (runLen + 1) / 2
			runLen = 1
		}
	}

	// Tail
	ans += runLen * (runLen + 1) / 2

	ans += count2(s, 'a', 'b') + count2(s, 'b', 'c') + count2(s, 'a', 'c') + count3(s, 'a', 'b', 'c')

	fmt.Fprintln(out, ans)
}

// If the difference (count1 - count2) is identical at two different indices,
// the number of both characters added between those indices must be equal.
func count2(s string, char1, char2 byte) int {
	m2 := make(map[int]int)
	m2[0] = 1

	c1, c2, total := 0, 0, 0

	for i := range s {
		if s[i] == char1 {
			c1++
		} else if s[i] == char2 {
			c2++
		} else {
			m2 = make(map[int]int)
			m2[0] = 1
			c1 = 0
			c2 = 0
			continue
		}

		diff := c1 - c2
		total += m2[diff]
		m2[diff]++
	}

	return total
}

func count3(s string, char1, char2, char3 byte) int {
	m3 := make(map[[2]int]int)
	m3[[2]int{0, 0}] = 1

	c1, c2, c3, total := 0, 0, 0, 0

	for i := range s {
		if s[i] == char1 {
			c1++
		} else if s[i] == char2 {
			c2++
		} else if s[i] == char3 {
			c3++
		}

		key := [2]int{c1 - c2, c2 - c3}
		total += m3[key]
		m3[key]++
	}

	return total
}
