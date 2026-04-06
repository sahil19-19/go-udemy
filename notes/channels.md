-   channels and go routines are used to perform concurrent programming

#   go routines
-   we use the `go` keyword
    -   we only use `go` keyword while calling functions
-   it creates a new go routine to specifically execute the code inside the function call
-   when the control runs into a blocking call, it tells program to continue executing rest of the code
-   when we start executing our program, a default go routine is created, `main routine`
    -   `main routine` has to be made aware to wait for child routines to complete

##  go scheduler
-   interface btween the single CPU core and the go routines
-   by default go on only on 1 core, but we can change this behavior
-   scheduler in such case would run one thread on each logical core

##  concurrency/ parallelism
-   concurrency > our program has the ability to run threads of our code
    -   we schedule code and change btw one thread to another

-   parallelism > multiple threads executed at the same time 
    -   requires multiple CPUs


#   Channels
-   a way for routines to communicate with one another
-   channels can be made just like other variables like stirng, int etc
-   channels of a type can only pass one type of data, ie., string or int or float
-   writing into channel
```
c <- "hello"
```

-   reading from a channel
```
temp := <-c
```

```
fmt.Println(<- myChan)
```

#   function literal 
-   like lambda function in C++, anonymous function in js
-   a function without a name

#   note
-   we dont use same variable in 2 different routines