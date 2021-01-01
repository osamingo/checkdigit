package checkdigit_test

import (
	"testing"

	"github.com/osamingo/checkdigit"
)

func TestVerhoeff_Verify(t *testing.T) {
	cases := map[string]struct {
		in  string
		out bool
	}{
		"Regular 1": {
			in:  "938472210036",
			out: true,
		},
		"Regular 2": {
			in:  "0973652",
			out: true,
		},
		"Regular 3": {
			in:  "27",
			out: true,
		},
		"Irregular 1": {
			in: "2361",
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
			ret := checkdigit.NewVerhoeff().Verify(c.in)
			if c.out != ret {
				t.Errorf("not equal, expected = %v, given = %v", c.out, ret)
			}
		})
	}
}

func TestVerhoeff_Generate(t *testing.T) {
	cases := map[string]struct {
		in      string
		out     int
		isError bool
	}{
		"Regular 1": {
			in:  "236",
			out: 3,
		},
		"Regular 2": {
			in:  "097365",
			out: 2,
		},
		"Regular 3": {
			in:  "93847221003",
			out: 6,
		},
		"Regular 4": {
			in:  "2",
			out: 7,
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
			r, err := checkdigit.NewVerhoeff().Generate(c.in)
			if c.isError && err == nil {
				t.Error("unexpected error")
			}
			if c.out != r {
				t.Errorf("not equal, expected = %d, given = %d", c.out, r)
			}
		})
	}
}
