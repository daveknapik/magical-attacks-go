package main

import (
	"flag"
	"fmt"
)

func main() {
	length := flag.Int("length", 3, "Number of words in the magical attack")
	fandom := flag.String("fandom", "all", "If specified, restricts attack to words from that fandom (options: precure, sailorMoon)")

	flag.Parse()
	generateAttack(*length, *fandom)
}

func generateAttack(l int, f string) {
	// accept data and length
	data := getDataByFandom(f)
	fmt.Println(data)
	// choose randomly and return
	// perhaps solve selection by generating 3 random indices and then selecting based on them
}
