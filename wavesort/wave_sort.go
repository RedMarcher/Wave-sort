package wavesort

type Sequence []int

const UPRATE = 2

var Comparisons, Swaps, BlockSwaps int

// partition splits the sequence into two parts, the left part is less than the pivot and the right part is greater than the pivot.
// It returns the index of the split point which is involved in the greater part.
// l is the index of the left bound of the unsorted part.
// r is the index of the right bound of the unsorted part which is the firt element of the sorted part.
// p is the index of the pivot.
func (s Sequence) partition(l, r, p int) int {
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

func (s Sequence) less(i, j int) bool {
	Comparisons++
	return s[i] < s[j]
}

// blockSwap swaps the block from m to r and the one from r to p.
// m is the index of the first element of the left block.
// r is the index of the right bound of the left block and the first element of the right block.
// p is the index of the last element of the right block.
func (s Sequence) blockSwap(m, r, p int) {
	ll := r - m
	if ll == 0 {
		return
	}
	l := p - m + 1
	BlockSwaps += l // count block swapped elements
	tmp := s[m]
	init := m
	j := m
	nm := p - ll + 1
	k := j
	for range l {
		if j >= nm {
			k = j - nm + m
			if k == init {
				init++
				s[j] = tmp
				j = init
				tmp = s[j]
				continue
			}
			s[j] = s[k]
			j = k
			continue
		}
		k = j + ll
		s[j] = s[k]
		j = k
	}
}

// blockSwap_e swaps the block from m to r and the one from r to p.
// m is the index of the first element of the left block.
// r is the index of the right bound of the left block and the first element of the right block.
// p is the index of the last element of the right block.
func (s Sequence) blockSwap_e(m, r, p int) {
	ll := r - m
	if ll == 0 {
		return
	}

	lr := p - r + 1

	if lr == 1 {
		s[m], s[p] = s[p], s[m]
		Swaps++
		return
	}

	if lr <= ll {
		s.blockSwap_sr(m, r, p)
		BlockSwaps += lr << 1
		return
	}

	s.blockSwap_sl(m, p, ll)
	BlockSwaps += ll + lr
}

func (s Sequence) blockSwap_sl(m, p, ll int) {
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

func (s Sequence) blockSwap_sr(m, r, p int) {
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

func (s Sequence) swap(i, j int) {
	Swaps++
	s[i], s[j] = s[j], s[i]
}

func (s Sequence) downwave(start, sortedStart, end int) {
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
	s.blockSwap_e(m, sortedStart, p)
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

func (s Sequence) upwave(start, end int) {
	if start == end {
		return
	}
	sortedStart := s.preSorted(start, end)
	sortedLength := end - sortedStart + 1
	leftBound := end - sortedLength<<1 + 1
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

func (s Sequence) WaveSort() {
	if len(s) < 2 {
		return
	}
	s.upwave(0, len(s)-1)
}

func (s Sequence) upwave_e(start, end int) {
	length := end - start + 1
	if length == 1 {
		return
	}
	if length <= 8 {
		s.insertSort(start, end)
		return
	}
	sortedStart := s.preSorted(start, end)
	sortedLength := end - sortedStart + 1
	sortedLength <<= 4
	leftBound := end - sortedLength + 1
	for leftBound > start {
		s.downwave_e(leftBound, sortedStart, end)
		sortedStart = leftBound
		sortedLength = end - sortedStart + 1
		if length <= sortedLength<<6 {
			break
		}
		sortedLength <<= 4
		leftBound = end - sortedLength + 1
	}
	s.downwave_e(start, sortedStart, end)
}

func (s Sequence) downwave_e(start, sortedStart, end int) {
	unsortedLength := sortedStart - start
	if unsortedLength == 0 {
		return
	}
	lr := end - sortedStart
	if lr == 0 {
		if unsortedLength < 8 {
			s.insertSort(start, end)
			return
		}
		s.upwave_e(start, end)
		return
	}
	p := sortedStart + lr/2
	m := s.partition(start, sortedStart, p)
	if m == sortedStart {
		if p == sortedStart {
			s.upwave_e(start, sortedStart-1)
			return
		}
		s.downwave_e(start, sortedStart, p-1)
		return
	}
	s.blockSwap_e(m, sortedStart, p)
	if m == start {
		if p == sortedStart {
			s.upwave_e(m+1, end)
			return
		}
		s.downwave_e(m+p-sortedStart+1, p+1, end)
		return
	}
	if p == sortedStart {
		s.upwave_e(start, m-1)
		s.upwave_e(m+1, end)
		return
	}
	s.downwave_e(start, m, m+p-sortedStart-1)
	s.downwave_e(m+p-sortedStart+1, p+1, end)
}

func (s Sequence) WaveSort_e() {
	if len(s) < 2 {
		return
	}
	s.upwave_e(0, len(s)-1)
}

func (s Sequence) preSorted(start, end int) int {
	du := false
	i := end
	for i > start {
		if du {
			if s.less(i, i-1) {
				return i
			}
			i--
			continue
		}
		if s.less(i-1, i) {
			if i == end {
				du = true
				i--
				continue
			}
			s.reverse(i, end)
			return i
		}
		i--
	}
	if du {
		return start
	}
	s.reverse(start, end)
	return start
}

func (s Sequence) reverse(start, end int) {
	i := start
	j := end
	for range (end - start + 1) / 2 {
		s.swap(i, j)
		i++
		j--
	}
}

func (s Sequence) insertSort(l, r int) {
	sortedIndex := l + 1
	for sortedIndex <= r {
		i := sortedIndex
		for i > l {
			if s[sortedIndex] < s[i-1] {
				i--
				Comparisons++
				continue
			}
			break
		}
		if i < sortedIndex {
			tmp := s[sortedIndex]
			for j := sortedIndex; j > i; j-- {
				s[j] = s[j-1]
			}
			s[i] = tmp
			BlockSwaps += sortedIndex - i + 1
		}
		sortedIndex++
	}
}
