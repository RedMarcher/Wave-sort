package wavesort

import (
	"math/rand/v2"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_partition(t *testing.T) {
	s := Sequence{3, 1, 4, 1, 5, 9, 2, 6}
	got := s.partition(0, len(s)-1, len(s)-1)
	want := 6
	if got != want {
		t.Errorf("partition() = %v, want %v", got, want)
	}
}

func Test_blockSwap_1(t *testing.T) {
	ts := []struct {
		name    string
		input   Sequence
		want    Sequence
		m, r, p int
	}{
		{
			input: Sequence{5, 6, 7, 1, 2, 3, 4},
			want:  Sequence{1, 2, 3, 4, 5, 6, 7},
			m:     0,
			r:     3,
			p:     6,
		},
		{"empty", Sequence{1}, Sequence{1}, 0, 0, 0},
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
	s := Sequence{3, 1, 4, 1, 5, 9, 2, 6}
	s.blockSwap(0, 4, 7)
	want := Sequence{5, 9, 2, 6, 1, 4, 1, 3}
	if diff := cmp.Diff(want, s); diff != "" {
		t.Errorf("blockSwap() = %v, want %v", s, want)
	}
}
func TestWaveSort(t *testing.T) {
	tests := []struct {
		input Sequence
		want  Sequence
	}{
		{Sequence{3, 1, 4, 1, 5, 9, 2, 6, 7}, Sequence{1, 1, 2, 3, 4, 5, 6, 7, 9}},
		{Sequence{1}, Sequence{1}},
		{Sequence{}, Sequence{}},
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
	input := make(Sequence, seqLength)
	for i := range seqLength {
		input[i] = i
	}
	for range testRun {
		rand.Shuffle(seqLength, input.swap)
		// input.reverse(0, seqLength-1)
		Comparisons, Swaps, BlockSwaps = 0, 0, 0
		input.WaveSort()
	}
	for i := range seqLength {
		if input[i] != i {
			t.Errorf("WaveSort() = %v, want %v", input[i], i)
		}
	}
}

func TestWaveSort_e(t *testing.T) {
	tests := []struct {
		input Sequence
		want  Sequence
	}{
		{Sequence{3, 1, 4, 1, 5, 9, 2, 6, 7}, Sequence{1, 1, 2, 3, 4, 5, 6, 7, 9}},
		{Sequence{1}, Sequence{1}},
		{Sequence{}, Sequence{}},
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
	input := make(Sequence, seqLength)
	for i := range seqLength {
		input[i] = i
	}
	for range testRun {
		rand.Shuffle(seqLength, input.swap)
		// input.reverse(0, seqLength-1)
		Comparisons, Swaps, BlockSwaps = 0, 0, 0
		input.WaveSort_e()
	}
	for i := range seqLength {
		if input[i] != i {
			t.Errorf("WaveSort() = %v, want %v", input[i], i)
		}
	}
}

func Test_waveInsertSort(t *testing.T) {
	seqLength := 1 << 4
	input := make(Sequence, seqLength)
	for i := range seqLength {
		input[i] = i
	}
	rand.Shuffle(seqLength, input.swap)
	input.insertSort(0, seqLength-1)
	for i := range seqLength {
		if input[i] != i {
			t.Errorf("WaveSort() = %v, want %v", input[i], i)
		}
	}
}

func Test_preSorted(t *testing.T) {
	tcs := []struct {
		name   string
		input  Sequence
		output Sequence
		start  int
	}{
		{
			name:   "sorted",
			input:  Sequence{1, 2, 3, 4, 5},
			output: Sequence{1, 2, 3, 4, 5},
			start:  0,
		},
		{
			name:   "reversed",
			input:  Sequence{5, 4, 3, 2, 1},
			output: Sequence{1, 2, 3, 4, 5},
			start:  0,
		},
		{
			name:   "partial sorted",
			input:  Sequence{3, 1, 2, 4, 5},
			output: Sequence{3, 1, 2, 4, 5},
			start:  1,
		},
		{
			name:   "partial reversed",
			input:  Sequence{3, 1, 2, 5, 4},
			output: Sequence{3, 1, 2, 4, 5},
			start:  3,
		},
		{
			name:   "empty",
			input:  Sequence{},
			output: Sequence{},
			start:  0,
		},
		{
			name:   "one element",
			input:  Sequence{1},
			output: Sequence{1},
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
