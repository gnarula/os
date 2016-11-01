### Synopsis
The main goroutine spawns two goroutines, even and odd which increase the count
variable accordingly. The increments are done until the count variable reaches
a value of 15, after which the program exits.

```
go run main.go
```

```
even count = 1
odd count = 2
even count = 3
odd count = 4
even count = 5
odd count = 6
even count = 7
odd count = 8
even count = 9
odd count = 10
even count = 11
odd count = 12
even count = 13
odd count = 14
even count = 15
Waited for two threads. Final count value 15 done
```

### Why?
This problem was given as an OS assignment to be done using pthreads in C.
Thought it might be a good opportunity to try this in golang as well.
