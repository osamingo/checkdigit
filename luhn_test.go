package checkdigit

import "testing"

func TestLuhn_Verify(t *testing.T) {

	cases := map[string]struct {
		in  string
		out bool
	}{
		"Regular 1": {
			in:  "4242424242424242",
			out: true,
		},
		"Regular 2": {
			in:  "5105105105105100",
			out: true,
		},
		"Regular 3": {
			in:  "34",
			out: true,
		},
		"Irregular 1": {
			in: "510510510510511",
		},
		"Irregular 2": {
			in: "",
		},
		"Irregular 3": {
			in: "a",
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			ret := NewLuhnProvider().Verify(c.in)
			if c.out != ret {
				t.Errorf("not equal, expected = %v, given = %v", c.out, ret)
			}
		})
	}
}

func TestLuhn_Generate(t *testing.T) {

	cases := map[string]struct {
		in      string
		out     int
		isError bool
	}{
		"Regular 1": {
			in:  "424242424242424",
			out: 2,
		},
		"Regular 2": {
			in:  "510510510510510",
			out: 0,
		},
		"Regular 3": {
			in:  "37144963539843",
			out: 1,
		},
		"Regular 4": {
			in:  "3056930902590",
			out: 4,
		},
		"Regular 5": {
			in:  "353011133330000",
			out: 0,
		},
		"Regular 6": {
			in:  "3",
			out: 4,
		},
		"Irregular 1": {
			isError: true,
		},
		"Irregular 2": {
			in:      "a",
			isError: true,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			r, err := NewLuhnProvider().Generate(c.in)
			if c.isError && err == nil {
				t.Error("unexpected error")
			}
			if c.out != r {
				t.Errorf("not equal, expected = %d, given = %d", c.out, r)
			}
		})
	}
}
