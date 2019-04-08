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

func TestNewISBN10(t *testing.T) {

	v := checkdigit.NewISBN10()
	if v == nil {
		t.Error("failed to generate isbn10 verifier")
	}
}

func TestNewISBN13(t *testing.T) {

	p := checkdigit.NewISBN13()
	if p == nil {
		t.Error("failed to generate isbn13 provider")
	}
}

func TestNewEAN8(t *testing.T) {

	p := checkdigit.NewEAN8()
	if p == nil {
		t.Error("failed to generate ean8 provider")
	}
}

func TestNewEAN13(t *testing.T) {

	p := checkdigit.NewEAN13()
	if p == nil {
		t.Error("failed to generate ean13 provider")
	}
}

func TestNewJAN8(t *testing.T) {

	p := checkdigit.NewJAN8()
	if p == nil {
		t.Error("failed to generate jan8 provider")
	}
}

func TestNewJAN13(t *testing.T) {

	p := checkdigit.NewJAN13()
	if p == nil {
		t.Error("failed to generate jan13 provider")
	}
}

func TestNewITF(t *testing.T) {

	p := checkdigit.NewITF()
	if p == nil {
		t.Error("failed to generate itf provider")
	}
}

func TestNewUPC(t *testing.T) {

	p := checkdigit.NewUPC()
	if p == nil {
		t.Error("failed to generate upc provider")
	}
}

func TestNewSSCC(t *testing.T) {

	p := checkdigit.NewSSCC()
	if p == nil {
		t.Error("failed to generate sscc provider")
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

func ExampleNewISBN10() {

	v := checkdigit.NewISBN10()

	code := "155860832X"
	ok := v.Verify(code)
	fmt.Printf("code: %s, verify: %t\n", code, ok)

	// Output:
	// code: 155860832X, verify: true
}

func ExampleNewISBN13() {

	p := checkdigit.NewISBN13()

	seed := "978000271217"
	cd, err := p.Generate(seed)
	if err != nil {
		log.Fatalln("failed to generate check digit")
	}

	ok := p.Verify(seed + strconv.Itoa(cd))
	fmt.Printf("seed: %s, check digit: %d, verify: %t\n", seed, cd, ok)

	// Output:
	// seed: 978000271217, check digit: 0, verify: true
}

func ExampleNewEAN8() {

	p := checkdigit.NewEAN8()

	seed := "9638507"
	cd, err := p.Generate(seed)
	if err != nil {
		log.Fatalln("failed to generate check digit")
	}

	ok := p.Verify(seed + strconv.Itoa(cd))
	fmt.Printf("seed: %s, check digit: %d, verify: %t\n", seed, cd, ok)

	// Output:
	// seed: 9638507, check digit: 4, verify: true
}

func ExampleNewEAN13() {

	p := checkdigit.NewEAN13()

	seed := "590123412345"
	cd, err := p.Generate(seed)
	if err != nil {
		log.Fatalln("failed to generate check digit")
	}

	ok := p.Verify(seed + strconv.Itoa(cd))
	fmt.Printf("seed: %s, check digit: %d, verify: %t\n", seed, cd, ok)

	// Output:
	// seed: 590123412345, check digit: 7, verify: true
}

func ExampleNewJAN8() {

	p := checkdigit.NewJAN8()

	seed := "4996871"
	cd, err := p.Generate(seed)
	if err != nil {
		log.Fatalln("failed to generate check digit")
	}

	ok := p.Verify(seed + strconv.Itoa(cd))
	fmt.Printf("seed: %s, check digit: %d, verify: %t\n", seed, cd, ok)

	// Output:
	// seed: 4996871, check digit: 2, verify: true
}

func ExampleNewJAN13() {

	p := checkdigit.NewJAN13()

	seed := "456995111617"
	cd, err := p.Generate(seed)
	if err != nil {
		log.Fatalln("failed to generate check digit")
	}

	ok := p.Verify(seed + strconv.Itoa(cd))
	fmt.Printf("seed: %s, check digit: %d, verify: %t\n", seed, cd, ok)

	// Output:
	// seed: 456995111617, check digit: 9, verify: true
}

func ExampleNewITF() {

	p := checkdigit.NewITF()

	seed := "1456995111617"
	cd, err := p.Generate(seed)
	if err != nil {
		log.Fatalln("failed to generate check digit")
	}

	ok := p.Verify(seed + strconv.Itoa(cd))
	fmt.Printf("seed: %s, check digit: %d, verify: %t\n", seed, cd, ok)

	// Output:
	// seed: 1456995111617, check digit: 6, verify: true
}

func ExampleNewUPC() {

	p := checkdigit.NewUPC()

	seed := "01234567890"
	cd, err := p.Generate(seed)
	if err != nil {
		log.Fatalln("failed to generate check digit")
	}

	ok := p.Verify(seed + strconv.Itoa(cd))
	fmt.Printf("seed: %s, check digit: %d, verify: %t\n", seed, cd, ok)

	// Output:
	// seed: 01234567890, check digit: 5, verify: true
}

func ExampleNewSSCC() {

	p := checkdigit.NewSSCC()

	seed := "04569951110000001"
	cd, err := p.Generate(seed)
	if err != nil {
		log.Fatalln("failed to generate check digit")
	}

	ok := p.Verify(seed + strconv.Itoa(cd))
	fmt.Printf("seed: %s, check digit: %d, verify: %t\n", seed, cd, ok)

	// Output:
	// seed: 04569951110000001, check digit: 6, verify: true
}
