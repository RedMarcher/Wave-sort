package algorithm

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
			tc.input.blockSwap_e(tc.m, tc.r, tc.p)
			if diff := cmp.Diff(tc.want, tc.input); diff != "" {
				t.Errorf("blockSwap_1() failed with diff %v", diff)
			}
		})
	}
}

func Test_blockSwap(t *testing.T) {
	s := seq{3, 1, 4, 1, 5, 9, 2, 6}
	s.blockSwap(0, 4, 7)
	want := seq{5, 9, 2, 6, 3, 1, 4, 1}
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
		compare, swap, bs = 0, 0, 0
		input.WaveSort()
		sumCmp += compare
		sumSwap += swap
		sumBs += bs
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

func TestWaveSort_e(t *testing.T) {
	tests := []struct {
		input seq
		want  seq
	}{
		{seq{3, 1, 4, 1, 5, 9, 2, 6, 7}, seq{1, 1, 2, 3, 4, 5, 6, 7, 9}},
		{seq{1}, seq{1}},
		{seq{}, seq{}},
	}

	for _, test := range tests {
		test.input.WaveSort_e()
		if diff := cmp.Diff(test.want, test.input); diff != "" {
			t.Errorf("WaveSort() = %v, want %v", test.input, test.want)
		}
	}
}

func TestWaveSort_e_LongSeq(t *testing.T) {
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
		compare, swap, bs = 0, 0, 0
		input.WaveSort_e()
		sumCmp += compare
		sumSwap += swap
		sumBs += bs
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
func Test_waveInsertSort(t *testing.T) {
	seqLength := 1 << 4
	input := make(seq, seqLength)
	for i := range seqLength {
		input[i] = i
	}
	rand.Shuffle(seqLength, input.swap)
	compare, swap, bs = 0, 0, 0
	input.insertSort(0, seqLength-1)
	for i := range seqLength {
		if input[i] != i {
			t.Errorf("WaveSort() = %v, want %v", input[i], i)
		}
	}
	log.Printf("\nnumber of comparisons: %d\n", compare)
	log.Printf("number of swaps: %d\n", swap)
	log.Printf("number of block swaps: %d\n", bs)
}

func Test_preSorted(t *testing.T) {
	tcs := []struct {
		name   string
		input  seq
		output seq
		start  int
	}{
		{
			name:   "sorted",
			input:  seq{1, 2, 3, 4, 5},
			output: seq{1, 2, 3, 4, 5},
			start:  0,
		},
		{
			name:   "reversed",
			input:  seq{5, 4, 3, 2, 1},
			output: seq{1, 2, 3, 4, 5},
			start:  0,
		},
		{
			name:   "partial sorted",
			input:  seq{3, 1, 2, 4, 5},
			output: seq{3, 1, 2, 4, 5},
			start:  1,
		},
		{
			name:   "partial reversed",
			input:  seq{3, 1, 2, 5, 4},
			output: seq{3, 1, 2, 4, 5},
			start:  3,
		},
		{
			name:   "empty",
			input:  seq{},
			output: seq{},
			start:  0,
		},
		{
			name:   "one element",
			input:  seq{1},
			output: seq{1},
			start:  0,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			start := tc.input.preSorted(0, len(tc.input)-1)
			if start != tc.start {
				t.Errorf("preSorted() = %v, want %v", start, tc.start)
			}
			for i := range tc.input {
				if tc.input[i] != tc.output[i] {
					t.Errorf("preSorted() = %v, want %v", tc.input[i], tc.output[i])
				}
			}
		})
	}
}
