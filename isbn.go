package checkdigit

type (
	isbn10 struct{}
	isbn13 struct{}
)

// Verify implements checkdigit.Verifier interface.
func (i10 isbn10) Verify(code string) bool {
	if len(code) != 10 {
		return false
	}

	sum, multiply := 0, 10
	for _, n := range code {
		var digit int
		switch {
		case n == 'X':
			digit = 10
		case isNotNumber(n):
			return false
		default:
			digit = int(n - '0')
		}

		sum += multiply * digit
		multiply--
	}

	return sum%11 == 0
}

// Generate implements checkdigit.Generator interface.
// This will return a "10" instead of an "X", since the interface expects an int.
func (i10 *isbn10) Generate(seed string) (int, error) {
	if len(seed) != 9 {
		return 0, ErrInvalidArgument
	}

	sum, multiply := 0, 10

	for _, n := range seed {
		if isNotNumber(n) {
			return 0, ErrInvalidArgument
		}

		sum += multiply * int(n-'0')
		multiply--
	}

	return 11 - sum%11, nil
}

// Verify implements checkdigit.Verifier interface.
func (i13 isbn13) Verify(code string) bool {
	if len(code) != 13 {
		return false
	}
	i, err := i13.Generate(code[:len(code)-1])

	return err == nil && i == int(code[len(code)-1]-'0')
}

// Generate implements checkdigit.Generator interface.
func (i13 *isbn13) Generate(seed string) (int, error) {
	if len(seed) != 12 {
		return 0, ErrInvalidArgument
	}

	sum, weight := 0, 1
	for _, n := range seed {
		if isNotNumber(n) {
			return 0, ErrInvalidArgument
		}
		sum += int(n-'0') * weight
		if weight == 1 {
			weight = 3
		} else {
			weight = 1
		}
	}

	d := 10 - sum%10
	if d == 10 {
		d = 0
	}

	return d, nil
}
