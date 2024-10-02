---
theme: ./
titleTemplate: "%s - Shodo"
author: Manon Carbonnel
layout: cover
---

<div class="subtitle">Shodo</div>

# Go basics for algorithmic foundations

---

```yaml
layout: intro
```

**I'm Alex(andre)**

- Backend developer
- Worked one year for OVH
- Loves to write articles about IT

---

```yaml
layout: center
```

<strong>How about you?</strong>

---

```yaml
layout: three-cols-header
```

### What's the plan?

::left::

#### Learn Go basics 

`int`, `for`, `slice`, `func`…

::center::

#### Revisit theoretical concepts in Go

Control flows and algorithms.

::right::

#### Pratice a lot

1/3 theory, 2/3 exercises.

---

```yaml
layout: default
```

# Slides

https://amoevi.github.io/go-algorithms-basics-course

# Source code

https://github.com/amoevi/go-algorithms-basics-course

---

```yaml
layout: section
index: 1
```

## What is Go?

---

```yaml
layout: three-cols-header
```

### A little bit of history

::left::
2007

Project began at Google.


::center::
2009

Publicly announced.

::right::
2012

Version 1.0 released.

---

```yaml
layout: three-cols-header
```

### Why Go?

::left::
**Simplicity**

Clean and straightforward syntax, easy to learn.


::center::
**Performance**

Compiled language with efficient execution.

::right::
**Scalability**

Ideal for building large-scale, distributed systems.


---

```yaml
layout: full
```

# 📝 Exercise : verify Go Installation

- Open a terminal and run `go version` command. If Go is properly installed, the result should look like this.

```console
$ go version
go version go1.23.1 linux/amd64
$
```

🚨 Go is not installed? Download Go from the official website and follow installation instructions : https://go.dev/doc/install

---

```yaml
layout: full
```

# 📝 Exercise : download the source code

- Clone the repository to download the project source.
- Navigate to the project folder.
- Check the project structure by listing its contents.

```console
$ git clone https://github.com/amoevi/go-algorithms-basics-course.git
Cloning into 'go-algorithms-basics-course'...
remote: Enumerating objects: 288, done.
remote: Counting objects: 100% (288/288), done.
remote: Compressing objects: 100% (206/206), done.
remote: Total 288 (delta 68), reused 270 (delta 53), pack-reused 0 (from 0)
Receiving objects: 100% (288/288), done.
$ cd go-algorithms-basics-course
$ ls
docs  exercises  lab  README.md
$
```

🚨 Git is not installed? Download the project as a ZIP file from GitHub : https://github.com/amoevi/go-algorithms-basics-course

---

```yaml
layout: full
```

# Go commands

- `go run` compiles Go code on the fly and immediately runs it.
```console
$ ls
hello-world.go
$ go run hello-world.go
hello, world
```

- `go build` compiles source code into an executable file.
```console
$ ls
hello-world.go
$ go build hello-world.go
$ ls
hello-world.go hello-world
$ ./hello-world
hello, world
```

---

```yaml
layout: full
```

# Go documentation and resources

|                                        |                                                                                                        |
| -------------------------------------- | ------------------------------------------------------------------------------------------------------ |
| https://go.dev/doc                     | The official documentation for Go.                                                                     |
| https://go.dev/tour                    | An interactive tour that introduces the basics.                                                        |
| https://go.dev/doc/effective_go        | A guide on best practices and idiomatic ways to write Go code.                                         |
| https://go.dev/play                    | An online tool to write, run, and share Go code.                                                       |
| https://go.dev/ref/spec                | The official specification for the language.                                                           |


---

```yaml
layout: full
```

# 📝 Exercise : run your first program

- Navigate to the `ex-01` folder and run the `hello-world.go` program with the following using `go run`.

```console
$ cd exercises
$ cd ex-01
$ go run hello-world.go
hello, world
$
```

---

```yaml
layout: section
index: 2
```

## First steps

---

```yaml
layout: full
```

# Basic types

|             |                                  |
| ----------- | -------------------------------  |
| `bool`      | `true` / `false`                 |
| `int`       | …, `-1`, `-2`, `0`, `1`, `2`, …  |
| `float64`   | `19.99`, `7.0`                   |
| `string`    | `"hello, world"`                 |
| `rune`      | `'h'`, `'e'`, `','`, `'π'`, `'€'`|

---

```yaml
layout: full
```

# Basic operations

### Boolean operations

AND (`&&`), OR (`||`), NOT (`!`)

### Arithmetic operations

Addition (`+`), Subtraction (`-`), Multiplication (`*`), Division (`/`), Modulus (`%`)

###  Comparison operations

Equal (`==`), Not equal (`!=`), Less than (`<`), Greater than (`>`), Less than or equal (`<=`), Greater than or equal (`>=`)

### String operations

Concatenation (`+`)

---

```yaml
layout: full
```

# 📝 Exercise : play with basic types

- Navigate to the `ex-02` folder to edit the `basics.go` file.
- Add the missing code in place of the TODO.

```console
$ cd exercises
$ cd ex-02
$ go run basics.go
11 plus 4 IS 15
22.0 divided 7.0 IS 3.142857142857143
true AND false IS false
NOT false IS true
"go" plus "lang" IS golang
-------------------------------------------
256 times 128 IS TODO
28.0 divided by 7.0 IS TODO
28 divided by 16 IS TODO
28 mod 16 IS TODO
"basics" plus "." plus "go" IS TODO
...
```

---

```yaml
layout: full
```

# Variables

- Multiple ways to declare a variable.

```go {all|6|7|8|all}
package main

import "fmt"

func main() {
	var a int = -6
	var b int
	c := 13
	fmt.Println(a, b, c) // -6 0 13
}
```

---

```yaml
layout: full
```

# Variables

- Variables without an initial value are automatically set to their zero value.
- Variables in Go can be reassigned with new values after their initial declaration

```go {all|6-9|11-14|all}
package main

import "fmt"

func main() {
	var d int       // 0
	var e float64   // 0.0
	var f string    // ""
	var g bool      // false

	i := 2          // i == 2
	i = 3           // i == 3
	i = i + 1       // i == 4
	i = i * 2       // i == 8
}

```

---

```yaml
layout: three-cols-header
```

### `:=`, `=` and `==` are different 

::left::

#### `:=`

Declare a new variable
```go
x := 42
```

::center::

#### `=`

Assign a new value to a variable.
```go
x = x * 2
```

::right::

#### `==`

Compare two values for equality.
```go
11+4 == 15
```

---

```yaml
layout: full
```

# 📝 Exercise : create and use variables

- Navigate to the `ex-03` folder and edit the `variables.go` file.
- Assign values to `quotient` and `remainder` using appropriate basic operations.
- Replace `TODO` with logic to check if `a` and `b` are even.

```console
$ cd exercises
$ cd ex-03
$ go run variables.go
a = 28 b = 3
Quotient: 0
Remainder: 0
Is 28 even ? TODO
Is 3 even ? TODO
$
```
---

```yaml
layout: section
index: 3
```

## Control structures

---

```yaml
layout: full
```

# `if` statements

- If the condition is `true`, the code inside the `if` block is executed.
- `else if` is used to test multiple conditions in sequence.
- `else` keyword provides an alternative block of code if the conditions are false.

```go {all|1-2|3-6|7-9|11-13|all}
if (28 % 2) == 0 {
	fmt.Println("28 is an even number")
} else if (28 > 3) && (28-3 >= 0) {
	fmt.Println("28 is really greater than 3")
} else if 3 == 3 {
	fmt.Println("Three is the magic number")
} else {
	fmt.Println("Well, I guess none of the above is true")
}

if result := someFunction("INFO", 1726667324, "192.168.1.1"); result > 5 {
	fmt.Println("The computed result was superior to 5: ", result)
}
```

---

```yaml
layout: full
```

# `switch` statements

- `switch` is used to select one of many blocks of code to be executed.
- Multiple expressions can be evaluated in the same `case` by separating them with commas.

```go {all|1-3|4-9|10-11|all}
food := "apple"

switch food {
case "apple":
	fmt.Println(food, "is a fruit")
case "carrot", "spinach":
	fmt.Println(food, "is a vegetable")
case "chicken", "beef", "pork":
	fmt.Println(food, "is a meat")
default:
	fmt.Println("Unknown food type")
}
```

---

```yaml
layout: full
```

# `for` statements

- Unlike other languages, Go has only one looping construct, the `for` loop.

```go {all|1-3|5-9|all}
for i := 0; i < 10; i++ {
	fmt.Println(i)
}

j := 1
for j < 1000000 {
	fmt.Println(j)
	j *= 2
}
```

---

```yaml
layout: full
```

# `for` statements

- `break` terminates the loop immediately.
- `continue` skips the remaining code in the loop and moves to the next iteration.

```go {all|5-7|9-11|all}
i := 0
for {
	i++

	if i > 9 {
		break
	}

	if i%2 == 0 {
		continue
	}

	fmt.Println(i)
}
// 1
// 3
// 5
// 7
// 9
```

---

```yaml
layout: full
```

# `range` statements
- `range` iterates over elements of a collection.
- It returns the index and the value itself.
- Use `_` to discard unused values.

```go {all|1-4|6-11|13-16|all}
floating := []float64{ 3.14159, 2.71828, 1.61803 }
for index, value := range floating {
	fmt.Println("This is the element", index, "in my slice:", value)
}

for _, char := range "Hello World" {
	if char == ' ' {
		fmt.Println("Found a space!")
		break
	}
}

// Range over integers since Go 1.22 (February 2024)
for number := range 10 {
	fmt.Println(number)
}
```

---

```yaml
layout: full
```

# 📝 Exercise : work with control structures

- Navigate to the `ex-04` folder and edit the `statements.go` file.
- Replace `TODO` with logic that sets `isPrime` based on whether the number is prime.

```console
$ cd exercises
$ cd ex-04
$ go run statements.go
1948764641 is not a prime number
$
```

---

```yaml
layout: section
index: 4
```

## Slices

---

```yaml
layout: full
```

# Arrays or slices?
- You may see "arrays" and "slices" in Go documentation.
- Arrays and slices are both sequences of elements of a single type.
- Arrays have a fixed length, not flexible.
- Slices can grow or shrink, easier to work with.
- In this course, when we refer to arrays or lists, we’ll mostly mean slices.

---

```yaml
layout: full
```

# Creating slices

```go {all|1-2|4-5|7-8|all}
// Declare an empty slice of integers
var sli []int

// Declare a slice of integers with initial values
sli := []int{1, 2, 3} 

// Create a slice of integers with 5 elements
sli := make([]int, 5)
```
---

```yaml
layout: full
```

# Accessing and modifying slices

```go {all|3-4|6-7|9-11|13-14|all}
sli := []string{"perl", "go", "java", "python", "c"}

// Access an element
lang := sli[0]

// Accessing an out-of-bounds index will cause a runtime error
// oops := sli[5]

// Edit an element
sli[1] = "golang"
sli[4] += "++" 

// Swap elements
sli[2], sli[3] = sli[3], sli[2]
```
---

```yaml
layout: full
```
# Slicing slices

- Slicing creates a new slice from an existing one.
- `slice[start:end]` returns a sub-slice from start (inclusive) to end (exclusive).
- Modifying a sub-slice may also modify the original slice.

```go {all|3-5|7-11|all}
sli := []string{"perl", "golang", "python", "java", "c++"}

sli1 := sli[1:3] // ["golang", "python"]
sli2 := sli[2:] // ["python", "java", "c++"]
sli3 := sli[:3] // ["perl", "golang", "python"]

sli2[0] = "ruby"

fmt.Println(sli2) // ["ruby", "java", "c++"]
fmt.Println(sli) // ["perl", "golang", "ruby", "java", "c++"]
fmt.Println(sli3) // ["perl", "golang", "ruby"]
```

---

```yaml
layout: full
```
# Common slices operations

```go {all|3-4|6-8|9-11|13-21|all}
sli := []float64{ 3.14159, 2.71828, 1.61803 }

// len() returns the number of elements in the slice
fmt.Println(len(sli))

// append() add elements to a slice
sli = append(sli, 1.41421)

// copy() copies elements from one slice to another
dst := make([]string, len(sli))
copy(dst, sli)

// for statements to iterate over a slice
for i := 0; i < len(sli); i++ {
	value := sli[i]
	fmt.Println("This is the element", i, "in my slice:", value)
}

for i, value := range sli {
	fmt.Println("This is the element", i, "in my slice:", value)
}
```

---

```yaml
layout: full
```

# 📝 Exercise : compare two slices

- Navigate to the `ex-05` folder and edit the `slices.go` file.
- Replace `TODO` with logic that sets `areSlicesEqual` based on whether the slices are equal.

```console
$ cd exercises
$ cd ex-05
$ go run slices.go
false
false
false
false
$
```

---

```yaml
layout: section
index: 5
```

## Maps

---

```yaml
layout: full
```

# Maps
- Maps store data in key-value pairs.
- Each key in a map must be unique.
- Values are accessed by using the key.

```go {all|1-2|4-5|7-9|11-13|all}
// Declare a map with initial values
m := map[string]int{"foo": 1, "bar": 2}

// Create an empty map
n := make(map[string]int)

// Add elements in the maps
n["k1"] = 7
n["k2"] = 13

// Update an element
// The previous value of "k1" (7) is overwritten with 9
n["k1"] = 9
```

---

```yaml
layout: full
```

# Common maps operations

```go {all|3-8|10-11|13-14|16-17|19-22|all}
m := map[string]float64{"pi": 3.14159, "phi": 1.61803, "ln10":  2.30258}

// Access an element
v1 := m["ln10"]
v2, exists := m["sqrt2"]
if exists {
	fmt.Println("sqrt2 is", v2)
}    

// Delete an element
delete(m, "phi")

// Delete ALL elements
clear(m)

// Get the length of the map
length := len(m)

// Iterate over the map
for key, value := range m {
    fmt.Println(key, ":", value)
}
```

---

```yaml
layout: full
```

# 📝 Exercise : count with maps

- Navigate to the `ex-06` folder and edit the `maps.go` file.
- Replace `TODO` with the proper function implementation.

```console
$ cd exercises
$ cd ex-06
$ go run maps.go
The most frequent digit is 0 with 0 occurrences
$
```

---

```yaml
layout: section
index: 6
```

## Functions
---

```yaml
layout: full
```
# Functions
- A function is a block of code that can be reused and organized to perform tasks.

```go
func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func calculateArea(width, height float64) float64 {
	return width * height
}

func greetBasedOnTime(hour int) string {
	switch {
	case hour < 12:
		return "Good morning"
	case hour < 18:
		return "Good afternoon"
	default:
		return "Good evening"
	}
}
```

---

```yaml
layout: full
```
#  Declaring and calling a function

- Functions are defined with a **signature** that specifies what kind of data the function accepts and what it returns.

```go {all|1|all}
func add(a int, b int) int {
	return a + b
}
```

- To execute a function, you need to call it by passing arguments that match the parameters defined in the signature.

```go {all|1|all}
result := add(1, 5)
fmt.Println("Result:", result)
```

---

```yaml
layout: full
```

# Parameters
- Functions can accept zero, one, or multiple parameters.

```go {all|1-4|6-10|12-16|all}
func sayHi() {
	fmt.Println("Hello, World!")
}
sayHi()

func square(n int) int {
	return n * n
}
result := square(5)
fmt.Println("Square of 5:", result)

func checkStringLength(s string, length int) bool {
	return len(s) == length
}
output := checkStringLength("hello", 5)
fmt.Println("Does it match?", output)
```

---

```yaml
layout: full
```
# Return values
- Functions may return zero, one, or multiple values.
- Use `_` to ignore unused return values.
```go {all|1-4|6-10|12-16|all}
func greet(name string) {
	fmt.Println("Hello", name)
}
greet("Alex")

func isEven(n int) bool {
	return n%2 == 0
}
result := isEven(5)
fmt.Println("Is 101 event?", result)

func divide(a int, b int) (int, int) {
	return a / b, a % b
}
quotient, remainder := divide(9, 4)
_, parity := divide(9, 4)
```

---

```yaml
layout: full
```
#  Recursion
- Base case prevents infinite recursion.
- Recursive case calls the function with a smaller argument, breaking the problem down.
- Recursion is useful for divide and conquer algorithms.

```go {all|2-5|6-7|all}
func factorial(n int) int {
    // Base case: when n is 0 or 1.
    if n == 0 || n == 1 {
        return 1
    }
    // Recursive case: n * factorial(n-1).
    return n * factorial(n-1)
}

fmt.Println(factorial(5))
```

---

```yaml
layout: full
```

# 📝 Exercise : declare and call functions

- Navigate to the `ex-07` folder and edit the `functions.go` file.
- Replace `TODO` with the proper function implementation.

```console
$ cd exercises
$ cd ex-07
$ go run functions.go

$
```

---

```yaml
layout: section
index: 7
```

## Packages

---

```yaml
layout: full
```
# Packages

- Packages are a way to organize code in Go.
- Packages allow code reuse across multiple files.
- Each file belongs to a package.
- The `main` package is where execution starts, and the `main()` function must be present.

```go {all|1|all}
package main

import "fmt"

func main() {
	fmt.Println("hello, world")
}
```

---

```yaml
layout: full
```
# `import` statement

- `import` is used to include code from another package.
- You can import multiple packages by using parentheses.

```go {all|3-7|all}
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("The square root of 25 is:", math.Sqrt(25))
	fmt.Println("Current time is:", time.Now())
}
```

---

```yaml
layout: two-cols-header
```

- Uppercase names are exported and accessible across packages.
- Lowercase keeps names private to the package.

::left::

```go
// mathutils.go
package mathutils

func Add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}
```

::right::

```go
// main.go
package main

import (
	"fmt"
	"path/to/your/project/mathutils"
)

func main() {
	sum := mathutils.Add(3, 5)
	fmt.Println("Sum:", sum)

	// Will cause a compilation error
	// difference := mathutils.subtract(5, 3)
}
```

---

```yaml
layout: section
index: 8
```

## Standard library

---

```yaml
layout: full
```
# `fmt`

- `fmt` package provides formatted I/O operations.

```go {all|1-9|11-15|17-21|all}
// Prints using the default format to the standard output.
fmt.Print("hello, world")

// Prints using the default format and adds a new line.
fmt.Println("hello, world")

// Prints custom formatted output.
cost, currency := 9.99, "EUR"
fmt.Printf("This costs %.2f %s.", cost, currency)

// Stores output in a string.
n := 22.0 / 7.0
message1 := fmt.Sprint("22.0 / 7.0 =", n, ".")
message2 := fmt.Sprintln("22.0 / 7.0 =", n, ".")
message3 := fmt.Sprintf("22.0 / 7.0 = %.2f.", n)

// Reads formatted input using Scanf.
var year int
fmt.Print("Enter the year: ")
fmt.Scanf("%d", &year)
fmt.Printf("We are in the year %d!\n", year)
```

---

```yaml
layout: full
```
# `math`

- `math` package provides basic constants and mathematical functions.

```go
// Constants
fmt.Println(math.E, math.Phi, math.Ln10)
fmt.Println(math.MinInt, math.MaxInt)

// Functions
n := 22.0
fmt.Println(math.Abs(n), math.Cos(n), math.Exp(n), math.IsNaN(n), math.Round(n), math.Sqrt(n))

m := 7.0
fmt.Println(math.Max(n, m), math.Min(n, m), math.Mod(n, m), math.Pow(n, m))
```
---

```yaml
layout: full
```
# `strings`

- `strings` package provides functions to manipulate strings.

```go
strings.ToUpper("hello") // "HELLO"
strings.ToLower("GOODBYE") // goodbye

strings.Contains("golang", "go") // true
strings.ContainsRune("golang", 'g') // true

strings.Index("golang", "lan") // 2
strings.Index("golang", "py") // -1

strings.Split("apple,banana,cherry", ",") // [apple banana cherry]
strings.Join([]string{"one", "two", "three"}, "-") // one-two-three

strings.TrimSpace("  hello  ") // "hello"
strings.Trim("...hello...", ".") // "hello"
```

---

```yaml
layout: full
```
# `os`

- `os` package provides functions to interact with the operating system.

```go
// Reads a file
data, err := os.ReadFile("example.txt")
if err != nil {
	fmt.Println("Error while reading file", err)
} else {
	fmt.Println(string(data))
}

// Writes to a file
err := os.WriteFile("output.txt", []byte("hello, world"), 0644)
if err != nil {
    fmt.Println("Error while writing file", err)
}

// Gets an environment variable
home := os.Getenv("HOME")
fmt.Println("Home directory:", home)

// Gets command-line arguments
// go run main.go 123 abc hello
args := os.Args
fmt.Println("Program arguments:", args)
```
---

```yaml
layout: full
```
# ...and many more!

|                    |                      |                    |
| ------------------ | -------------------- | ------------------ |
| `bytes`            | `bufio`              | `cmp`              |
| `encoding/json`    | `errors`             | `flag`             |
| `html/template`    | `image/png`          | `log/slog`         |
| `net/http`         | `path/filepath`      | `regexp`           |
| `slices`           | `strconv`            | `time`             |

---

```yaml
layout: section
index: 9
```

## Algorithms

---

```yaml
layout: full
```

# Project structure

- `lab` directory is the workspace where all activities take place.
	- `main.go` serves as a playground for running and testing the algorithms implemented.
	- `go.mod` is required to enable importing the `lab/utils` local package.
	- `utils` directory contains various algorithms to implement.
	- `solutions` folder provides the solutions, available for review after attempting the exercises.

```console
$ cd lab
$ go run main.go
[]
$
```

---

```yaml
layout: full
```

# 📝 Exercise : create slices

- Navigate to the `utils` folder and edit the `slices.go` file.
- Implement the method `GenerateSequentialSlice`.
```go
generated := utils.GenerateSequentialSlice(5)
fmt.Println(generated) // [0 1 2 3 4]
```

- Implement the method `GenerateIdenticalSlice`.
```go
generated := utils.GenerateIdenticalSlice(3, 28)
fmt.Println(generated) // [28 28 28]
```

- Implement the method `GenerateRandomSlice`.
```go
generated := utils.GenerateRandomSlice(4, 16)
fmt.Println(generated) // [7 1 12 6]
```

- Bonus: implement the method `GenerateUserSlice`.
```console
$ go run main.go
Enter the elements separated by space: 10 20 30 40
User-defined slice: [10 20 30 40]
$
```

---

```yaml
layout: full
```

# 📝 Exercise : verify if a slice is sorted

- Navigate to the `utils` folder and edit the `sorts.go` file.
- Implement the method `IsSorted`.

```go
slice1 := []int{0, 2, 4, 5, 6, 9}
fmt.Println(utils.IsSorted(slice1)) // true

slice2 := []int{1, 1, 1}
fmt.Println(utils.IsSorted(slice2)) // true

slice3 := []int{5, 1, 2, 3, 0}
fmt.Println(utils.IsSorted(slice3)) // false
```

---

```yaml
layout: full
```

# Selection sort

- The algorithm repeatedly selects the smallest element from the unsorted portion of the list.
- After selecting the smallest element, it swaps it with the first element of the unsorted portion.
- This process is repeated, each time moving the boundary between the sorted and unsorted portions.

<Youtube id="92BfuxHn2XE" width="592px" height="333px" />

---

```yaml
layout: full
```

# 📝 Exercise : sort a slice using selection sort

- Navigate to the `utils` folder and edit the `sorts.go` file.
- Implement the method `SelectionSort`.

---

```yaml
layout: full
```

# Bubble sort

- It is called _bubble_ because elements gradually move up (like bubbles rising to the surface). 
- The algorithm repeatedly compares adjacent elements.
- If two adjacent elements are in the wrong order, they are swapped.

<Youtube id="Cq7SMsQBEUw" width="592px" height="333px" />

---

```yaml
layout: full
```

# 📝 Exercise : sort a slice using bubble sort

- Navigate to the `utils` folder and edit the `sorts.go` file.
- Implement the method `BubbleSort`.

---

```yaml
layout: full
```

# Merge sort

- This is a divide-and-conquer sorting algorithm.
- It recursively divides the array into two halves.
- Each subarray is then merged back together in sorted order.

<Youtube id="ZRPoEKHXTJg" width="592px" height="333px" />

---

```yaml
layout: full
```

# Merge sort

```
Function MergeSort(array)
    IF length(array) <= 1 THEN
        return  // Base case: array has 1 or 0 elements, already sorted
    END IF

    middle <- length(array) / 2
    left <- array[0...middle]
    right <- array[middle...end]

    MergeSort(left)
    MergeSort(right)

    merge(array, left, right)  // Merge the sorted halves back into the original array
END Function
```

---

```yaml
layout: two-cols-header
```

# Merge sort

::left::

```
Function merge(array, left, right)
	i, j, k <- 0

	// Compare elements from both arrays 
	// and merge them in sorted order
	WHILE i < length(left) AND j < length(right) DO
		IF left[i] < right[j] THEN
			array[k] <- left[i]
			i <- i + 1
		ELSE
			array[k] <- right[j]
			j <- j + 1
		END IF
		k <- k + 1
	END WHILE
```

::right::

```
	// Add remaining elements from the left array
	WHILE i < length(left) DO
		array[k] <- left[i]
        i <- i + 1
		k <- k + 1
	END WHILE

	// Add remaining elements from the right array
	WHILE j < length(right) DO
		array[k] <- right[j]
		j <- j + 1
		k <- k + 1
	END WHILE
END Function
```


---

```yaml
layout: full
```

# 📝 Exercise : sort a slice using merge sort

- Navigate to the `utils` folder and edit the `sorts.go` file.
- Implement the method `MergeSort`.

---

```yaml
layout: full
```

# Linear search

- A simple search algorithm that checks each element in the array one by one.
- Algorithm steps
	- Start at the first element of the array.
	- Compare each element with the target value.
	- Continue until the target is found or all elements are checked.
- Works with unsorted arrays or arrays where sorting is not important.

---

```yaml
layout: full
```

# Binary search

- A faster search algorithm that repeatedly divides a **sorted** slice in half to find the target.
- Algorithm steps
	- Check the middle element of the slice.
	- If the middle element is the target, return its index.
	- If the target is smaller, focus on the left half. If larger, focus on the right half.
	- Repeat this process until the target is found or the sub-slice size becomes 0.

---

```yaml
layout: full
```

# 📝 Exercise : search data in slices

- Navigate to the `utils` folder and edit the `search.go` file.
- Implement the method `LinearSearch`.
- Implement the method `BinarySearch`.

---

```yaml
layout: end
```

**Any questions?**