# Go Interfaces: Things I Learned

## 1. A Nil Interface Can Be a Bug

An interface can contain a nil pointer and still not be nil itself. This can cause unexpected bugs in Go programs.

### Example

```go
type Payment interface {
    Pay()
}

type MTN struct{}

func (m *MTN) Pay() {
    fmt.Println("Paying with MTN")
}

var mtn *MTN = nil
var payment Payment = mtn

fmt.Println(payment == nil) // false
```

### Why does this happen?

- `mtn` is a pointer to an MTN struct, but it currently holds `nil`.
- When we assign `mtn` to `payment`, the interface receives the type `*MTN` and its value, which is nil.
- The interface is no longer a completely nil interface because it now knows the type of the value it holds.

Therefore, `payment == nil` returns `false`.

### What is a truly nil interface?

```go
var payment Payment

fmt.Println(payment == nil) // true
```

Here, the interface holds neither a type nor a value. This is a truly nil interface.

### How can we avoid this bug?

Check whether the pointer is nil before assigning it to an interface.

```go
var mtn *MTN = nil

if mtn != nil {
    var payment Payment = mtn
    fmt.Println(payment == nil)
} else {
    fmt.Println("MTN provider is unavailable")
}
```

By checking the pointer first, we avoid assigning a nil pointer to the interface.

---

## 2. What Is a Pointer?

A pointer is a value that holds the memory address of another value.

```go
x := 10
p := &x
```

- `x` holds the value `10`.
- `&x` obtains the address of `x`.
- `p` stores that address.

Because `p` holds the address of `x`, we say that `p` points to `x`.

## 3. Understanding `&` and `*`

### The `&` symbol

The `&` symbol obtains the address of a value.

```go
x := 10
p := &x
```

`&x` means: Give me the address of `x`.

### The `*` symbol

The `*` symbol has two common uses when working with pointers.

**A. Declaring a pointer type**

```go
var p *int
```

This means `p` is a pointer to an integer.

**B. Dereferencing a pointer**

```go
x := 10
p := &x

fmt.Println(*p) // 10
```

`*p` accesses the value stored at the address that `p` holds.

We can also change the original value through the pointer:

```go
x := 10
p := &x

*p = 50

fmt.Println(x) // 50
```

Because `p` points to `x`, changing the value through `*p` changes `x` itself.

### Quick summary

| Syntax | Meaning                                 |
| ------ | --------------------------------------- |
| `&x`   | Obtains the address of `x`              |
| `*int` | Declares a pointer type to an integer   |
| `*p`   | Accesses the value that `p` points to   |
| `p`    | The pointer variable holding an address |

The pointer is the variable holding the address, not the `&` or `*` symbol.

---

## 4. Returning Pointers from Functions

A Go function can return a pointer to a struct.

```go
type MTN struct {
    Name string
}

func GetMTNProvider() *MTN {
    return &MTN{}
}
```

### Understanding the function

- `*MTN` in the function signature means the function returns a pointer to an MTN struct.
- `MTN{}` creates an MTN struct value.
- `&MTN{}` obtains the address of that struct value.
- `&MTN{}` produces a pointer of type `*MTN`.

The function promises to return a pointer, and it does exactly that.

### Using the returned pointer

```go
provider := GetMTNProvider()

provider.Name = "Airtel"

fmt.Println(provider.Name) // Airtel
```

- `GetMTNProvider()` returns a pointer to an MTN struct.
- `provider` stores that pointer.
- `provider.Name = "Airtel"` modifies the struct through that pointer.

Go automatically dereferences the pointer when accessing struct fields.

We could also write:

```go
(*provider).Name = "Airtel"
```

But `provider.Name` is the simpler and normal way in Go.

### Value return vs pointer return

Returning a struct value:

```go
func GetMTNProvider() MTN {
    return MTN{}
}
```

Returning a pointer:

```go
func GetMTNProvider() *MTN {
    return &MTN{}
}
```

| Return type | Meaning                            |
| ----------- | ---------------------------------- |
| `MTN`       | Returns an MTN struct value        |
| `*MTN`      | Returns a pointer to an MTN struct |

---

## 5. Important Lessons

- An interface can hold a nil pointer and still not be nil itself.
- A truly nil interface has neither a dynamic type nor a dynamic value.
- A pointer holds the memory address of another value.
- `&` obtains an address.
- `*` declares a pointer type or dereferences a pointer.
- `*MTN` is a pointer type to an MTN struct.
- `&MTN{}` creates and returns a pointer to an MTN struct.
- Go allows us to access struct fields through pointers without explicitly dereferencing them.

**Learning Go one step at a time.**
