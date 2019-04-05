package checkdigit

type damm struct {
	matrix [][]int
}

// Verify implements checkdigit.Verifier interface.
func (d *damm) Verify(code string) bool {
	return false
}

// Generate implements checkdigit.Generator interface.
func (d *damm) Generate(seed string) (rune, error) {
	return 0, nil
}
