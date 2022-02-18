package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	length := flag.Int("length", 3, "Number of words in the magical attack")
	fandom := flag.String("fandom", "all", "If specified, restricts attack to words from that fandom (options: precure, sailorMoon)")

	flag.Parse()
	// join the attack slice into a space-separated string
	fmt.Println(strings.Join(generateAttack(*length, *fandom), " "))
}

func generateAttack(l int, f string) []string {
	// accept data and length
	data := getDataByFandom(f)

	// choose randomly
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// initialize slice of strings to hold randomly selected words
	attack := []string{}

	// iterate over randomized slice of indices and append to attack slice
	for _, i := range r.Perm(len(data))[:l] {
		attack = append(attack, data[i])
	}

	return attack
}
