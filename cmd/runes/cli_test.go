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
	// U+215C	⅜	VULGAR FRACTION THREE EIGHTHS
	// U+215D	⅝	VULGAR FRACTION FIVE EIGHTHS
	// U+215E	⅞	VULGAR FRACTION SEVEN EIGHTHS
	// 3 characters found
}

func Example_single_result() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", "registered"}
	main()
	// Output:
	// U+00AE	®	REGISTERED SIGN
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

func Example_get_names() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", "?abc"}
	main()
	// Output:
	// U+003F	?	QUESTION MARK
	// U+0061	a	LATIN SMALL LETTER A
	// U+0062	b	LATIN SMALL LETTER B
	// U+0063	c	LATIN SMALL LETTER C
	// 4 characters found
}

func Example_no_args() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{""}
	main()
	// Output:
	// Please provide one or more words or characters to search.
}
