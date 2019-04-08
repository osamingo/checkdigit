package checkdigit

type gtin struct {
	digit              int
	correctionPosition bool
}

func (g *gtin) Verify(code string) bool {

	if len(code) != g.digit {
		return false
	}

	i, err := g.Generate(code[:len(code)-1])

	return err == nil && i == int(code[len(code)-1]-'0')
}

func (g *gtin) Generate(seed string) (int, error) {

	if len(seed) != g.digit-1 {
		return 0, ErrInvalidArgument
	}

	var oddSum, evenSum int
	for i, n := range seed {
		if isNotNumber(n) {
			return 0, ErrInvalidArgument
		}
		p := i
		if g.correctionPosition {
			p++
		}
		if p%2 == 0 {
			oddSum += int(n - '0')
		} else {
			evenSum += int(n - '0')
		}
	}

	d := 10 - (oddSum*33+evenSum)%10
	if d == 10 {
		d = 0
	}

	return d, nil
}
