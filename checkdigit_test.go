package checkdigit_test

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/osamingo/checkdigit"
)

func TestNewLuhn(t *testing.T) {
	t.Parallel()

	p := checkdigit.NewLuhn()
	if p == nil {
		t.Error("failed to generate luhn provider")
	}
}

func TestNewDamm(t *testing.T) {
	t.Parallel()

	p := checkdigit.NewDamm()
	if p == nil {
		t.Error("failed to generate damm provider")
	}
}

func TestNewVerhoeff(t *testing.T) {
	t.Parallel()

	p := checkdigit.NewVerhoeff()
	if p == nil {
		t.Error("failed to generate verhoeff provider")
	}
}

func TestNewISBN10(t *testing.T) {
	t.Parallel()

	v := checkdigit.NewISBN10()
	if v == nil {
		t.Error("failed to generate isbn10 verifier")
	}
}

func TestNewISBN13(t *testing.T) {
	t.Parallel()

	p := checkdigit.NewISBN13()
	if p == nil {
		t.Error("failed to generate isbn13 provider")
	}
}

func TestNewEAN8(t *testing.T) {
	t.Parallel()

	p := checkdigit.NewEAN8()
	if p == nil {
		t.Error("failed to generate ean8 provider")
	}
}

func TestNewEAN13(t *testing.T) {
	t.Parallel()

	p := checkdigit.NewEAN13()
	if p == nil {
		t.Error("failed to generate ean13 provider")
	}
}

func TestNewJAN8(t *testing.T) {
	t.Parallel()

	p := checkdigit.NewJAN8()
	if p == nil {
		t.Error("failed to generate jan8 provider")
	}
}

func TestNewJAN13(t *testing.T) {
	t.Parallel()

	p := checkdigit.NewJAN13()
	if p == nil {
		t.Error("failed to generate jan13 provider")
	}
}

func TestNewITF(t *testing.T) {
	t.Parallel()

	p := checkdigit.NewITF()
	if p == nil {
		t.Error("failed to generate itf provider")
	}
}

func TestNewUPC(t *testing.T) {
	t.Parallel()

	p := checkdigit.NewUPC()
	if p == nil {
		t.Error("failed to generate upc provider")
	}
}

func TestNewSSCC(t *testing.T) {
	t.Parallel()

	p := checkdigit.NewSSCC()
	if p == nil {
		t.Error("failed to generate sscc provider")
	}
}

func ExampleNewLuhn() {
	p := checkdigit.NewLuhn()

	const seed = "411111111111111"
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

	const seed = "572"
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

	const seed = "236"
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
	p := checkdigit.NewISBN10()

	const seed = "155860832"
	cd, err := p.Generate(seed)
	if err != nil {
		log.Fatalln("failed to generate check digit")
	}

	digit := "X"
	if cd != 10 {
		digit = strconv.Itoa(cd)
	}

	ok := p.Verify(seed + digit)
	fmt.Printf("seed: %s, check digit: %s, verify: %t\n", seed, digit, ok)

	// Output:
	// seed: 155860832, check digit: X, verify: true
}

func ExampleNewISBN13() {
	p := checkdigit.NewISBN13()

	const seed = "978000271217"
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

	const seed = "9638507"
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

	const seed = "590123412345"
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

	const seed = "4996871"
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

	const seed = "456995111617"
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

	const seed = "1456995111617"
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

	const seed = "01234567890"
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

	const seed = "04569951110000001"
	cd, err := p.Generate(seed)
	if err != nil {
		log.Fatalln("failed to generate check digit")
	}

	ok := p.Verify(seed + strconv.Itoa(cd))
	fmt.Printf("seed: %s, check digit: %d, verify: %t\n", seed, cd, ok)

	// Output:
	// seed: 04569951110000001, check digit: 6, verify: true
}
