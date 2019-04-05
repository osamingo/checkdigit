package checkdigit

type luhn struct{}

// Verify implements checkdigit.Verifier interface.
func (l *luhn) Verify(code string) bool {
	return false
}

// Generate implements checkdigit.Generator interface.
func (l *luhn) Generate(seed string) (rune, error) {
	return 0, nil
}
