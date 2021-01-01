package checkdigit_test

import (
	"testing"

	"github.com/osamingo/checkdigit"
)

func TestDamm_Verify(t *testing.T) {
	cases := map[string]struct {
		in  string
		out bool
	}{
		"Regular 1": {
			in:  "224564332323",
			out: true,
		},
		"Regular 2": {
			in:  "543525432346",
			out: true,
		},
		"Regular 3": {
			in:  "37",
			out: true,
		},
		"Irregular 1": {
			in: "835323233227",
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
			ret := checkdigit.NewDamm().Verify(c.in)
			if c.out != ret {
				t.Errorf("not equal, expected = %v, given = %v", c.out, ret)
			}
		})
	}
}

func TestDamm_Generate(t *testing.T) {
	cases := map[string]struct {
		in      string
		out     int
		isError bool
	}{
		"Regular 1": {
			in:  "22456433232",
			out: 3,
		},
		"Regular 2": {
			in:  "54352543234",
			out: 6,
		},
		"Regular 3": {
			in:  "10493839530",
			out: 5,
		},
		"Regular 4": {
			in:  "08989435403",
			out: 5,
		},
		"Regular 5": {
			in:  "54994384990",
			out: 4,
		},
		"Regular 6": {
			in:  "3",
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
			r, err := checkdigit.NewDamm().Generate(c.in)
			if c.isError && err == nil {
				t.Error("unexpected error")
			}
			if c.out != r {
				t.Errorf("not equal, expected = %d, given = %d", c.out, r)
			}
		})
	}
}
