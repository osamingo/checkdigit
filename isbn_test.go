package checkdigit_test

import (
	"testing"

	"github.com/osamingo/checkdigit"
)

func TestIsbn10_Verify(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in  string
		out bool
	}{
		"Regular 1": {
			in:  "0026515628",
			out: true,
		},
		"Regular 2": {
			in:  "007231592X",
			out: true,
		},
		"Regular 3": {
			in:  "155860832X",
			out: true,
		},
		"Irregular 1": {
			in: "155860831X",
		},
		"Irregular 2": {
			in: "9780002715096",
		},
		"Irregular 3": {
			in: "155860831",
		},
		"Irregular 4": {
			in: "",
		},
		"Irregular 5": {
			in: "aaaaaaaaaa",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ret := checkdigit.NewISBN10().Verify(c.in)
			if c.out != ret {
				t.Errorf("not equal, expected = %v, given = %v", c.out, ret)
			}
		})
	}
}

func TestIsbn10_Generate(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in      string
		out     int
		isError bool
	}{
		"Regular 1": {
			in:  "002651562",
			out: 8,
		},
		"Regular 2": {
			in:  "007231592",
			out: 10,
		},
		"Regular 3": {
			in:  "155860832",
			out: 10,
		},
		"Irregular 1": {
			isError: true,
		},
		"Irregular 2": {
			in:      "9780002715096",
			isError: true,
		},
		"Irregular 3": {
			in:      "15586",
			isError: true,
		},
		"Irregular 4": {
			in:      "155860832X",
			isError: true,
		},
		"Irregular 5": {
			in:      "aaaaaaaaa",
			isError: true,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			r, err := checkdigit.NewISBN10().Generate(c.in)
			if c.isError && err == nil {
				t.Error("unexpected error")
			}
			if c.out != r {
				t.Errorf("not equal, expected = %d, given = %d", c.out, r)
			}
		})
	}
}

func TestIsbn13_Verify(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in  string
		out bool
	}{
		"Regular 1": {
			in:  "9780002712095",
			out: true,
		},
		"Regular 2": {
			in:  "9780002715096",
			out: true,
		},
		"Regular 3": {
			in:  "9780002713306",
			out: true,
		},
		"Irregular 1": {
			in: "155860831X",
		},
		"Irregular 2": {
			in: "9780002712520",
		},
		"Irregular 3": {
			in: "9780002712709",
		},
		"Irregular 4": {
			in: "",
		},
		"Irregular 5": {
			in: "aaaaaaaaaaaaa",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ret := checkdigit.NewISBN13().Verify(c.in)
			if c.out != ret {
				t.Errorf("not equal, expected = %v, given = %v", c.out, ret)
			}
		})
	}
}

func TestIsbn13_Generate(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in      string
		out     int
		isError bool
	}{
		"Regular 1": {
			in:  "978000271217",
			out: 0,
		},
		"Regular 2": {
			in:  "978000271330",
			out: 6,
		},
		"Regular 3": {
			in:  "978000271363",
			out: 4,
		},
		"Regular 4": {
			in:  "978000271236",
			out: 1,
		},
		"Irregular 1": {
			isError: true,
		},
		"Irregular 2": {
			in:      "0026515628",
			isError: true,
		},
		"Irregular 3": {
			in:      "155860832X",
			isError: true,
		},
		"Irregular 4": {
			in:      "a",
			isError: true,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			r, err := checkdigit.NewISBN13().Generate(c.in)
			if c.isError && err == nil {
				t.Error("unexpected error")
			}
			if c.out != r {
				t.Errorf("not equal, expected = %d, given = %d", c.out, r)
			}
		})
	}
}
