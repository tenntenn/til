package main

import (
	"testing"
	"unicode"
)

type pair [2]interface{}

func __(t *testing.T, pattern string) []func(s *Scanner) interface{} {
	fs := make([]func(s *Scanner) interface{}, len(pattern))

	for i, p := range pattern {
		switch unicode.ToUpper(p) {
		case 'S':
			fs[i] = func(s *Scanner) interface{} {
				return s.Scan()
			}
		case 'R':
			fs[i] = func(s *Scanner) interface{} {
				return s.Rune()
			}
		case 'N':
			fs[i] = func(s *Scanner) interface{} {
				r, ok := s.Next()
				return pair{r, ok}
			}
		default:
			t.Fatal("unexpected pattern:", p)
		}
	}

	return fs
}

func TestScanner(t *testing.T) {
	type V = []interface{}
	const none rune = 0
	cases := map[string]struct {
		str     string
		pattern string
		want    []interface{}
	}{
		// S:Scan, R:Rune, N:Next
		"empty-R":   {"", "R", V{none}},
		"empty-S":   {"", "S", V{false}},
		"empty-N":   {"", "N", V{pair{none, false}}},
		"empty-SR":  {"", "SR", V{false, none}},
		"empty-SN":  {"", "SN", V{false, pair{none, false}}},
		"single-R":  {"a", "R", V{none}},
		"single-S":  {"a", "S", V{true}},
		"single-N":  {"a", "N", V{pair{'a', true}}},
		"single-SR": {"a", "SR", V{true, 'a'}},
		"single-SS": {"a", "SS", V{true, false}},
		"single-SN": {"a", "SN", V{true, pair{none, false}}},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			s := NewScanner(tt.str)
			got := make([]interface{}, len(tt.pattern))
			for i, f := range __(t, tt.pattern) {
				got[i] = f(s)
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("%dth result want %v but got %v", i, tt.want[i], got[i])
				}
			}
		})
	}
}
