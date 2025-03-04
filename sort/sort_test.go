package sort

import (
	"datnalgs/data"
	"fmt"
	"testing"
)

func testSortAsc(y []int) error {
	for i := 0; i < len(y)-2; i++ {
		if y[i] > y[i+1] {
			return fmt.Errorf("Array not sorted at index %v\n", i)
		}
	}
	return nil
}

func TestBubbleSort(t *testing.T) {
	v := data.GetLv()
	BubbleSort(v)
	t.Log("BubbleSort")
	e := testSortAsc(v)
	if e != nil {
		t.Error(e)
	}
}

func TestInsertionSort(t *testing.T) {
	v := data.GetLv()
	InsertionSort(v)
	t.Log("InsertionSort")
	e := testSortAsc(v)
	if e != nil {
		t.Error(e)
	}
}

func TestSelectionSort(t *testing.T) {
	v := data.GetLv()
	SelectionSort(v)
	t.Log("SelectionSort")
	e := testSortAsc(v)
	if e != nil {
		t.Error(e)
	}
}

func TestMergeSort(t *testing.T) {
	v := data.GetLv()
	MergeSort(v)
	t.Log("MergeSort")
	e := testSortAsc(v)
	if e != nil {
		t.Error(e)
	}
}

func TestQuickSort(t *testing.T) {
	v := data.GetLv()
	QuickSort(v)
	t.Log("QuickSort")
	e := testSortAsc(v)
	if e != nil {
		t.Error(e)
	}
}

func TestCountSort(t *testing.T) {
	v := data.GetLv()
	CountSort(v, 1, 1001)
	t.Log("CountSort")
	e := testSortAsc(v)
	if e != nil {
		t.Error(e)
	}
}

func TestRadixSort(t *testing.T) {
	v := data.GetLv()
	RadixSort(v)
	t.Log("RadixSort")
	e := testSortAsc(v)
	if e != nil {
		t.Error(e)
	}
}

func TestShellSort(t *testing.T) {
	v := data.GetLv()
	ShellSort(v)
	t.Log("ShellSort")
	e := testSortAsc(v)
	if e != nil {
		t.Error(e)
	}
}

func TestBucketSort(t *testing.T) {
	v := data.GetLv()
	BucketSort(v)
	t.Log("BucketSort")
	e := testSortAsc(v)
	if e != nil {
		t.Error(e)
	}
}

func TestPartition01Sort(t *testing.T) {
	v := []int{0, 1, 1, 0, 1, 0, 1, 1, 0, 0, 0, 1}
	Partition01Sort(v)
	t.Log("v:", v)
	e := testSortAsc(v)
	if e != nil {
		t.Error(e)
	}
}

func TestPartition012V2Sort(t *testing.T) {
	v := []int{0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1}
	Partition012Sort(v)
	t.Log("v:", v)
	e := testSortAsc(v)
	if e != nil {
		t.Error(e)
	}
}

// find the minimum number of swaps required to bring all elementes less than the given value
// together at the start of the array
// given the input [2,7, 5, 6, 1, 3, 4, 9, 10, 8
// and the key value 10, the expected minumum number of swap should be 2.

func TestMinSwaps(t *testing.T) {
	v := []int{2, 7, 5, 6, 1, 3, 4, 9, 10, 8}
	t.Log("v before:", v)
	ms := MinSwaps(v, 5)
	t.Log("v after:", v)
	a := 4
	if ms != a {
		t.Error(fmt.Errorf("Min swaps not match. Expected:%v, actual:%v", a, ms))
	}
	t.Log("Min swaps found:", ms)
}

func TestEvenOdd(t *testing.T) {
	v := []int{2, 7, 5, 6, 1, 3, 4, 9, 10, 8}
	EvenOdd(v)
	if v[0]%2 != 0 && v[len(v)-1] == 0 {
		t.Error("Array not sorted")
	}

	t.Log("v:", v)
}

func TestSortArrays(t *testing.T) {
	v1 := []int{1, 5, 9, 10, 15, 20}
	t.Log("v1 before:", v1)
	v2 := []int{2, 3, 8, 13}
	t.Log("v2 before:", v2)
	SortArrays(v1, v2)
	t.Log("v1:", v1)
	t.Log("v2:", v2)
	if v1[len(v1)-1] > v2[0] {
		t.Error("arrays not sorted")
	}
}

func TestUnionIntersect(t *testing.T) {
	v1 := []int{1, 11, 2, 3, 14, 5, 6, 8, 9}
	t.Log("v1", v1)
	v2 := []int{2, 4, 5, 12, 7, 8, 13, 10}
	t.Log("v2", v2)
	vu, vi := UnionIntersect(v1, v2)
	t.Log("vi:", vi)
	t.Log("vu:", vu)

	if len(vi) != 3 || vi[0] != 2 || vi[1] != 5 || vi[2] != 8 {
		t.Error(fmt.Printf("Intersect array is invalid"))
	}
	if len(vu) != 14 || vu[0] != 1 || vu[13] != 14 {
		t.Error(fmt.Printf("Union array is invalid"))
	}
}

func TestCheckReverse(t *testing.T) {
	v := []int{1, 2, 8, 5, 4, 3, 10, 11, 12, 18, 28}
	chxr := CheckReverse(v)

	if !chxr {
		t.Error("Expected true but got false")
	}
}

func TestCheckReverseShouldReturnFalse(t *testing.T) {
	v := []int{1, 2, 8, 5, 4, 10, 3, 11, 12, 18, 28}
	chxr := CheckReverse(v)

	if chxr {
		t.Error("Expected fale but got true")
	}
}
func TestRotateArray(t *testing.T) {
	v := []int{1, 2, 3, 4, 5, 6}
	RotateArray(v, 2)
	t.Log("v:", v)
	if len(v) != 6 {
		t.Error(fmt.Printf("lenght of v is invalid. Expected: 6; actual: %v", len(v)))
	}
	if v[0] != 3 {
		t.Error(fmt.Printf("Incorrect value at index 0. Expected: 3; actual: %v", v[0]))
	}
	if v[5] != 2 {
		t.Error(fmt.Printf("Incorrect value at index 5. Expected: 2; actual: %v", v[5]))
	}
}

func TestIndexZeroZort(t *testing.T) {
	v := []int{8, -1, 6, 1, 9, 3, 2, 7, 4, -1}
	IndexZeroSort(v)
	t.Log("v:", v)
	if len(v) != 10 {
		t.Error(fmt.Printf("length of v is invalid.Expected 10, but got %v", len(v)))
	}
	if v[0] != -1 {
		t.Error(fmt.Printf("Invalid sorting at index 0. Expceted -1, but got %v", v[0]))
	}
	if v[9] != 9 {
		t.Error(fmt.Printf("invalid sorting at index 9. Expected 9, but got %v", v[9]))
	}
	if v[5] != -1 {
		t.Error(fmt.Printf("Invalid sorting at index 5. Expected -1, but got %v", v[5]))
	}
}

func TestIndexOneToNSort(t *testing.T) {
	v := []int{8, 5, 6, 1, 9, 3, 2, 7, 4, 10}
	IndexOneToNSort(v)
	t.Log("v:", v)
	for i, x := range v {
		if x != i+1 {
			t.Error(fmt.Printf("Invalid sort at index %v. Expected %v, but got %v\n", i, i, x))
			break
		}
	}
}

func TestMinMaxSort(t *testing.T) {
	v := []int{1, 2, 3, 4, 5, 6, 7, 8}
	MaxMinSort(v)
	t.Log("v:", v)
	if v[0] != 8 {
		t.Error(fmt.Errorf("Invalid sort at index 0. Expected 8 but got %v\n", v[0]))
	}
	if v[1] != 1 {
		t.Error(fmt.Errorf("Invalid sort at index 1. Expected 1, but got %v\n", v[1]))
	}
	if v[6] != 5 {
		t.Error(fmt.Errorf("Invalid sort at index 6. Expected 5, but got %v\n", v[6]))
	}
	if v[7] != 4 {
		t.Error(fmt.Errorf("Invalid sort at index 7. Expected 4, but got %v\n", v[7]))
	}
}

func TestMinMaxONSort(t *testing.T) {
	v := []int{1, 2, 3, 4, 5, 6, 7, 8}
	MaxMinSortON(v)
	t.Log("v:", v)
	if v[0] != 8 {
		t.Error(fmt.Errorf("Invalid sort at index 0. Expected 8 but got %v\n", v[0]))
	}
	if v[1] != 1 {
		t.Error(fmt.Errorf("Invalid sort at index 1. Expected 1, but got %v\n", v[1]))
	}
	if v[6] != 5 {
		t.Error(fmt.Errorf("Invalid sort at index 6. Expected 5, but got %v\n", v[6]))
	}
	if v[7] != 4 {
		t.Error(fmt.Errorf("Invalid sort at index 7. Expected 4, but got %v\n", v[7]))
	}
}

func TestCircularSum(t *testing.T) {
	v := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	r := MaxCircularSum(v)
	if r != 290 {
		t.Error(fmt.Printf("Invalid max circular sum. Expected 290, but got %v\n", r))
	}
}
