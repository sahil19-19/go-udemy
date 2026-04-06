-   every file in a package must declare the package name at the top

-   myInt, myString := func hello() (int, string)
    -   format of returning and saving 2 variables

#   go mod init, go mod tidy

#   type conversion
-   one type of value and convert it into another

#   byte slice
-   if we want to store anything to local machine, we need to convert it into a slice of byte, which is nothing but string in ASCII form
-   type conversion of a string into byte slice
    >   []byte(string_name)

#   runes ?
-   'a' is a rune?

#   nil
-   nil is a value in go which means no value
-   a function that returns an `error` will return nil if no error occurs

#   Exit
-   os.Exit(code int)
    -   makes current program to exit with status given code
    -   if code == 0 > means code executed succesfully, else not

#   rand.Intn()
-   pseudo random generate
-   to make it more random we can use seed
-   wen can use current time as seed
    -   time.Now().UnixNano()
    -   UnixNano returns time as nanoseconds (64 bit number)

##  
-   rand.Seed() sets the seed for the global rand package, which is not thread-safe and can lead to unexpected behavior in concurrent programs.

-   Better to use a local rand.Rand instance to avoid global state issues.
-   ```
    src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	// Generate random integers
	fmt.Println(rng.Int())    // Large random number
	fmt.Println(rng.Intn(100)) // Random number between 0 and 99
    ```

##  Reference types and value types
- slices, map, channels, pointers and functions are reference types
- 