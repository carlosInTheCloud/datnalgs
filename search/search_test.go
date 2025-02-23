package search

import (
	"datnalgs/data"
	"testing"
)

func testSorter(v []int) {
	s := len(v) - 1
	swap := true
	for i := 0; i < s && swap; i++ {
		swap = false
		for j := 0; j < s-i; j++ {
			if v[j] > v[j+1] {
				v[j], v[j+1] = v[j+1], v[j]
				swap = true
			}
		}
	}
}

func TestLinearUnsortedSearch(t *testing.T) {
	v := []int{2, 4, 6, 8, 10, 12, 14, 18, 21, 23, 24}
	r := LinearUnsortedSearch(v, 18)
	if !r {
		t.Error("error while testing LinearUnsortedSearchTest for value 18. Expected true, but got false\n")
	}
	r = LinearUnsortedSearch(v, 22)
	if r {
		t.Error("error while testing LinearUnsortedSearchTest for value 22. Expected false, but got true\n")
	}
}

func TestLinearSortedSearch(t *testing.T) {
	v := []int{2, 4, 6, 8, 10, 12, 14, 18, 21, 23, 24}
	r := LinearSortedSearch(v, 18)
	if !r {
		t.Errorf("error while testing LinearSortedSearchTest for value 18. Expected true, but got false\n")
	}
	r = LinearSortedSearch(v, 22)
	if r {
		t.Errorf("error while testing LinearSortedSearchTest for value 22. Expected false, but got true\n")
	}
}

func TestBinarySearch(t *testing.T) {
	v := data.GetLv()
	testSorter(v)
	r := BynarySearch(v, 359)
	if !r {
		t.Errorf("invalid binarySearch result with value 1000. Expected true, but false")
	}
	r = BynarySearch(v, 1001)
	if r {
		t.Errorf("Invalid binarySearch result with value 1001. Expected false, but got true")
	}
}

func TestFibonacciSearch(t *testing.T) {
	v := data.GetLv()
	testSorter(v)
	r := FibonacciSearch(v, 359)
	if !r {
		t.Errorf("invalid fibonacciSearch result with value 1000. Expected true, but false")
	}
	r = FibonacciSearch(v, 1001)
	if r {
		t.Errorf("Invalid fibonacchiSearch result with value 1001. Expected false, but got true")
	}
}

func TestSumArrayItems(t *testing.T) {
	v := data.GetLv()
	r := SumArrayItems(v)
	ex := 500500
	if r != ex {
		t.Errorf("Invalid SumArrayItems result. Expected %v, but got %v", ex, r)
	}
}

func TestFindDups(t *testing.T) {
	v := []int{1, 3, 5, 7, 9, 7, 25, 21, 30}
	vr := FindDups(v)
	exl := 1
	if len(vr) != exl {
		t.Errorf("Invalid array size. Expected %v, but got %v.", exl, len(vr))
	}
	exv := 7
	if vr[0] != exv {
		t.Errorf("Invalid duplicate value. Expected %v, but got %v.", exv, vr[0])
	}
}

func TestFindDupsSorting(t *testing.T) {
	v := []int{1, 3, 5, 7, 9, 7, 25, 21, 30}
	vr := FindDupsSorting(v)
	exl := 1
	if len(vr) != exl {
		t.Errorf("Invalid array size. Expected %v, but got %v.", exl, len(vr))
	}
	exv := 7
	if vr[0] != exv {
		t.Errorf("Invalid duplicate value. Expected %v, but got %v.", exv, vr[0])
	}
}

func TestRemoveDups(t *testing.T) {
	v := []int{1, 3, 5, 3, 9, 1, 30}
	vr := RemoveDups(v)
	exl := 5
	if len(vr) != exl {
		t.Errorf("Invalid array size. Expected %v, but got %v.", exl, len(vr))
	}
	exv := 1
	if vr[0] != exv {
		t.Errorf("Invalid item found. Expected %v, but got %v.", exv, vr[0])
	}
	exv = 3
	if vr[1] != exv {
		t.Errorf("invalid item found. expected %v, but got %v.", exv, vr[0])
	}
	exv = 5
	if vr[2] != exv {
		t.Errorf("invalid item found. expected %v, but got %v.", exv, vr[0])
	}
	exv = 9
	if vr[3] != exv {
		t.Errorf("invalid item found. expected %v, but got %v.", exv, vr[0])
	}
	exv = 30
	if vr[4] != exv {
		t.Errorf("invalid item found. expected %v, but got %v.", exv, vr[0])
	}
}

func TestRemoveDupsMap(t *testing.T) {
	v := []int{1, 3, 5, 3, 9, 1, 30}
	vr := RemoveDupsMap(v)
	exl := 5
	if len(vr) != exl {
		t.Errorf("Invalid array size. Expected %v, but got %v.", exl, len(vr))
	}
	testSorter(vr)
	t.Log("vr:", vr)
	exv := 1
	if vr[0] != exv {
		t.Errorf("Invalid item found. Expected %v, but got %v.", exv, vr[0])
	}
	exv = 3
	if vr[1] != exv {
		t.Errorf("invalid item found. expected %v, but got %v.", exv, vr[0])
	}
	exv = 5
	if vr[2] != exv {
		t.Errorf("invalid item found. expected %v, but got %v.", exv, vr[0])
	}
	exv = 9
	if vr[3] != exv {
		t.Errorf("invalid item found. expected %v, but got %v.", exv, vr[0])
	}
	exv = 30
	if vr[4] != exv {
		t.Errorf("invalid item found. expected %v, but got %v.", exv, vr[0])
	}
}

func TestMissingNumber(t *testing.T) {
	v := []int{1, 10, 3, 9, 4, 8, 5, 7, 6}
	en := 2
	n, m := FindMissingNumber(v)
	if n != en {
		t.Errorf("invalid missing number. expected number %v, but got %v", en, n)
	}
	em := true
	if m != em {
		t.Errorf("invalid missing assertion. expected %v, but got %v", em, m)
	}
}

func TestMissingNumberMap(t *testing.T) {
	v := []int{1, 10, 3, 9, 4, 8, 5, 7, 6}
	en := 2
	n, m := FindMissingNumberMap(v)
	if n != en {
		t.Errorf("invalid missing number. expected number %v, but got %v", en, n)
	}
	em := true
	if m != em {
		t.Errorf("invalid missing assertion. expected %v, but got %v", em, m)
	}
}

func TestMissingNumberSum(t *testing.T) {
	v := []int{1, 10, 3, 9, 4, 8, 5, 7, 6}
	en := 2
	n, m := FindMissingNumberSum(v)
	if n != en {
		t.Errorf("invalid missing number. expected number %v, but got %v", en, n)
	}
	em := true
	if m != em {
		t.Errorf("invalid missing assertion. expected %v, but got %v", em, m)
	}
}
