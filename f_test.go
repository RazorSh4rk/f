package f_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/RazorSh4rk/f"
)

func TestSimple(t *testing.T) {
	fn := f.From([]int{1, 2, 3})

	t.Log(fn)
}

func TestEq(t *testing.T) {
	fn0 := f.From([]string{"a", "b", "c"})
	fn1 := f.F[string]{
		Val: []string{"a", "b", "c"},
	}

	if !fn0.Is(fn1) {
		t.Fail()
	}
}

func TestForEach(t *testing.T) {
	fn := f.From([]int{1, 2, 3})
	fn.ForEach(func(i int) {
		t.Log(i * i)
	})
}

func TestForAll(t *testing.T) {
	fnT := f.From([]int{2, 4, 6})
	fnF := f.From([]int{1, 2, 3})

	predicate := func(i int) bool {
		return i%2 == 0
	}

	_true := fnT.ForAll(predicate)
	_false := fnF.ForAll(predicate)

	if _true != true || _false != false {
		t.Fail()
	}
}

func TestHas(t *testing.T) {
	fnT := f.From([]int{2, 4, 6})
	fnF := f.From([]int{1, 3, 5})

	predicate := func(i int) bool {
		return i%2 == 0
	}

	_true := fnT.Has(predicate)
	_false := fnF.Has(predicate)

	if _true != true || _false != false {
		t.Fail()
	}
}

func TestHead(t *testing.T) {
	fn := f.From([]int{1, 2, 3})
	head := fn.Head()

	if head != 1 {
		t.Fail()
	}
}

func TestLast(t *testing.T) {
	fn := f.From([]int{1, 2, 3})
	last := fn.Last()

	if last != 3 {
		t.Fail()
	}
}

func TestTail(t *testing.T) {
	fn := f.From([]int{1, 2, 3})
	tail := fn.Tail()

	fmt.Println(tail)

	if tail.Val[0] != 2 || tail.Val[1] != 3 {
		t.Fail()
	}
}

func TestReverse(t *testing.T) {
	fn := f.From([]int{1, 2, 3})
	r := fn.Reverse()

	if !r.Is(f.F[int]{
		Val: []int{3, 2, 1},
	}) {
		t.Fail()
	}
}

func TestFilter(t *testing.T) {
	fn := f.From([]int{1, 2, 3})
	res := fn.Filter(func(i int) bool {
		return i%2 == 1
	})

	if !res.Is(f.F[int]{
		Val: []int{1, 3},
	}) {
		t.Fail()
	}
}

func TestMap(t *testing.T) {
	fn := f.From([]int{1, 2, 3})
	res := f.Map(fn, func(i int) string {
		return fmt.Sprint(i)
	})

	if !res.Is(f.F[string]{
		Val: []string{"1", "2", "3"},
	}) {
		t.Fail()
	}
}

func TestFold(t *testing.T) {
	fn := f.From([]int{1, 2, 3})
	res := f.Fold(fn, 0, func(l int, r int) int {
		return l + r
	})

	if res != 6 {
		fmt.Println(res)
		t.Fail()
	}
}

func TestTakeWhile(t *testing.T) {
	fn := f.From([]int{1, 2, 3, 4, 5, 6})
	res := fn.TakeWhile(func(i int) bool {
		return i < 4
	})

	if !res.Is(f.F[int]{
		Val: []int{1, 2, 3},
	}) {
		t.Fail()
	}
}

func TestDropWhile(t *testing.T) {
	fn := f.From([]int{1, 2, 3, 4, 5, 6})
	res := fn.DropWhile(func(i int) bool {
		return i < 4
	})

	if !res.Is(f.F[int]{
		Val: []int{4, 5, 6},
	}) {
		t.Fail()
	}
}

func TestZip(t *testing.T) {
	fn0 := f.From([]int{1, 2, 3})
	fn1 := f.From([]int{4, 5, 6})
	fn2 := f.From([]int{1, 4, 2, 5, 3, 6})

	if !fn0.ZipWith(fn1).Is(fn2) {
		t.Fail()
	}

}

func TestChain(t *testing.T) {
	fn := f.Gen(func(i int) int64 {
		return int64(i * i)
	}, 100)

	res := f.Fold(
		f.Map(fn.Reverse().TakeWhile(func(i int64) bool {
			return i%2 == 0
		}).Filter(func(i int64) bool {
			return strings.Contains(fmt.Sprint(i), "3")
		}), func(i int64) int64 {
			return i % 10
		}), 5, func(l int64, r int64) int64 {
			return l + r
		},
	)

	if res != 57 {
		t.Fail()
	}
}
