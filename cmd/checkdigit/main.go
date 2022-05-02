package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/osamingo/checkdigit"
)

func main() {
	const (
		errExitCode = 1
		generateCmd = "generate"
		verifyCmd   = "verify"
	)

	flag.Usage = usage

	var pname string
	flag.StringVar(&pname, "provider", "luhn", "")
	flag.StringVar(&pname, "provider", "luhn", "")
	flag.Parse()

	provider := takeProvider(pname)
	if provider == nil {
		fmt.Fprintf(os.Stderr, "Unimplemented %s provider\n", pname)
		os.Exit(errExitCode)
	}

	args := flag.Args()
	if len(args) < 2 {
		flag.Usage()
	}

	switch args[0] {
	case generateCmd:
		if err := generate(provider, args[1]); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(errExitCode)
		}
	case verifyCmd:
		if !verify(provider, args[1]) {
			os.Exit(errExitCode)
		}
	default:
		flag.Usage()
	}
}

//nolint: cyclop
func takeProvider(name string) checkdigit.Provider {
	switch strings.ToLower(strings.ReplaceAll(name, "-", "")) {
	case "luhn":
		return checkdigit.NewLuhn()
	case "verhoeff":
		return checkdigit.NewVerhoeff()
	case "damm":
		return checkdigit.NewDamm()
	case "isbn10":
		return checkdigit.NewISBN10()
	case "isbn13", "isbn":
		return checkdigit.NewISBN13()
	case "ean8":
		return checkdigit.NewEAN8()
	case "ean13", "ean":
		return checkdigit.NewEAN13()
	case "jan8":
		return checkdigit.NewJAN8()
	case "jan13", "jan":
		return checkdigit.NewJAN13()
	case "itf":
		return checkdigit.NewITF()
	case "upc":
		return checkdigit.NewUPC()
	case "sscc":
		return checkdigit.NewSSCC()
	}

	return nil
}

func generate(g checkdigit.Generator, seed string) error {
	cd, err := g.Generate(seed)
	if err != nil {
		return fmt.Errorf("failed to generate with seed, message: %w", err)
	}
	if cd > 9 {
		fmt.Fprintf(os.Stderr, "%s\n", "X")
	} else {
		fmt.Fprintf(os.Stderr, "%d\n", cd)
	}

	return nil
}

func verify(v checkdigit.Verifier, target string) bool {
	ret := v.Verify(target)
	fmt.Fprintf(os.Stderr, "%t\n", ret)

	return ret
}

func usage() {
	const usage = `NAME:
    checkdigit - generate or verify numbers

USAGE:
    checkdigit [global options] command [command options] [arguments...]

COMMANDS:
    generate	Generates a check digit by implemented algorithm or calculator.
    verify	Verify to code by implemented algorithm or calculator.
    help	Shows a list of commands or help for one command

GLOBAL OPTIONS
    -p, --provider	Select use provider [default: luhn]`
	fmt.Fprintln(os.Stderr, usage)
	os.Exit(2)
}
