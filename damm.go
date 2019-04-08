package checkdigit

type damm struct {
	matrix [][]int
}

// Verify implements checkdigit.Verifier interface.
func (d *damm) Verify(code string) bool {

	if len(code) < 2 {
		return false
	}

	i, err := d.Generate(code[:len(code)-1])

	return err == nil && i == int(code[len(code)-1]-'0')
}

// Generate implements checkdigit.Generator interface.
func (d *damm) Generate(seed string) (int, error) {

	if seed == "" {
		return 0, ErrInvalidArgument
	}

	interim := 0
	for _, n := range seed {
		if isNotNumber(n) {
			return 0, ErrInvalidArgument
		}
		interim = d.matrix[interim][int(n-'0')]
	}

	return interim, nil
}
