# F()

## functional programming in go

```bash
go get github.com/RazorSh4rk/f
```

### Usage

Creating a F:

```golang
fn := f.From([]T{ ... })

fn := f.F[T]{
    Val: []T{ ... },
}

fn := f.Gen(func (index int) T {
    return index * T // or something else
})
```

Consumer function over an F:

```golang
// run a function over each element
fn.ForEach(func (element T) {
    // do something with element
})

// does every element in the F satisfy the predicate
fn.ForAll(func (element T) bool {
    // return some condition
})

// does any element in the F satisfy the predicate
fn.Has(func (element T) bool {
    // return some condition
})
```

Modifier function over an F (do not change type):

```golang
// Get first value
fn.Head()

// Get last value
fn.Last()

// Get all values except the first
fn.Tail()

// Reverse the order of the values
fn.Reverse()

// Take values from the front as long as they satisfy the predicate
fn.TakeWhile(func (element T) bool {
    // return some condition
})

// Drop values from the front as long as they satisfy the predicate
fn.DropWhile(func (element T) bool {
    // return some condition
})

// Filter values based on a predicate
fn.Filter(func (element T) bool {
    // return some condition
})

// Zip two Fs together
// They have to hold the same type
// Example: 
// zipping {1,2,3} and {4,5,6,7,8,9} will result in {1,4,2,5,3,6,7,8,9}
fn.Zip(f.F[T]{
    Val: []T{ ... },
})
```

Transformer function over an F (change type):

```golang
// Map values to a new type
f.Map(fn, func (element T) U {
    // return some new value of type U
})

// Fold (or reduce) the values to a single value
f.Fold(fn, startingValue U, func (acc U, element T) U {
    // return some new value of type U
})
```


```golang
// Example of method chaining
fn := f.Gen(func(i int) int64 {
	return int64(i * i)
}, 100)
// first 100 squares

res := f.Fold(
            // descending order
	f.Map(fn.Reverse().TakeWhile(func(i int64) bool {
        // only even numbers
		return i%2 == 0
	}).Filter(func(i int64) bool {
        // only numbers with a 3 in them
		return strings.Contains(fmt.Sprint(i), "3")
	}), func(i int64) int64 {
        // get the last digit
		return i % 10
        // sum them all up + 5
	}), 5, func(l int64, r int64) int64 {
		return l + r
	},
)
// 57
```