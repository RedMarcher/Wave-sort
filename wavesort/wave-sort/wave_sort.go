package wave_sort

type seq []int

var compare, swap, blockSwap int

// partition splits the sequence into two parts, the left part is less than the pivot and the right part is greater than the pivot.
// It returns the index of the split point which is involved in the greater part.
// l is the index of the left bound of the unsorted part.
// r is the index of the right bound of the unsorted part which is the firt element of the sorted part.
// p is the index of the pivot.
func (s seq) partition(l, r, p int) int {
	i, j := l-1, r

	for {
		for {
			i++
			if i == j {
				return i
			}
			if s.less(p, i) {
				break
			}
		}
		for {
			j--
			if j == i {
				return i
			}
			if s.less(j, p) {
				break
			}
		}
		s.swap(i, j)
	}
}

func (s seq) less(i, j int) bool {
	compare++
	return s[i] < s[j]
}

// blockSwap swaps the block from m to r and the one from r to p.
// m is the index of the first element of the left block.
// r is the index of the right bound of the left block and the first element of the right block.
// p is the index of the last element of the right block.
func (s seq) blockSwap(m, r, p int) {
	ll := r - m
	if ll == 0 {
		return
	}

	lr := p - r + 1

	if lr == 1 {
		s[m], s[p] = s[p], s[m]
		swap++
		return
	}

	if lr <= ll {
		s.blockSwap_sr(m, r, p)
		blockSwap += lr << 1
		return
	}

	s.blockSwap_sl(m, p, ll)
	blockSwap += ll + lr
}

func (s seq) blockSwap_sl(m, p, ll int) {
	tmp := s[m]
	init := m
	j := m
	nm := p - ll + 1
	var k int
	for range p - m + 1 {
		if j >= nm {
			k = j - nm + m
			if k == init {
				init++
				s[j] = tmp
				j = init
				tmp = s[j]
				continue
			}
		} else {
			k = j + ll
		}
		s[j] = s[k]
		j = k
	}
}

func (s seq) blockSwap_sr(m, r, p int) {
	i := m
	tmp := s[i]
	j := r
	for j < p {
		s[i] = s[j]
		i++
		s[j] = s[i]
		j++
	}
	s[i] = s[j]
	s[j] = tmp
}

func (s seq) swap(i, j int) {
	swap++
	s[i], s[j] = s[j], s[i]
}

func (s seq) downwave(start, sortedStart, end int) {
	if sortedStart-start == 0 {
		return
	}
	p := sortedStart + (end-sortedStart)/2
	m := s.partition(start, sortedStart, p)
	if m == sortedStart {
		if p == sortedStart {
			s.upwave(start, sortedStart-1)
			return
		}
		s.downwave(start, sortedStart, p-1)
		return
	}
	s.blockSwap(m, sortedStart, p)
	if m == start {
		if p == sortedStart {
			s.upwave(m+1, end)
			return
		}
		p++
		s.downwave(m+p-sortedStart, p, end)
		return
	}
	if p == sortedStart {
		s.upwave(start, m-1)
		s.upwave(m+1, end)
		return
	}
	s.downwave(start, m, m+p-sortedStart-1)
	s.downwave(m+p-sortedStart+1, p+1, end)
}

func (s seq) upwave(start, end int) {
	if start == end {
		return
	}
	sortedStart := end
	sortedLength := 1
	leftBound := end - 1
	length := end - start + 1
	for leftBound > start {
		s.downwave(leftBound, sortedStart, end)
		sortedStart = leftBound
		sortedLength = end - sortedStart + 1
		if length < sortedLength<<2 {
			break
		}
		leftBound = end - sortedLength<<1 + 1
	}
	s.downwave(start, sortedStart, end)
}

func (s seq) WaveSort() {
	if len(s) < 2 {
		return
	}
	s.upwave(0, len(s)-1)
}
