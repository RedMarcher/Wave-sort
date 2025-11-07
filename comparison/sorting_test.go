package algorithms

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

func Test_recursiveMergeSort_longSeq(t *testing.T) {
	seqLength := 1 << 20
	input := make(nums, seqLength)
	for i := range input {
		input[i] = i
	}
	sumCmp := 0
	runs := 100
	for range runs {
		rand.Shuffle(seqLength, input.swap)
		input.recursiveMergeSort()
	}
	sumCmp /= runs
	for i := range input {
		if input[i] != i {
			t.Errorf("recursive merge sort = %v, want %v", input[i], i)
		}
	}
	log.Printf("\nnumber of compare: %d\n", sumCmp)
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
	runs := 1
	for range runs {
		rand.Shuffle(seqLength, input.swap)
		input.qSort()
	}
	sumCmp /= runs
	sumSwap /= runs
	for i := range input {
		if input[i] != i {
			t.Errorf("qSort() = %v, want %v", input[i], i)
		}
	}
	log.Printf("\nnumber of comparisons: %d\n", sumCmp)
	log.Printf("number of swaps: %d\n", sumSwap)
}
