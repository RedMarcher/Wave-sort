package main

import (
	"fmt"
	"math/rand"
	"time"
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
	start := time.Now()
	waveSequence.WaveSort()
	waveTime := time.Since(start)
	waveCmp, waveSwap, waveBs := wavesort.Comparisons, wavesort.Swaps, wavesort.BlockSwaps

	// InsertionSort
	insertInput := make([]int, N)
	copy(insertInput, input)
	algorithms.Comparisons, algorithms.Swaps = 0, 0
	start = time.Now()
	insertNums := algorithms.Nums(insertInput)
	insertNums.InsertSort()
	insertTime := time.Since(start)
	insertCmp, insertSwap := algorithms.Comparisons, algorithms.Swaps

	// ShellSort
	shellInput := make([]int, N)
	copy(shellInput, input)
	algorithms.Comparisons, algorithms.Swaps = 0, 0
	start = time.Now()
	shellNums := algorithms.Nums(shellInput)
	shellNums.ShellSort()
	shellTime := time.Since(start)
	shellCmp, shellSwap := algorithms.Comparisons, algorithms.Swaps

	// MergeSort
	mergeInput := make([]int, N)
	copy(mergeInput, input)
	algorithms.Comparisons, algorithms.Swaps = 0, 0
	start = time.Now()
	mergeNums := algorithms.Nums(mergeInput)
	mergeNums.MergeSort()
	mergeTime := time.Since(start)
	mergeCmp, mergeSwap := algorithms.Comparisons, algorithms.Swaps

	// QuickSort
	quickInput := make([]int, N)
	copy(quickInput, input)
	algorithms.Comparisons, algorithms.Swaps = 0, 0
	start = time.Now()
	quickNums := algorithms.Nums(quickInput)
	quickNums.QSort()
	quickTime := time.Since(start)
	quickCmp, quickSwap := algorithms.Comparisons, algorithms.Swaps

	fmt.Printf("WaveSort:\n")
	fmt.Printf("  Time: %v\n", waveTime)
	fmt.Printf("  Comparisons: %d\n", waveCmp)
	fmt.Printf("  Swaps: %d\n", waveSwap)
	fmt.Printf("  Block Swaps: %d\n", waveBs)

	fmt.Printf("\nInsertionSort:\n")
	fmt.Printf("  Time: %v\n", insertTime)
	fmt.Printf("  Comparisons: %d\n", insertCmp)
	fmt.Printf("  Swaps: %d\n", insertSwap)

	fmt.Printf("\nShellSort:\n")
	fmt.Printf("  Time: %v\n", shellTime)
	fmt.Printf("  Comparisons: %d\n", shellCmp)
	fmt.Printf("  Swaps: %d\n", shellSwap)

	fmt.Printf("\nMergeSort:\n")
	fmt.Printf("  Time: %v\n", mergeTime)
	fmt.Printf("  Comparisons: %d\n", mergeCmp)
	fmt.Printf("  Swaps: %d\n", mergeSwap)

	fmt.Printf("\nQuickSort:\n")
	fmt.Printf("  Time: %v\n", quickTime)
	fmt.Printf("  Comparisons: %d\n", quickCmp)
	fmt.Printf("  Swaps: %d\n", quickSwap)
}
