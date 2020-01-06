// 不要
package main

type Scanner struct {
	i  int
	rs []rune
}

func NewScanner(s string) *Scanner {
	return &Scanner{
		i:  -1,
		rs: ([]rune)(s),
	}
}

func (s *Scanner) Scan() bool {
	s.i++
	return len(s.rs) != 0 && s.i < len(s.rs)
}

func (s *Scanner) Rune() rune {
	if s.i < 0 || s.i >= len(s.rs) {
		return 0
	}
	return s.rs[s.i]
}

func (s *Scanner) Next() (rune, bool) {
	if s.i+1 < len(s.rs) {
		return s.rs[s.i+1], true
	}
	return 0, false
}
