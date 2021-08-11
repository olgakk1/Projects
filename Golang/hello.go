//importing packages
package main

//creating global variables without setting
//their types
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

//creating global variables with setting
//their types
var (
	a      int
	b      bool
	s      string
	apple  int
	bee    bool
	apple2 string
)

//creating a struct
type Person struct {
	name string
	job  string
}

//practice user Terminal input and output
func userIO() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter you city: ")
	city, _ := reader.ReadString('\n')
	fmt.Printf("You live in " + city)
}

//practice setting up variables locally
func setLocalVariables() {
	a, b, s, apple = 5, false, "abc", 6
	fmt.Printf("a: %d apple: %d\n", a, apple)
	a, apple = apple, a
	fmt.Printf("a: %d apple: %d\n", a, apple)
}

//run for loops
func runForLoops1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

//run for loops
func runForLoops2() {
	// var arr1 = new([5]int)
	// var a = []int64{1, 2, 3, 4}
	// var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	var arrLazy = [...]int{5, 6, 7, 8, 22}
	for _, element := range arrLazy {
		fmt.Println(element)
	}
}

//setting up if and else statements
func setConditionals() {
	var x = 1
	if x > 2 {
		fmt.Printf("x is greater than 2")
	} else {
		fmt.Printf("condition is false (x > 2)")
	}
}

//run while loops
func runWhileLoops() {
	var i = 1
	var max = 20
	for i < max {
		fmt.Println(i)
		i += 1
	}
}

//check if the file exists
func checkFileExists1() {
	if _, err := os.Stat("file-exists.go"); err == nil {
		fmt.Printf("File exists\n")
	} else {
		fmt.Printf("File does not exist\n")
	}
}

//check if the file exists
func checkFileExists2() {
	if _, err := os.Stat("file-exists2.file"); os.IsNotExist(err) {
		fmt.Printf("File does not exist\n")
	}
}

//read a file
func readFile() {
	b, err := ioutil.ReadFile("fileIn.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	fmt.Println(str)
}

//read a file line by line
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

//read a file line by line main
func readLinesMain() {
	lines, err := readLines("fileIn.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for i, line := range lines {
		fmt.Println(i, line)
	}
}

// write to a file
func writeFile() {
	file, err := os.Create("fileOut.txt")
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString("I wrote to this file")
}

//rename a file
func renameFile() {
	src := "changeName1.txt"
	dst := "changeName2.txt"
	os.Rename(src, dst)
}

//using structs
func useStructs() {
	var person Person
	person.name = "Albert"
	person.job = "Einstein"
	fmt.Printf("person.name = %s\n", person.name)
	fmt.Printf("person.job = %s\n", person.job)
}

//using maps
func useMaps() {
	alpha := make(map[string]int)
	alpha["A"] = 1
	fmt.Println(alpha["A"])
	// alpha2 := map[string]int{
	// 	"A": 1,
	// 	"B": 2,
	// 	"C": 3,
	// }
	website := map[string]map[string]string{
		"Google": map[string]string{
			"name": "Google",
			"type": "Search",
		},
		"YouTube": map[string]string{
			"name": "YouTube",
			"type": "video",
		},
	}
	fmt.Println(website["Google"]["name"])
	fmt.Println(website["Google"]["type"])
}

//using random numbers
func useRandom() {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(10-0) + 0
	fmt.Printf("Random number: %d\n", randomNum)
}

//using Addresses
func useAddresses1() {
	var x int = 5
	fmt.Printf("Address of variable x: %x\n", &x)
	fmt.Printf("Value of variable x: %d\n", x)
}

//using Addresses
func useAddresses2() {
	var x int = 5
	var ipointer *int
	ipointer = &x
	fmt.Printf("Memory address of variable x: %x\n", &x)
	fmt.Printf("Memory address stored in ipointer variable: %x\n", ipointer)
	fmt.Printf("Contents of *ipointer variable: %d\n", *ipointer)
}

//using Slices (Sets)
func useSlices() {
	myset := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s := myset[0:4]
	fmt.Println(s)
	mystring := "Go programming"
	mysubstring := mystring[0:2]
	fmt.Println(mysubstring)
}

//using defer
func hello() {
	fmt.Println("Hello")
}

func who() {
	fmt.Println("Go")
}

func useDefer() {
	defer who()
	hello()
}

//using multiple returns
func useMultipleReturns() (int, int) {
	return 2, 4
}

//using Variadic functions
func useVariadicFunctions(numbers ...int) {
	total := 0
	for _, num := range numbers {
		total += num
	}
	fmt.Println(total)
}

//setting up recursion
func useRecursion(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * useRecursion(x-1)
}

//main function calling on other functions
func main() {
	// userIO()
	// setLocalVariables()
	// runForLoops1()
	// runForLoops2()
	// setConditionals()
	// runWhileLoops()
	// checkFileExists1()
	// checkFileExists2()
	// readFile()
	// readLinesMain()
	// writeFile()
	// renameFile()
	// useStructs()]
	// useMaps()
	// useRandom()
	// useSlices()
	// useAddresses1()
	// useAddresses2()
	// useDefer()
	// x, y := useMultipleReturns()
	// fmt.Println(x)
	// fmt.Println(y)
	// useVariadicFunctions(2, 3, 4)
	// fmt.Println(useRecursion(5))
	//using go and channels
}

//ADDITIONAL NODES
// := is simple declaration
// = is more difficult declaration (need to declare the variable somewhere else)

//CODE REFERENCE
//https://golangr.com/exercises/
