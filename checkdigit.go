package checkdigit

import "errors"

type (
	// A Verifier is verify to code by implemented algorithm.
	Verifier interface {
		Verify(code string) bool
	}
	// A Generator is generate a check digit by implemented algorithm.
	Generator interface {
		Generate(seed string) (int, error)
	}
	// A Provider has Verifier and Generator interfaces.
	Provider interface {
		Verifier
		Generator
	}
)

// ErrInvalidArgument is happen when given wrong argument.
var ErrInvalidArgument = errors.New("checkdigit: invalid argument")

func isNotNumber(n rune) bool {
	return n < '0' || '9' < n
}

// NewLuhn returns new Provider that implemented luhn algorithm.
func NewLuhn() Provider {
	return &luhn{}
}

// NewDamm returns new Provider that implemented damm algorithm.
func NewDamm() Provider {
	return &damm{
		matrix: [][]int{
			{0, 3, 1, 7, 5, 9, 8, 6, 4, 2},
			{7, 0, 9, 2, 1, 5, 4, 8, 6, 3},
			{4, 2, 0, 6, 8, 7, 1, 3, 5, 9},
			{1, 7, 5, 0, 9, 8, 3, 4, 2, 6},
			{6, 1, 2, 3, 0, 4, 5, 9, 7, 8},
			{3, 6, 7, 4, 2, 0, 9, 5, 8, 1},
			{5, 8, 6, 9, 7, 2, 0, 1, 3, 4},
			{8, 9, 4, 5, 3, 6, 2, 0, 1, 7},
			{9, 4, 3, 8, 6, 1, 7, 2, 0, 5},
			{2, 5, 8, 1, 4, 3, 6, 7, 9, 0},
		},
	}
}

// NewVerhoeff returns new Provider that implemented verhoeff algorithm.
func NewVerhoeff() Provider {
	return &verhoeff{
		multiplication: [][]int{
			{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			{1, 2, 3, 4, 0, 6, 7, 8, 9, 5},
			{2, 3, 4, 0, 1, 7, 8, 9, 5, 6},
			{3, 4, 0, 1, 2, 8, 9, 5, 6, 7},
			{4, 0, 1, 2, 3, 9, 5, 6, 7, 8},
			{5, 9, 8, 7, 6, 0, 4, 3, 2, 1},
			{6, 5, 9, 8, 7, 1, 0, 4, 3, 2},
			{7, 6, 5, 9, 8, 2, 1, 0, 4, 3},
			{8, 7, 6, 5, 9, 3, 2, 1, 0, 4},
			{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		permutation: [][]int{
			{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			{1, 5, 7, 6, 2, 8, 3, 0, 9, 4},
			{5, 8, 0, 3, 7, 9, 6, 1, 4, 2},
			{8, 9, 1, 6, 0, 4, 3, 5, 2, 7},
			{9, 4, 5, 3, 1, 2, 6, 8, 7, 0},
			{4, 2, 8, 6, 5, 7, 3, 9, 0, 1},
			{2, 7, 9, 3, 8, 0, 6, 4, 1, 5},
			{7, 0, 4, 6, 9, 1, 3, 2, 5, 8},
		},
		inverse: []int{0, 4, 3, 2, 1, 5, 6, 7, 8, 9},
	}
}

// NewISBN10 returns new Verifier that implemented modulus 11 weight 10 to 2 calculator.
func NewISBN10() Verifier {
	return &isbn10{}

}

// NewISBN10 returns new Provider that implemented modulus 10 weight 3 calculator.
func NewISBN13() Provider {
	return &isbn13{}
}

// NewEAN8 returns new Provider that implemented GTIN-8 calculator.
func NewEAN8() Provider {
	return &gtin{
		digit:              8,
		correctionPosition: true,
	}
}

// NewEAN13 returns new Provider that implemented GTIN-13 calculator.
func NewEAN13() Provider {
	return &gtin{
		digit:              13,
		correctionPosition: true,
	}
}

// NewJAN8 returns new Provider that implemented GTIN-8 calculator.
func NewJAN8() Provider {
	return &gtin{
		digit:              8,
		correctionPosition: true,
	}
}

// NewJAN13 returns new Provider that implemented GTIN-13 calculator.
func NewJAN13() Provider {
	return &gtin{
		digit:              13,
		correctionPosition: true,
	}
}

// NewITF returns new Provider that implemented GTIN-14 calculator.
func NewITF() Provider {
	return &gtin{
		digit: 14,
	}
}

// NewUPC returns new Provider that implemented GTIN-12 calculator.
func NewUPC() Provider {
	return &gtin{
		digit:              12,
		correctionPosition: true,
	}
}

// NewSSCC returns new Provider that implemented GTIN-18 calculator.
func NewSSCC() Provider {
	return &gtin{
		digit: 18,
	}
}
