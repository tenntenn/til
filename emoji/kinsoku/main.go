package main

import (
	"fmt"
)

type Measurer interface {
	Measure(r rune) (w, h float64)
}

type MeasurerFunc func(r rune) (w, h float64)

func (m MeasurerFunc) Measure(r rune) (w, h float64) {
	return m(r)
}

func main() {
	m := MeasurerFunc(func(r rune) (w, h float64) {
		return 1, 1
	})

	for _, s := range WordWrap("123。456[8", m, 3) {
		fmt.Println(s)
	}
}

func WordWrap(s string, m Measurer, width float64) []string {

	rs := ([]rune)(s)

	var (
		lines []string
		//line  []rune
		w	  float64
		start, i int
	)

	for i < len(rs) {
		_, dw := m.Measure(rs[i])
		if w + dw <= width {
			w += dw
			i++
			continue
		}

		for i-start > 0 && i-1 > 0 && gyomatsuKinsoku[rs[i-1]] {
			_, dw := m.Measure(rs[i])
			w -= dw
			if w < 0 {
				w = 0
			}
			i--
		}

		for i-start > 0 && gyotouKinsoku[rs[i]] {
			_, dw := m.Measure(rs[i])
			w -= dw
			if w < 0 {
				w = 0
			}
			i--
		}
	
		lines = append(lines, string(rs[start:i]))
		start = i
		w = 0
	}

	if i-start > 0 {
		lines = append(lines, string(rs[start:i]))
	}

	return lines
}
