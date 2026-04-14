package main

/*  go build -o vet-test
- # (No error! The program compiles successfully.)
	./vet-test
- Hello, %!s(MISSING)!
- go vet ./...
	- main.go:10:21: fmt.Printf format %s reads arg #1, but call has 0 args

- This shows the clear difference: go build creates a working but buggy program,
while go vet finds the bug before it ever gets to production.

*/

// fmt - formatter // vet - linter
func main() {

	/* return
	fmt.Println("Hello World")
	*/
	/*
		- If we compile above code using go build it will compile successfully versus if we run go vet ./... it will
		give us error - unreachable code
	*/

	//fmt.Printf("Hello, %s! \n") // main.go:10:21: fmt.Printf format %s reads arg #1, but call has 0 args
	//fmt.Printf("Hello, %s! \n", "world")

	checkNetworkLimitation()

}
