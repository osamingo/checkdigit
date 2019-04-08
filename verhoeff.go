package checkdigit

type verhoeff struct {
	multiplication [][]int
	permutation    [][]int
	inverse        []int
}

// Verify implements checkdigit.Verifier interface.
func (v *verhoeff) Verify(code string) bool {

	if len(code) < 2 {
		return false
	}

	i, err := v.Generate(code[:len(code)-1])

	return err == nil && i == int(code[len(code)-1]-'0')
}

// Generate implements checkdigit.Generator interface.
func (v *verhoeff) Generate(seed string) (int, error) {

	if seed == "" {
		return 0, ErrInvalidArgument
	}

	interim, l := 0, len(seed)
	for i, n := range seed {
		if isNotNumber(n) {
			return 0, ErrInvalidArgument
		}
		interim = v.multiplication[interim][v.permutation[(i+1)%8][seed[l-i-1]-'0']]
	}

	return v.inverse[interim], nil
}
