package algorithms

import (
	"math"
	"math/rand/v2"
	"slices"
)

type Nums []int

var Comparisons, Swaps, BlockSwaps int

func (nms Nums) InsertSort() {
	ln := len(nms)
	sortedIndex := 1
	for sortedIndex < ln {
		tmp := nms[sortedIndex]
		i := sortedIndex
		for i > 0 {
			Comparisons++
			if tmp < nms[i-1] {
				nms[i] = nms[i-1]
				Swaps++
				i--
				continue
			}
			break
		}
		nms[i] = tmp
		Swaps++
		sortedIndex++
	}
}

// https://en.wikipedia.org/wiki/Shellsort
// The fastest one seems to be floor(2.25k)
// https://stackoverflow.com/questions/2539545/fastest-gap-sequence-for-shell-sort
func (nms Nums) ShellSort() {
	ln := len(nms)
	incs := []int{1391376, 463792, 198768, 86961,
		33936, 13776, 4592, 1968, 861,
		336, 112, 48, 21, 7, 3, 1}
	for k := 0; k < len(incs); k++ {
		h := incs[k]
		for i := h; i < ln; i++ {
			for j := i; j >= h; j -= h {
				Comparisons++
				if nms[j] < nms[j-h] {
					nms[j], nms[j-h] = nms[j-h], nms[j]
					Swaps++
					continue
				}
				break
			}
		}
	}
}

func merge(a, aux Nums, l, m, r int) {
	Comparisons++
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
			Swaps++
			continue
		}
		if j >= r {
			a[k] = aux[i]
			i++
			Swaps++
			continue
		}
		Comparisons++
		if aux[i] > aux[j] {
			a[k] = aux[j]
			j++
			Swaps++
			continue
		}
		a[k] = aux[i]
		i++
		Swaps++
	}
}

func (nms Nums) MergeSort() {
	ln := len(nms)
	aux := make(Nums, ln)

	for m := 1; m < ln; m += m {
		for i := 0; i < ln-m; i += m + m {
			merge(nms, aux, i, i+m, int(math.Min(float64(i+m+m), float64(ln))))
		}
	}
}

func recursiveMerge(a, aux Nums, low, mid, high int) {
	i := low
	j := mid + 1
	Comparisons++
	if aux[j] > aux[mid] {
		copy(a[low:high+1], aux[low:high+1])
		return
	}
	for k := low; k <= high; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
			Swaps++
			continue
		}
		if j > high {
			a[k] = aux[i]
			i++
			Swaps++
			continue
		}
		Comparisons++
		if aux[i] > aux[j] {
			a[k] = aux[j]
			j++
			Swaps++
			continue
		}
		a[k] = aux[i]
		i++
		Swaps++
	}
}

func recursiveMergeSort(a, aux Nums, low, high int) {
	if high <= low {
		return
	}
	mid := low + (high-low)/2
	recursiveMergeSort(aux, a, low, mid)
	recursiveMergeSort(aux, a, mid+1, high)
	recursiveMerge(a, aux, low, mid, high)
}

func (nms Nums) RecursiveMergeSort() {
	aux := slices.Clone(nms)
	recursiveMergeSort(nms, aux, 0, len(nms)-1)
}

func (nms Nums) partition(l, r int) int {
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

func (nms Nums) quickSort(l, r int) {
	if r <= l {
		return
	}
	m := nms.partition(l, r)
	nms.quickSort(l, m-1)
	nms.quickSort(m+1, r)
}

func (nms Nums) QSort() {
	rand.Shuffle(len(nms), nms.swap)
	nms.quickSort(0, len(nms)-1)
}

func (nms Nums) less(i, j int) bool {
	Comparisons++
	return nms[i] < nms[j]
}

func (nms Nums) swap(i, j int) {
	Swaps++
	nms[i], nms[j] = nms[j], nms[i]
}
