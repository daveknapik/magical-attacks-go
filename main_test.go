package main

import (
	"testing"
)

func TestGenerateAttack(t *testing.T) {
	attack := generateAttack(3, "all")
	if len(attack) != 3 {
		t.Errorf("Expected 3 words in the attack but got %v", len(attack))
	}

	attack = generateAttack(5, "all")
	if len(attack) != 5 {
		t.Errorf("Expected 5 words in the attack but got %v", len(attack))
	}

	attack = generateAttack(6, "precure")
	for _, v := range attack {
		if !(contains(getDataByFandom("precure"), v)) {
			t.Errorf("Expected all items to be in the precure dataset but %s came from outside of it", v)
		}
	}
}
