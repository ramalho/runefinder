package main

import (
	"os"
	)

func Example() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", "EIGHTHS", "fraction"}
	main()
	// Output:
	// U+215C	⅜	VULGAR FRACTION THREE EIGHTHS (FRACTION THREE EIGHTHS)
	// U+215D	⅝	VULGAR FRACTION FIVE EIGHTHS (FRACTION FIVE EIGHTHS)
	// U+215E	⅞	VULGAR FRACTION SEVEN EIGHTHS (FRACTION SEVEN EIGHTHS)
	// 3 characters found
}

func Example_single_result() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", "registered"}
	main()
	// Output:
	// U+00AE	®	REGISTERED SIGN (REGISTERED TRADE MARK SIGN)
	// 1 character found
}

func Example_no_result() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", "nosuchcharacter"}
	main()
	// Output:
	// no character found
}
