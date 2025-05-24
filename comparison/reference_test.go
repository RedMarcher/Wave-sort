package algorithms

import (
	"math/rand/v2"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_insertSort(t *testing.T) {
	nms := Nums{3, 1, 4, 1, 5, 9, 2, 6}
	nms.InsertSort()
	want := Nums{1, 1, 2, 3, 4, 5, 6, 9}
	if diff := cmp.Diff(want, nms); diff != "" {
		t.Errorf("sorted list mismatch (-want +got):\n%s", diff)
	}
}

func Test_shellSort(t *testing.T) {
	nms := Nums{3, 1, 4, 1, 5, 9, 2, 6}
	nms.ShellSort()
	want := Nums{1, 1, 2, 3, 4, 5, 6, 9}
	if diff := cmp.Diff(want, nms); diff != "" {
		t.Errorf("sorted list mismatch (-want +got):\n%s", diff)
	}
}

func Test_mergeSort(t *testing.T) {
	nms := Nums{3, 1, 4, 1, 5, 9, 2, 6}
	nms.MergeSort()
	want := Nums{1, 1, 2, 3, 4, 5, 6, 9}
	if diff := cmp.Diff(want, nms); diff != "" {
		t.Errorf("sorted list mismatch (-want +got):\n%s", diff)
	}
}

func Test_recursiveMergeSort(t *testing.T) {
	nms := Nums{3, 1, 4, 1, 5, 9, 2, 6}
	nms.RecursiveMergeSort()
	want := Nums{1, 1, 2, 3, 4, 5, 6, 9}
	if diff := cmp.Diff(want, nms); diff != "" {
		t.Errorf("sorted list mismatch (-want +got):\n%s", diff)
	}
}

func Test_qSort(t *testing.T) {
	nms := Nums{3, 1, 4, 1, 5, 9, 2, 6}
	nms.QSort()
	want := Nums{1, 1, 2, 3, 4, 5, 6, 9}
	if diff := cmp.Diff(want, nms); diff != "" {
		t.Errorf("sorted list mismatch (-want +got):\n%s", diff)
	}
}
func Test_qSort_longSeq(t *testing.T) {
	seqLength := 1 << 20
	input := make(Nums, seqLength)
	for i := range input {
		input[i] = i
	}
	for range 10 {
		rand.Shuffle(seqLength, input.swap)
		input.QSort()
	}
	for i := range input {
		if input[i] != i {
			t.Errorf("qSort() = %v, want %v", input[i], i)
		}
	}
}
