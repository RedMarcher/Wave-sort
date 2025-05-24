package algorithm

import (
	"log"
	"math/rand/v2"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_insertSort(t *testing.T) {
	nms := nums{3, 1, 4, 1, 5, 9, 2, 6}
	nms.insertSort()
	want := nums{1, 1, 2, 3, 4, 5, 6, 9}
	if diff := cmp.Diff(want, nms); diff != "" {
		t.Errorf("sorted list mismatch (-want +got):\n%s", diff)
	}
}

func Test_shellSort(t *testing.T) {
	nms := nums{3, 1, 4, 1, 5, 9, 2, 6}
	nms.shellSort()
	want := nums{1, 1, 2, 3, 4, 5, 6, 9}
	if diff := cmp.Diff(want, nms); diff != "" {
		t.Errorf("sorted list mismatch (-want +got):\n%s", diff)
	}
}

func Test_mergeSort(t *testing.T) {
	nms := nums{3, 1, 4, 1, 5, 9, 2, 6}
	nms.mergeSort()
	want := nums{1, 1, 2, 3, 4, 5, 6, 9}
	if diff := cmp.Diff(want, nms); diff != "" {
		t.Errorf("sorted list mismatch (-want +got):\n%s", diff)
	}
}

func Test_recursiveMergeSort(t *testing.T) {
	nms := nums{3, 1, 4, 1, 5, 9, 2, 6}
	nms.recursiveMergeSort()
	want := nums{1, 1, 2, 3, 4, 5, 6, 9}
	if diff := cmp.Diff(want, nms); diff != "" {
		t.Errorf("sorted list mismatch (-want +got):\n%s", diff)
	}
}

func Test_qSort(t *testing.T) {
	nms := nums{3, 1, 4, 1, 5, 9, 2, 6}
	nms.qSort()
	want := nums{1, 1, 2, 3, 4, 5, 6, 9}
	if diff := cmp.Diff(want, nms); diff != "" {
		t.Errorf("sorted list mismatch (-want +got):\n%s", diff)
	}
}
func Test_qSort_longSeq(t *testing.T) {
	seqLength := 1 << 20
	input := make(nums, seqLength)
	for i := range input {
		input[i] = i
	}
	sumCmp, sumSwap := 0, 0
	for range 10 {
		rand.Shuffle(seqLength, input.swap)
		compare, swap = 0, 0
		input.qSort()
		sumCmp += compare
		sumSwap += swap
	}
	sumCmp /= 10
	sumSwap /= 10
	for i := range input {
		if input[i] != i {
			t.Errorf("qSort() = %v, want %v", input[i], i)
		}
	}
	log.Printf("\nnumber of comparisons: %d\n", sumCmp)
	log.Printf("number of swaps: %d\n", sumSwap)
}
