package main

import (
	"fmt"
	"sort"
)

func main() {
	var strs = []string{"flower", "flow", "flight"}
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	fmt.Println(strs)
	for j := range strs[1:] {
		print(j)
		println(strs[j])
	}
}
