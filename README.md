# F()

## functional programming in go

F is an unobtrusive functional programming library, that leaves your original data alone. Every transformation can be extracted separately and you can build transpormation chains, just like you would in something like Scala or OCaml.

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
    return T(index) // or something else
}, limit int)
// limit is the number of elements to generate
```

Some functions will use `Option`s:

```golang
var o f.Option[T]

// check if the option is empty
var ok bool = o.Ok()

// get the value of the option or error
res, err := o.Get()

// get the value of the option or a default value
res := o.GetOrElse(T)

```

Consumer function over a F:

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

// find the first element that satisfies the predicate
fn.Find(func (element T) bool {
    // return some condition
}
```

Modifier function over a F (no type change):

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

Transformer function over an F (type change):

```golang
// Map values to a new type
f.Map(fn, func (element T) U {
    // return some new value of type U
})

// Fold (or reduce) the values to a single value
f.Fold(fn, startingValue U, func (acc U, element T) U {
    // return some new value of type U
})

// Go down a dimension in a F's value
// for example from a [][]int{} to a []int{}
var flattened Option[T] = f.Flatten[T](fn)
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