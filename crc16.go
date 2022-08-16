package checkdigit

type CRC16 struct {
	name            string
	polynomial      uint16
	initial         uint16
	inputReflected  bool
	outputReflected bool
	finalXor        uint16
	table           [256]uint16
}

// Verify implements checkdigit.Verifier interface.
func (c CRC16) Verify(code string) bool {
	if len(code) < 2 {
		return false
	}
	i, err := c.Generate(code[:len(code)-1])

	return err == nil && i == int(code[len(code)-1]-'0')
}

// Generate implements checkdigit.Generator interface.
func (c *CRC16) Generate(seed string) (int, error) {
	if seed == "" {
		return 0, ErrInvalidArgument
	}

	interim := 0
	for _, n := range seed {
		if isNotNumber(n) {
			return 0, ErrInvalidArgument
		}

		// TODO (osamimgo): Write logic.
	}

	return interim, nil
}
