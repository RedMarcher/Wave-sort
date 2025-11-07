package wave_sort

import (
	"log"
	"math/rand/v2"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_partition(t *testing.T) {
	s := seq{3, 1, 4, 1, 5, 9, 2, 6}
	got := s.partition(0, len(s)-1, len(s)-1)
	want := 6
	if got != want {
		t.Errorf("partition() = %v, want %v", got, want)
	}
}

func Test_blockSwap_1(t *testing.T) {
	ts := []struct {
		name    string
		input   seq
		want    seq
		m, r, p int
	}{
		{
			input: seq{5, 6, 7, 1, 2, 3, 4},
			want:  seq{1, 2, 3, 4, 5, 6, 7},
			m:     0,
			r:     3,
			p:     6,
		},
		{"empty", seq{1}, seq{1}, 0, 0, 0},
	}
	for _, tc := range ts {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.input.blockSwap(tc.m, tc.r, tc.p)
			if diff := cmp.Diff(tc.want, tc.input); diff != "" {
				t.Errorf("blockSwap_1() failed with diff %v", diff)
			}
		})
	}
}

func Test_blockSwap(t *testing.T) {
	s := seq{3, 1, 4, 1, 5, 9, 2, 6}
	s.blockSwap(0, 4, 7)
	want := seq{5, 9, 2, 6, 1, 4, 1, 3}
	if diff := cmp.Diff(want, s); diff != "" {
		t.Errorf("blockSwap() = %v, want %v", s, want)
	}
}
func TestWaveSort(t *testing.T) {
	tests := []struct {
		input seq
		want  seq
	}{
		{seq{3, 1, 4, 1, 5, 9, 2, 6, 7}, seq{1, 1, 2, 3, 4, 5, 6, 7, 9}},
		{seq{1}, seq{1}},
		{seq{}, seq{}},
	}

	for _, test := range tests {
		test.input.WaveSort()
		if diff := cmp.Diff(test.want, test.input); diff != "" {
			t.Errorf("WaveSort() = %v, want %v", test.input, test.want)
		}
	}
}

func TestWaveSort_LongSeq(t *testing.T) {
	const testRun = 1
	seqLength := 1 << 20
	input := make(seq, seqLength)
	for i := range seqLength {
		input[i] = i
	}
	sumCmp, sumSwap, sumBs := 0, 0, 0
	for range testRun {
		rand.Shuffle(seqLength, input.swap)
		// input.reverse(0, seqLength-1)
		compare, swap, blockSwap = 0, 0, 0
		input.WaveSort()
		sumCmp += compare
		sumSwap += swap
		sumBs += blockSwap
	}
	sumCmp /= testRun
	sumSwap /= testRun
	sumBs /= testRun
	for i := range seqLength {
		if input[i] != i {
			t.Errorf("WaveSort() = %v, want %v", input[i], i)
		}
	}
	log.Printf("\nnumber of comparisons: %d\n", sumCmp)
	log.Printf("number of swaps: %d\n", sumSwap)
	log.Printf("number of block swaps: %d\n", sumBs)

}
