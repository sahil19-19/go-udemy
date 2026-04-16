# new vs make
- new in golang can allocate in both heap and stack
- is decided by the compiler
```
func foo() *int {
    a := new(int)
    *a = 10
    return a
}

Now a escapes the function → Go allocates it on the heap.
```

# type conversion
- int to character
```
i := 90
fmt.Println(string(i)) // > "Z" (ascii of 90)
```
- treats it like Unicode character

- int to string representation of int
```
i:= 90
// need import "strconv"
s := strconv.Itoa(i)
fmt.Println(s) // Outputs: "90"
```

- string to int
```
s := "90"
i, err := strconv.Atoi(s)

i64, err := strconv.ParseInt(s, 10, 64) // base, bit size
```
