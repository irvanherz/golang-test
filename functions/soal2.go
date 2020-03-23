package functions

import (
	"strconv"
)

type Mystery struct {
	Str   string
	Index int
}

func (m *Mystery) Init(str string) {
	m.Str = str
	m.Index = 0
}

func (m *Mystery) get(length int) int {
	a := m.Index
	b := m.Index + length

	if a == len(m.Str) {
		return -1
	}

	if b > len(m.Str) {
		return -2
	}
	val, _ := strconv.Atoi(m.Str[a:b])
	return val
}

func (m *Mystery) skip(length int) {
	m.Index += length
}

func countDigits(i int) (count int) {
	for i != 0 {
		i /= 10
		count = count + 1
	}
	return count
}

func (m *Mystery) Scan() []int {
	var r = make([]int, 0)
	ndigits := 1
begin:
	c := m.get(ndigits)
	if c == -2 {
		goto end
	}
	m.skip(ndigits)
	for {
		g1 := c + 1
		g2 := c + 2
		nlen := countDigits(g1)
		n := m.get(nlen)

		if n == g1 {
			c = g1
			if c == -1 {
				goto end
			}
			m.skip(nlen)
			continue
		} else {
			nlen = countDigits(g2)
			n = m.get(nlen)

			if n == -1 && len(r) > 0 { //end of item
				goto end
			} else if n == g2 {
				c = g2
				r = append(r, g1)
				m.skip(nlen)
				continue
			} else {
				m.Index = 0
				r = nil
				ndigits++
				goto begin
			}

		}
	}
end:
	return r
}
