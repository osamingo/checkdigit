package checkdigit_test

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/osamingo/checkdigit"
)

func TestNewLuhn(t *testing.T) {

	p := checkdigit.NewLuhn()
	if p == nil {
		t.Error("failed to generate luhn provider")
	}
}

func TestNewDamm(t *testing.T) {

	p := checkdigit.NewDamm()
	if p == nil {
		t.Error("failed to generate damm provider")
	}
}

func TestNewVerhoeff(t *testing.T) {

	p := checkdigit.NewVerhoeff()
	if p == nil {
		t.Error("failed to generate verhoeff provider")
	}
}

func ExampleNewLuhn() {

	p := checkdigit.NewLuhn()

	seed := "411111111111111"
	cd, err := p.Generate(seed)
	if err != nil {
		log.Fatalln("failed to generate check digit")
	}

	ok := p.Verify(seed + strconv.Itoa(cd))
	fmt.Printf("seed: %s, check digit: %d, verify: %t\n", seed, cd, ok)

	// Output:
	// seed: 411111111111111, check digit: 1, verify: true
}

func ExampleNewDamm() {

	p := checkdigit.NewDamm()

	seed := "572"
	cd, err := p.Generate(seed)
	if err != nil {
		log.Fatalln("failed to generate check digit")
	}

	ok := p.Verify(seed + strconv.Itoa(cd))
	fmt.Printf("seed: %s, check digit: %d, verify: %t\n", seed, cd, ok)

	// Output:
	// seed: 572, check digit: 4, verify: true
}

func ExampleNewVerhoeff() {

	p := checkdigit.NewVerhoeff()

	seed := "236"
	cd, err := p.Generate(seed)
	if err != nil {
		log.Fatalln("failed to generate check digit")
	}

	ok := p.Verify(seed + strconv.Itoa(cd))
	fmt.Printf("seed: %s, check digit: %d, verify: %t\n", seed, cd, ok)

	// Output:
	// seed: 236, check digit: 3, verify: true
}
