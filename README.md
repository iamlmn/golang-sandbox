# golang-sandbox


Go is a statically typed, compiled programming language designed at Google.


## About Go.
Proceddural, functional and concurrent.
Compiled Language.
Statically typed language
Provides memory management.
Concurrency is inbuilt.

Go was developed by Google to solve challenges that they were facing, purpose built from the ground up. A language like Python is great for rapid development and many tasks, but, due to the GIL and dynamic typing, it can sometimes be slow. If you scale out Python to millions of instances, you pay quite the price in performance and cost as opposed to if you had written the same program in C++. The problem is, however, C++ isn't the most enjoyable language to program in, nor is it the most friendly to learn.

Go is a static-typed language (as opposed to dynamic typing with something like Python), and has concepts like concurrency baked in right from the start. For me personally, my main interest in Go is for web development, but I can see it being useful for a variety of high-volume types of tasks and just systems in general. Not only does Go have web frameworks, the language itself comes out of the gate with full web support, a built in web server, and could be considered already a low-level web framework.


Download go from here[https://golang.org/doc/install]

## Godoc
```
godoc fmt Println
```
## GO Language Types:
Go is a static-typed language, meaning you need to specify the type, and the type cannot change without you explicitly changing it

- bool
- string
- int  int8  int16  int32  int64
- uint uint8 uint16 uint32 uint64 uintptr
- byte // alias for uint8
- rune // alias for int32
     // represents a Unicode code point
- float32 float64
- complex64 complex128


## godoc on ResponseWriter
ResponseWriter: godoc net/http ResponseWriter


## Structs in GO!

NO Classes. Only Structs in GO.