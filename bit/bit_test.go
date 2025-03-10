package bit

import "testing"

func TestAnd(t *testing.T) {
	a := 9
	b := 9
	c := 2
	r := And(a, b)
	if r != 9 {
		t.Errorf("Error while testing AndEx: a {%v} and b {%v}. Expected 9, but got %v", a, b, r)
	}
	r = And(a, c)
	if r != 0 {
		t.Errorf("Error while testing AndEx: a {%v} and c {%v}. Expected 0, but got %v", a, c, r)
	}
}

func TestOr(t *testing.T) {
	a := 9
	b := 2
	c := 10
	r := Or(a, b)
	if r != 11 {
		t.Errorf("Error while testing OrEx: a {%v} and b {%v}. Expected 11, but got %v", a, b, r)
	}
	r = Or(a, c)
	if r != 11 {
		t.Errorf("Error while testing OrEx: a {%v} and c {%v}. Expected 11, but got %v", a, c, r)
	}
}

func TestXor(t *testing.T) {
	a := 9
	b := 2
	c := 10
	r := Xor(a, b)
	if r != 11 {
		t.Errorf("Error while testing XorEx: a {%v} and b {%v}. Expected 11, but got %v", a, b, r)
	}
	r = Xor(a, c)
	if r != 3 {
		t.Errorf("Error while testing XorEx: a {%v} and c {%v}. Expected 3, but got %v", a, c, r)
	}
}

func TestDivideBy2(t *testing.T) {
	v := 512
	hv := 512 / 2
	r := DivideBy2(v)
	if r != hv {
		t.Errorf("Error while dividing by 2. Expected %v, but got %v\n", hv, r)
	}
}

func TestMultiplyBy8(t *testing.T) {
	v := 1234567
	ocv := v * 8
	r := MultiplyBy8(v)
	if r != ocv {
		t.Errorf("Error while multiplying by 8. Expected %v, but got %v\n", ocv, r)
	}
}

func TestOnceComplement(t *testing.T) {
	a := 10
	r := OnceComplement(a)
	if r != -11 {
		t.Errorf("Error while finding the Once complement of 10. Expected -11 but got %v", r)
	}
}

func TestTwoComplement(t *testing.T) {
	a := 11
	r := TwoComplement(a)
	if r != -a {
		t.Errorf("Error while finding the Once complement of 10. Expected -11 but got %v", r)
	}
}

func TestKthBitCheck(t *testing.T) {
	a := 589
	k := 7
	if !KthBitCheck(a, k) {
		t.Errorf("Error while checking if the Kth {%v} bit of %v is set. Expected true, but got false", k, a)
	}
	k = 6
	if KthBitCheck(a, k) {
		t.Errorf("Error while checking if the Kth {%v} bit of %v is set. Expected false, but got true", k, a)
	}
}

func TestKthbitset(t *testing.T) {
	a := 589
	k := 2
	ex := 591
	r := KthBitSet(a, k)
	if r != ex {
		t.Errorf("error while setting the kth {%v} bit of %v (not set). expected %v, but got %v", k, a, ex, r)
	}
	k = 7
	r = KthBitSet(a, k)
	if r != a {
		t.Errorf("error while setting the kth {%v} bit of %v (already set). expected %v, but got %v", k, a, a, r)
	}
}

func TestKthBitReset(t *testing.T) {
	a := 5
	k := 3
	ex := 1
	r := KthBitReset(a, k)
	if r != ex {
		t.Errorf("error while resetting the kth {%v} bit of %v. expected %v, but got %v", k, a, ex, r)
	}
	k = 1
	ex = 4
	r = KthBitReset(a, k)

	if r != ex {
		t.Errorf("error while resetting the kth {%v} bit of %v. expected %v, but got %v", k, a, ex, r)
	}
}

func TestIsPowerOf2(t *testing.T) {
	a := 16
	if !IsPowerOf2(a) {
		t.Errorf("Error while checking if %v is of the power of 2. Got false, but expected true", a)
	}
	a = 15
	if IsPowerOf2(a) {
		t.Errorf("Error while checking if %v is of the power of 2. Got true, but expected false", a)
	}
}

func TestBitCount(t *testing.T) {
	a := 5
	ex := 2
	r := CountBits(a)
	if r != ex {
		t.Errorf("Error while counting the bits set in %v. Expected %v, but got %v", a, ex, r)
	}
	a = 4
	ex = 1
	r = CountBits(a)
	if r != ex {
		t.Errorf("Error while counting the bits set in %v. Expected %v, but got %v", a, ex, r)
	}
}

func TestFindOddValue(t *testing.T) {
	v := []int{1, 2, 3, 1, 2, 3, 4, 5, 6, 7, 5, 6, 7}
	o := FindOddValue(v)
	ex := 4
	if o != ex {
		t.Logf("invalid odd value. Expected %v, but got %v", ex, o)
	}
}
