package templates

import (
	"math/rand"
	"testing"
)

func sumRange(a []int, l, r int) int {
	sum := 0
	for i := l; i <= r; i++ {
		sum += a[i]
	}
	return sum
}

func addRange(a []int, l, r, delta int) {
	for i := l; i <= r; i++ {
		a[i] += delta
	}
}

func TestSegTree_BasicCases(t *testing.T) {
	nums := []int{1, 5, 4, 2, 3}
	st := NewSegTree(append([]int(nil), nums...))
	raw := append([]int(nil), nums...)

	if got, want := st.Query(1, 3), sumRange(raw, 1, 3); got != want {
		t.Fatalf("initial query mismatch: got=%d want=%d", got, want)
	}

	st.Update(1, 2, 2)
	addRange(raw, 1, 2, 2)
	if got, want := st.Query(2, 3), sumRange(raw, 2, 3); got != want {
		t.Fatalf("after update #1 query mismatch: got=%d want=%d", got, want)
	}

	st.Update(0, 4, 1)
	addRange(raw, 0, 4, 1)
	if got, want := st.Query(0, 3), sumRange(raw, 0, 3); got != want {
		t.Fatalf("after update #2 query mismatch: got=%d want=%d", got, want)
	}
}

func TestSegTree_PointAndWholeRange(t *testing.T) {
	nums := []int{7, -3, 8, 0, 5, -2}
	st := NewSegTree(append([]int(nil), nums...))
	raw := append([]int(nil), nums...)

	for i := range nums {
		if got, want := st.Query(i, i), raw[i]; got != want {
			t.Fatalf("point query mismatch at i=%d: got=%d want=%d", i, got, want)
		}
	}

	st.Update(2, 2, 10)
	addRange(raw, 2, 2, 10)
	if got, want := st.Query(0, len(nums)-1), sumRange(raw, 0, len(raw)-1); got != want {
		t.Fatalf("whole range query mismatch: got=%d want=%d", got, want)
	}
}

func TestSegTree_RandomAgainstBruteForce(t *testing.T) {
	const (
		n          = 40
		operations = 3000
		seed       = 20260213
	)

	rng := rand.New(rand.NewSource(seed))
	raw := make([]int, n)
	for i := range raw {
		raw[i] = rng.Intn(41) - 20
	}
	st := NewSegTree(append([]int(nil), raw...))

	for step := range operations {
		l := rng.Intn(n)
		r := rng.Intn(n)
		if l > r {
			l, r = r, l
		}

		if rng.Intn(2) == 0 {
			delta := rng.Intn(31) - 15
			st.Update(l, r, delta)
			addRange(raw, l, r, delta)
			continue
		}

		got := st.Query(l, r)
		want := sumRange(raw, l, r)
		if got != want {
			t.Fatalf("step=%d query[%d,%d] mismatch: got=%d want=%d", step, l, r, got, want)
		}
	}
}
