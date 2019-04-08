package checkdigit_test

import (
	"testing"

	"github.com/osamingo/checkdigit"
)

func TestGtin_Verify(t *testing.T) {

	cases := map[string]struct {
		fn  func() checkdigit.Provider
		in  string
		out bool
	}{
		"EAN8": {
			fn:  checkdigit.NewEAN8,
			in:  "96385074",
			out: true,
		},
		"EAN13": {
			fn:  checkdigit.NewEAN13,
			in:  "5901234123457",
			out: true,
		},
		"JAN8": {
			fn:  checkdigit.NewJAN8,
			in:  "49968712",
			out: true,
		},
		"JAN13": {
			fn:  checkdigit.NewJAN13,
			in:  "4569951116179",
			out: true,
		},
		"ITF": {
			fn:  checkdigit.NewITF,
			in:  "14569951116176",
			out: true,
		},
		"UPC": {
			fn:  checkdigit.NewUPC,
			in:  "012345678905",
			out: true,
		},
		"SSCC": {
			fn:  checkdigit.NewSSCC,
			in:  "045699511100000016",
			out: true,
		},
		"Empty": {
			fn: checkdigit.NewEAN8,
			in: "",
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			ret := c.fn().Verify(c.in)
			if c.out != ret {
				t.Errorf("not equal, expected = %v, given = %v", c.out, ret)
			}
		})
	}
}

func TestGtin_Generate(t *testing.T) {

	cases := map[string]struct {
		fn      func() checkdigit.Provider
		in      string
		out     int
		isError bool
	}{
		"EAN8": {
			fn:  checkdigit.NewEAN8,
			in:  "9638112",
			out: 0,
		},
		"EAN13": {
			fn:  checkdigit.NewEAN13,
			in:  "590123412345",
			out: 7,
		},
		"JAN8": {
			fn:  checkdigit.NewJAN8,
			in:  "4996871",
			out: 2,
		},
		"JAN13": {
			fn:  checkdigit.NewJAN13,
			in:  "456995111617",
			out: 9,
		},
		"ITF": {
			fn:  checkdigit.NewITF,
			in:  "1456995111617",
			out: 6,
		},
		"UPC": {
			fn:  checkdigit.NewUPC,
			in:  "01234567890",
			out: 5,
		},
		"SSCC": {
			fn:  checkdigit.NewSSCC,
			in:  "04569951110000001",
			out: 6,
		},
		"Empty": {
			fn:      checkdigit.NewEAN8,
			in:      "",
			isError: true,
		},
		"Alphabet": {
			fn:      checkdigit.NewEAN8,
			in:      "aaaaaaa",
			isError: true,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			r, err := c.fn().Generate(c.in)
			if c.isError && err == nil {
				t.Error("unexpected error")
			}
			if c.out != r {
				t.Errorf("not equal, expected = %d, given = %d", c.out, r)
			}
		})
	}
}
