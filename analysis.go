package main

import (
	"fmt"
	"math/rand"
	algorithms "wavesort/comparison"
	"wavesort/wavesort"
)

func main() {
	const N = 1 << 10

	// Generate random input
	input := make([]int, N)
	for i := range input {
		input[i] = rand.Intn(N)
	}

	// WaveSort
	waveInput := make([]int, N)
	copy(waveInput, input)
	wavesort.Comparisons, wavesort.Swaps, wavesort.BlockSwaps = 0, 0, 0
	waveSequence := wavesort.Sequence(waveInput)
	waveSequence.WaveSort()
	waveCmp, waveSwap, waveBs := wavesort.Comparisons, wavesort.Swaps, wavesort.BlockSwaps

	// WaveSort w/ tradeoff
	waveInput_e := make([]int, N)
	copy(waveInput_e, input)
	wavesort.Comparisons, wavesort.Swaps, wavesort.BlockSwaps = 0, 0, 0
	waveSequence_e := wavesort.Sequence(waveInput)
	waveSequence_e.WaveSort_e()
	waveCmp_e, waveSwap_e, waveBs_e := wavesort.Comparisons, wavesort.Swaps, wavesort.BlockSwaps

	// InsertionSort
	insertInput := make([]int, N)
	copy(insertInput, input)
	algorithms.Comparisons, algorithms.Swaps = 0, 0
	insertNums := algorithms.Nums(insertInput)
	insertNums.InsertSort()
	insertCmp, insertSwap := algorithms.Comparisons, algorithms.Swaps

	// ShellSort
	shellInput := make([]int, N)
	copy(shellInput, input)
	algorithms.Comparisons, algorithms.Swaps = 0, 0
	shellNums := algorithms.Nums(shellInput)
	shellNums.ShellSort()
	shellCmp, shellSwap := algorithms.Comparisons, algorithms.Swaps

	// MergeSort
	mergeInput := make([]int, N)
	copy(mergeInput, input)
	algorithms.Comparisons, algorithms.Swaps = 0, 0
	mergeNums := algorithms.Nums(mergeInput)
	mergeNums.MergeSort()
	mergeCmp, mergeSwap := algorithms.Comparisons, algorithms.Swaps

	// QuickSort
	quickInput := make([]int, N)
	copy(quickInput, input)
	algorithms.Comparisons, algorithms.Swaps = 0, 0
	quickNums := algorithms.Nums(quickInput)
	quickNums.QSort()
	quickCmp, quickSwap := algorithms.Comparisons, algorithms.Swaps

	fmt.Printf("WaveSort:\n")
	fmt.Printf("  Comparisons: %d\n", waveCmp)
	fmt.Printf("  Swaps: %d\n", waveSwap)
	fmt.Printf("  Block Swaps: %d\n", waveBs)

	fmt.Printf("\nWaveSort w\\ tradeoff:\n")
	fmt.Printf("  Comparisons: %d\n", waveCmp_e)
	fmt.Printf("  Swaps: %d\n", waveSwap_e)
	fmt.Printf("  Block Swaps: %d\n", waveBs_e)

	fmt.Printf("\nInsertionSort:\n")
	fmt.Printf("  Comparisons: %d\n", insertCmp)
	fmt.Printf("  Swaps: %d\n", insertSwap)

	fmt.Printf("\nShellSort:\n")
	fmt.Printf("  Comparisons: %d\n", shellCmp)
	fmt.Printf("  Swaps: %d\n", shellSwap)

	fmt.Printf("\nMergeSort:\n")
	fmt.Printf("  Comparisons: %d\n", mergeCmp)
	fmt.Printf("  Swaps: %d\n", mergeSwap)

	fmt.Printf("\nQuickSort:\n")
	fmt.Printf("  Comparisons: %d\n", quickCmp)
	fmt.Printf("  Swaps: %d\n", quickSwap)
}
