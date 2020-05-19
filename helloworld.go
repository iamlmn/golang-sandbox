package main //go programs consist of packages, and start with main
import "fmt"

// import a library/package. In this case we import fmt, which brings in 
// standard formatted I/O strings from C


// syntax for a function in Go. 
// define, name, parms, and your code goes in curly braces.
func main(){
	//print line just outputs what you mass. 
	// Println is a function from the fmt library
	// we can reference other packages from fmt with.
	fmt.Println("Hello world!")
}