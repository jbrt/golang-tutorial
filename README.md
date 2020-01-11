# Golang-tutorial

A brief tutorial focused on the Go language.


## Code organization

Golang use text files with .go extension. There's no preprocessor in Golang 
(no .h files like with C language).

## A sample Hello World

Example 01-hello.go file:

```golang
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello World !")
}
```

This code can be compiled like this:

```bash
$ go build 01-hello.go
```

Note: 

Another special Golang keywords print and println can be used for printing 
messages not only functions from fmt package.

## Golang variables

### Variable declaration

Goland variables coulb be declared in several ways:

```golang
package main

import (
	"fmt"
)

// Outside function without initialization
var myString string

// Outside function with variable initialization
var myHello string = "Hello"

// Variables can be declared as a group of variables
var (
	myWorld string = "World !"
	goodBye string = "Goodbye"
)

func main() {
	myString = "This is a string"
	fmt.Println(myString)
	fmt.Println(myHello)
	fmt.Println(myWorld)
	fmt.Println(goodBye)
}
```

### Short variable declaration

Variables can be declared without "var" keyword but only within a block of 
code (Golang inference of type).

```golang
package main

import (
	"fmt"
)

func main() {
    helloWorld := "Hello World !"
    fmt.Println(helloWorld)
}
```

### Pointers variable

Golang is using pointers like C language (but without pointer arithmetic).

```golang
package main

import (
	"fmt"
)

func main() {
    var integer int = 123 
    var myPointer *int // Declaration of the pointer
    myPointer = &integer // Storing memory address of integer variable
    fmt.Println(*myPointer)
}
```

Pointers can be used into function signatures as we show further.

### Default values of variables

Variables uninitialized have default values.

| Type      | Default value |
|-----------|---------------|
| bool      | false         |
| int       | 0             |
| float     | 0.0           |
| string    | ""            |
| pointer   | nil           |
| function  | nil           |
| interface | nil           |
| slice     | nil           |
| channel   | nil           |
| map       | nil           |


## Golang Function

### Functions declaration

Functions are declared with "func" keyword. 
Functions can have no arguments, multiple arguments, no return values or multiples return values.

```golang
package main

import (
	"fmt"
)

func withoutArgsAndNoReturnValue() {
	fmt.Println("No arguments and no return values")
}

func withArgsAndNoReturnValue(alpha int, beta int) {
	fmt.Println(alpha + beta)
}

func withArgsAndReturnValue(alpha int, beta int) int {
	return alpha + beta
}

func withArgsAndMultipleReturnValue(alpha int, beta int) (int, int) {
	return alpha * 2, beta * 2
}

func main() {
    withoutArgsAndNoReturnValue()
    withArgsAndNoReturnValue(100, 100)
    fmt.Println(withArgsAndReturnValue(200, 300))
    a, b := withArgsAndMultipleReturnValue(400, 400)
    fmt.Println(a, b)
}
```

### Variadic functions

A variadic function is a function that accept an undefined number of arguments.

```golang
package main

import (
	"fmt"
)

func variadicSumFunction(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}

func main() {
    fmt.Println(variadicSumFunction(1, 3, 5, 7))
}
```

## Basic Golang control flow

### if else statement

**A simple 'if' statement**

```golang
package main

import (
	"fmt"
)

func main() {
    booleanValue := true
    if booleanValue {
        fmt.Println("Within the if statement")
    }
}
```

**With else statement**

```golang
package main

import (
	"fmt"
)

func main() {
    booleanValue := false
    if booleanValue {
        fmt.Println("Within the if statement")
    } else {
        fmt.Println("Within the else section")
    }
}
```

**Else-if structure**

```golang
package main

import (
	"fmt"
)

func main() {
    level := 2
    if level == 1 {
        fmt.Println("Level value is 1")
    } else if level == 2 {
        fmt.Println("Level value is 2")
    }
}
```

### Switch structure

A switch structure can be useful for removing a lot of else/if statements. 
Switch statement can also have a default choice.

```golang
package main

import (
	"fmt"
)

func main() {
    level := 2
    switch level {
        case 1: 
            fmt.Println("Level 1")
        case 2: 
            fmt.Println("Level 2")
        default:
            fmt.Println("Default choice")
    }
}
```

### For loops

Golang support only one form of loop: the for loop.
But this loop can also mimic while loop.

Classic for loop (like C for loop) with init, condition and increment:

```golang
package main

import (
	"fmt"
)

func main() {
    for i:=0; i < 5; i++ {
        fmt.Println("Value: ", i)
    }
}
```

While-style for loop:

```golang
package main

import (
	"fmt"
)

func main() {
    count := 0
    for count < 5 {
        count++
        fmt.Println("Value: ", count)
    }
}
```

Iterate over data structure:

```golang
package main

import (
	"fmt"
)

func main() {
    data := []int{10, 20, 30, 40}
    for index, value := range data {
        fmt.Println("Index value: ", index, "Data value: ", value)
    }
}
```

### Defer keyword

The defer keyword is useful for executing some code at the end of a function.
For instance, if you open a file you can defer the close of that file at the 
end of the function.

```golang
package main

import (
	"fmt"
)

func main() {
    defer fmt.Println("This will be executed after all other statements")
    fmt.Println("Hello")
    fmt.Println("World")
}
```


## Comparison, arithmetic and logical operators

Comparison operators:

| Operator | Meaning                  |
|----------|--------------------------|
| ==       | Equal to                 |
| !=       | Not equal to             |
| <        | Less than                |
| <=       | Less than or equal to    |
| >        | Greater than             |
| >=       | Greater than or equal to |

Arithmetic operators:

| Operator | Meaning     |
|----------|-------------|
| +        | Addition    |
| -        | Subtraction |
| *        | Product     |
| /        | Division    |
| %        | Modulo      |

Logical operators:

| Operator | Meaning |
|----------|---------|
| &&       | AND     |
| \|\|     | OR      |
| !        | NOT     |


## Arrays, Slices and Maps

### Arrays

An array is a collection of data of the same type. This collection is immutable 
(for mutable collection of data, look to slice section).

Array element's can be accessed by an index number beginning at 0 value.

```golang
package main

import (
	"fmt"
)

func main() {
	// First kind of declaration
	var listOfDays [2]string
	listOfDays[0] = "Sunday"
	listOfDays[1] = "Monday"
	fmt.Println(listOfDays[1])

	// second kind of declaration
	listOfMonths := []string{"March", "April"}
	fmt.Println(listOfMonths[0])
}
```

Once the length of an array has been initialized, **it's not possible to add 
more elements beyond the length that has been defined.** 

### Slices

Slice offer more flexibility than arrays because it's possible to add elements 
beyond the pre established slice size's.

There is two ways for creating a new slice:

```golang
// First version, without initialization
months := make([]string, 3) 

// Second version with initialization 
months := []string{"January", "February", "March"} 
```

Note:

The make keyword is used for declaring several kind of variable without the 
need of initialized them.

Adding elements to a slice:

```golang
months[0] = "January"
months[1] = "February"
months[2] = "March"

months = append(months, "April")
fmt.Println(months[3])
```

Note:

The append function is a variadic one, it's possible to add more than one 
element at a time.

### Maps

A map is an unordered group of variable that is accessed by a key rather thant an index value like with a slice.

Like slices, there is two ways for creating a map. One with content initialization and one without.

```golang
// First version, without initialization
points := make(map[string]int) 

// Second version with initialization 
points := map[string]int{"John": 3, "Mathew": 5}
```

Elements can be added into a map like this:

```golang
points["Mark"] = 3
```

For deleting an element from a map, there is a special keyword for that: delete.

```golang
delete(points, "Mark")
```


## Structures 

A structure is a composite type, it contains one or more base type under a 
unique name.

### Struct declaration

Here is an example of structure declaration:

```golang
type Person struct {
    Name string
    Age  int
}
```

### Struct initialization

Initialization and usage of a structure:

```golang
package main

import (
	"fmt"
)

// Structure declaration
type Person struct {
    Name string
    Age  int
}

func main() {
    // Structure with initialization
    person := Person{Name: "John", Age: 32}
    // Access to structure member's
    fmt.Println(person.Name)
}
```

### Nested structure

Structures can be nested to modelize complex data.

```golang
package main

import (
	"fmt"
)

// Structure declaration
type Person struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Street string
	City   string
}

func main() {
	// Structure with initialization
	person := Person{
		Name: "John",
		Age:  32,
		Address: Address{
			Street: "Main street",
			City:   "London",
		},
	}
	// Access to structure member's
	fmt.Println(person.Address.Street)
}
```

### Public and private values

With Golang it's possible to restrict the visibility of function, variables and
structures to reproduce the objet concept of encapsulation.

This mecanism works with Golang naming convention, there is no specific keyword
"private" or "public" like with Java or C++ languages.

Any variables, functions, structures with names starting with an uppercase 
letter will be exported and visible from outside their package. 

Conversely, variables, functions structures with names starting with a 
lowercase letter won't be available outside their package.


## Using methods and interfaces

### Methods

Golang is not an oriented object language. There is no inheritance or class definition (but Golang provides a kind of polymorphism).

But you can create methods attached to structures (and other Golang types) for manipulating data and thus reproduce a kind of encapsulation (specificaly if 
using the public/private values decribes in the previous chapter).

Here is an example:

```golang
package main

import (
	"fmt"
)

// Structure declaration
type Person struct {
	name string
	age  int
}

func (person Person) GetName() string {
	return person.name
}

func main() {
	// Structure with initialization
	person := Person{
		name: "John",
		age:  32,
	}
	// Access to structure member's
	fmt.Println(person.GetName())
}
```

Method declaration are almost identicaly than normal function but they have an 
extra field between the func keyword and the name of the function/method.

This field is called a receiver and it describes the relationship between this 
method and the target Golang data type. Here, it's a relationship between this 
method called GetName and the Person data structure. 

Inside the method it's possible to access to the data field of the Person 
structure with the variable name given into the receiver.

### Methods and pointers

For changing structures values with methods, pointers must be used. If not, 
values won't change. Pointers must be declared with * character into the 
receiver.

```golang
package main

import (
	"fmt"
)

// Person structure declaration
type Person struct {
	name string
	age  int
}

// GetName return the name of the person
func (person Person) GetName() string {
	return person.name
}

// SetName can change the name of one person
// Here the receiver is a pointer to the structure.
// If not we can't change the person's name
func (person *Person) SetName(name string) {
	person.name = name
}

func main() {
	// Structure with initialization
	person := Person{
		name: "John",
		age:  32,
	}
	// Access to structure member's
	fmt.Println(person.GetName())

	// Change the name of this person
	person.SetName("William")
	fmt.Println(person.GetName())
}
```

### Interfaces

Interfaces can be compared as blueprints for method sets. An interface 
describes methods signatures without implement it.

```golang
type Rectangle interface {
    Perimeter() float64
}
```

There is no need to implement (with a keyword) interfaces in Golang, 
interfaces are **implicitly implemented**. If a structure implement all 
methods described into an interface, then this structure is compliant to the
interface.

```golang
package main

import (
	"fmt"
)

// Rectangle is an interface for describing Perimeter method
type Rectangle interface {
	Perimeter() float64
}

// Square is a kind of rectangle
type Square struct {
	length float64
	width  float64
}

// Perimeter method
func (square Square) Perimeter() float64 {
	return square.length*2 + square.width*2
}

// DisplayPerimeter will take a rectangle (so an interface) as parameter
func DisplayPerimeter(rectangle Rectangle) {
	fmt.Println(rectangle.Perimeter())
}

func main() {
	s := Square{length: 2.0, width: 2.0}
	DisplayPerimeter(s)
}
```