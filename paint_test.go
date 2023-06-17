package jsonpainter_test

import (
	"fmt"
	"testing"

	"github.com/michurin/jsonpainter"
)

func TestString(t *testing.T) {
	cases := []struct {
		name string
		in   string
		exp  string
	}{{
		name: "empty",
		in:   "",
		exp:  "",
	}, {
		name: "simple_value_token",
		in:   `{"one":true}`,
		exp:  `(C]{[O)(Q]"one"[O)(C]:[O)(S]true[O)(C]}[O)`,
	}, {
		name: "simple_value_string",
		in:   `{"one":"two"}`,
		exp:  `(C]{[O)(Q]"one"[O)(C]:[O)(s]"two"[O)(C]}[O)`,
	}, {
		name: "spaces",
		in:   ` { "one" : 12 } `,
		exp:  ` (C]{[O) (Q]"one"[O) (C]:[O) (S]12[O) (C]}[O) `,
	}, {
		name: "spaces_array",
		in:   ` [ "12" ] `,
		exp:  ` (C][[O) (s]"12"[O) (C]][O) `,
	}, {
		name: "escaped",
		in:   `{"o\"ne":1}`,
		exp:  `(C]{[O)(Q]"o\"ne"[O)(C]:[O)(S]1[O)(C]}[O)`,
	}, {
		name: "ignore_wrong_colon",
		in:   `{1:2}`,
		exp:  `(C]{[O)(S]1[O):(S]2[O)(C]}[O)`,
	}, {
		name: "invalid_has_to_be_closed",
		in:   `[1`,
		exp:  `(C][[O)(S]1[O)`,
	}, {
		name: "overclosing",
		in:   `[]]`,
		exp:  `(C][[O)(C]][O)]`,
	}, {
		name: "bug_2023_06_1a",
		in:   `""x`,
		exp:  `""x`, // controversial behavior: no open brackets no color
	}, {
		name: "bug_2023_06_1b",
		in:   `{} "" x`,
		exp:  `(C]{[O)(C]}[O) "" x`,
	}, {
		name: "bug_2023_06_1c",
		in:   `{""x`,
		exp:  `(C]{[O)(s]""[O)(S]x[O)`, // JSON is broken, however we color it. Not sure it is correct behavior
	}}
	// Setup fake escape sequences
	optsColors := []jsonpainter.Option{
		jsonpainter.ClrKey(jsonpainter.Color("(Q]")),
		jsonpainter.ClrStr(jsonpainter.Color("(s]")),
		jsonpainter.ClrSpecStr(jsonpainter.Color("(S]")),
		jsonpainter.ClrCtl(jsonpainter.Color("(C]")),
		jsonpainter.ClrOff(jsonpainter.Color("[O)")),
	}
	// Disable all kinds of highlighting
	optsNoColors := []jsonpainter.Option{
		jsonpainter.ClrKey(jsonpainter.None),
		jsonpainter.ClrStr(jsonpainter.None),
		jsonpainter.ClrSpecStr(jsonpainter.None),
		jsonpainter.ClrCtl(jsonpainter.None),
		jsonpainter.ClrOff(jsonpainter.None),
	}
	// Check highlighting
	for _, c := range cases {
		c := c
		t.Run("color_"+c.name, func(t *testing.T) {
			o := jsonpainter.String(c.in, optsColors...)
			if o != c.exp {
				t.Errorf("%s != %s", o, c.exp)
			}
		})
	}
	// Check nothing changes if highlighting is disabled
	for _, c := range cases {
		c := c
		t.Run("no_color_"+c.name, func(t *testing.T) {
			o := jsonpainter.String(c.in, optsNoColors...)
			if o != c.in {
				t.Errorf("%s != %s", o, c.in)
			}
		})
	}
}

func ExampleString_simplest() {
	// Example of log line with mixture JSON and non-JSON parts
	logline := `Req: {"rc": 1}`
	// As far as go examples doesn't support escape control sequences, it has to be used %q here.
	// In real application this line might look like fmt.Println(jsonpainter.String(logline))
	fmt.Printf("%q", jsonpainter.String(logline))
	// Output: "Req: \x1b[31;1m{\x1b[0m\x1b[33;1m\"rc\"\x1b[0m\x1b[31;1m:\x1b[0m \x1b[36;1m1\x1b[0m\x1b[31;1m}\x1b[0m"
}

func ExampleString_withOptions() {
	opts := []jsonpainter.Option{
		jsonpainter.ClrCtl(jsonpainter.None),     // We don't want to colorize JSON punctuation
		jsonpainter.ClrKey(jsonpainter.Cyan),     // Keys will be cyan
		jsonpainter.ClrSpecStr(jsonpainter.None), // Special values (true, false, null...) won't be painted
		jsonpainter.ClrStr(jsonpainter.None),     // Strings won't be painted as well
	}
	logline := `Req: {"rc": 1}`
	fmt.Printf("%q", jsonpainter.String(logline, opts...))
	// Output: "Req: {\x1b[36;1m\"rc\"\x1b[0m: 1}"
}
