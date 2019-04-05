package checkdigit

type verhoeff struct {
	multiplication [][]int
	permutation    [][]int
	inverse        []int
}

// Verify implements checkdigit.Verifier interface.
func (v *verhoeff) Verify(code string) bool {
	return false
}

// Generate implements checkdigit.Generator interface.
func (v *verhoeff) Generate(seed string) (rune, error) {
	return 0, nil
}
