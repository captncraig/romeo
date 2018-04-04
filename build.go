package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("digraph {")
	dat, _ := ioutil.ReadFile("input.txt")
	for i, line := range strings.Split(string(dat), "\n") {
		fmt.Printf("p%d [label=\"%d\"]\n", i+1, i+1)
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			segs := strings.Split(line, " ")
			for i, seg := range segs {
				segs[i] = "p" + seg
			}
			fmt.Printf("p%d [fillcolor=cyan style=filled]\n", i+1)
			fmt.Printf("p%d -> {%s}\n", i+1, strings.Join(segs, " "))
		} else {
			fmt.Printf("p%d [fillcolor=red style=filled]\n", i+1)
		}
	}
	fmt.Println("}")
}
