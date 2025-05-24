package algorithm

import (
	"math"
	"slices"
)

type nums []int

func (nms nums) insertSort() {
	ln := len(nms)
	sortedIndex := 1
	for sortedIndex < ln {
		tmp := nms[sortedIndex]
		i := sortedIndex
		for i > 0 {
			if tmp < nms[i-1] {
				nms[i] = nms[i-1]
				i--
				continue
			}
			break
		}
		nms[i] = tmp
		sortedIndex++
	}
}

func (nms nums) shellSort() {
	ln := len(nms)
	// https://en.wikipedia.org/wiki/Shellsort
	// The fastes one seems to be floor(2.25k)
	// https://stackoverflow.com/questions/2539545/fastest-gap-sequence-for-shell-sort
	incs := []int{1391376, 463792, 198768, 86961,
		33936, 13776, 4592, 1968, 861,
		336, 112, 48, 21, 7, 3, 1}
	for k := 0; k < len(incs); k++ {
		h := incs[k]
		for i := h; i < ln; i++ {
			for j := i; j >= h; j -= h {
				if nms[j] < nms[j-h] {
					nms[j], nms[j-h] = nms[j-h], nms[j]
					continue
				}
				break
			}
		}
	}
}

func merge(a, aux nums, l, m, r int) {
	if a[m] > a[m-1] {
		return
	}
	for i := l; i < m; i++ {
		aux[i] = a[i]
	}
	for j := m; j < r; j++ {
		aux[j] = a[m-j+r-1]
	}
	i := l
	j := r - 1
	for k := l; k < r; k++ {
		if aux[j] < aux[i] {
			a[k] = aux[j]
			j--
			continue
		}
		a[k] = aux[i]
		i++
	}
}

func merge_1(a, aux nums, l, m, r int) {
	if a[m] > a[m-1] {
		return
	}
	copy(aux[l:r], a[l:r])
	i := l
	j := m
	for k := l; k < r; k++ {
		if i >= m {
			a[k] = aux[j]
			j++
			continue
		}
		if j >= r {
			a[k] = aux[i]
			i++
			continue
		}
		if aux[i] > aux[j] {
			a[k] = aux[j]
			j++
			continue
		}
		a[k] = aux[i]
		i++
	}
}

func (nms nums) mergeSort() {
	ln := len(nms)
	aux := make(nums, ln)

	for m := 1; m < ln; m += m {
		for i := 0; i < ln-m; i += m + m {
			merge_1(nms, aux, i, i+m, int(math.Min(float64(i+m+m), float64(ln))))
		}
	}
}

func recursiveMerge(a, aux nums, low, mid, high int) {
	i := low
	j := mid + 1
	if aux[j] > aux[mid] {
		copy(a[low:high+1], aux[low:high+1])
		return
	}
	for k := low; k <= high; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
			continue
		}
		if j > high {
			a[k] = aux[i]
			i++
			continue
		}
		if aux[i] > aux[j] {
			a[k] = aux[j]
			j++
			continue
		}
		a[k] = aux[i]
		i++
	}
}

func recursiveMergeSort(a, aux nums, low, high int) {
	if high <= low {
		return
	}
	mid := low + (high-low)/2
	recursiveMergeSort(aux, a, low, mid)
	recursiveMergeSort(aux, a, mid+1, high)
	recursiveMerge(a, aux, low, mid, high)
}

func (nms nums) recursiveMergeSort() {
	aux := slices.Clone(nms)
	recursiveMergeSort(nms, aux, 0, len(nms)-1)
}

func (nms nums) partition(l, r int) int {
	i, j := l-1, r

	for {
		for {
			i++
			if i == j {
				if j != r {
					nms.swap(i, r)
				}
				return i
			}
			if nms.less(r, i) {
				break
			}
		}
		for {
			j--
			if j == i {
				nms.swap(i, r)
				return i
			}
			if nms.less(j, r) {
				break
			}
		}
		nms.swap(i, j)
	}
}

func (nms nums) quickSort(l, r int) {
	if r <= l {
		return
	}
	m := nms.partition(l, r)
	nms.quickSort(l, m-1)
	nms.quickSort(m+1, r)
}

func (nms nums) qSort() {
	// rand.Shuffle(len(nms), nms.swap)
	nms.quickSort(0, len(nms)-1)
	// log.Printf("\nnumber of comparisons: %d\n", compare)
	// log.Printf("number of swaps: %d\n", swap)
}

func (nms nums) less(i, j int) bool {
	compare++
	return nms[i] < nms[j]
}

func (nms nums) swap(i, j int) {
	swap++
	nms[i], nms[j] = nms[j], nms[i]
}
