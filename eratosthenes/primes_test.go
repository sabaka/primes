package eratosthenes

import (
	"testing"
)

func TestPrimesUnder(t *testing.T) {
	primes := PrimesUnder(10)
	expected := []int{2, 3, 5, 7}
	for i := 0; i < len(primes); i++ {
		if primes[i] != expected[i] {
			t.Errorf("Unexpected number: %d", primes[i])
		}
	}
}

func TestPrime(t *testing.T) {
	prime := Prime(1)
	if prime != 2 {
		t.Errorf("First prime should be 2. Actual: %d", prime)
	}
	prime = Prime(2)
	if prime != 3 {
		t.Errorf("First prime should be 3. Actual: %d", prime)
	}
	prime = Prime(10001)
	if prime != 104743 {
		t.Errorf("First prime should be 2. Actual: %d", prime)
	}
}

func TestPrimes(t *testing.T) {
	actual := Primes(10)
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Unexpected number: %d", actual[i])
		}
	}
}
