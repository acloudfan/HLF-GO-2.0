/**
 * GoLang Snippets for reference
 * Run in the VM using the command: go run $GOPATH/src/token/snippets/*.go
 **/
package main

import "fmt"
import "time"
import "encoding/json"
import str "strings"
import "strconv"

func main() {

	fmt.Println("hello world")
	
	// Var declaration - type picked from value
	var  i=10
	// Var declare & initialize without keyword var
	v := 5

	// Using time
	var x = time.Now()

	// Multiple prints to console
	fmt.Println(x.Weekday(), v)

	// for loop
	for  k:=10; k < 13; k++ {
		fmt.Println(k)
		// if statement
		if k==i {
			fmt.Println("k==i")
		}
	}

	// switch without the arg is like if then else if ..
	t := time.Now().Hour()
	switch {
	case t < 12:
		fmt.Println("Good Morning")
	case t >= 12 && t<=17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Night")
	}

	// switch with arg
	switch t {
		// Same as above switch statement
	}


	// Array
	var arr[3] int;
	for  k:=10; k < 13; k++ {
		arr[k-10]=k
	}
	fmt.Println(arr)

	// Dynamic array referred to as slice
	// Array of string with initial size = 1
	darr := make([]string, 1)
	darr[0]="one"
	darr=append(darr, "two")
	fmt.Println(darr)

	// Map = associative arrays | 
	m:=make(map[string]int)
	m["one"]=1
	m["two"]=2
	fmt.Println("Original Map",m)
	delete(m,"one")
	fmt.Println("Map after delete", m)

	// Range = Iterator for Array|Slice|Map
	for index, value :=range darr {
		fmt.Printf("Range %d -> %s\n",index,value)
	}
	for key := range m {
		fmt.Printf("Map Range %s -> %d\n",key,m[key])
	}

	// Call mult return function
	_, k := multi()
	v, k = multi()
	fmt.Printf("Multi return func = %d, %s\n", v,k)

	// call to variadic func - add more numbers if you want
	variadic(1,2,3)

	bef_v, bef_f := 10,10
	fmt.Printf("Before v= %d,  f= f%d\n",bef_v, bef_f)
	// Pointer to variable - use &
	ptr(&bef_v, bef_f)
	fmt.Printf("After v= %d,  f= f%d\n",bef_v, bef_f)

	// Refer to the person struct defined below
	personInstance := person{name:"johny",age:43}
	fmt.Println("Struct-1", personInstance)

	// Methods = functions defined for structs
	newAge := personInstance.doubleAge()
	fmt.Println("Struct-Method", newAge, )

	// JSON - look at the jsonFunc()
	jsonFunc()

	// String funcs
	strFunc()

	// conversion functions in package strconv
	conversions()
}

// funtion
func plus(a int, b int) int {
	return a+b
}

// multiple return func
func multi() (int, string){
	return 10, "ten"
}

// Variadic function - takes multiple arguments
func variadic(nums ... int) {
	fmt.Println("Variadic func ", nums)
}

func  ptr(v *int, f int){
	*v++	// This will change value of original var
	f++		// No impact on the original
}

// Structs
type person struct {
	name string
	age  int  // Initialized to 0 if not provided
}

// Method - struct instance may be passed by value or reference
func  (p *person) doubleAge() (int) {
	p.age = 2*p.age
	return p.age
}

// Struct with JSON tags
type account struct {
	Name    string `json:"name"`
	AcNum   int    `json:"accountnumber"`
}

func jsonFunc() {
	johnAccount := account{Name: "john", AcNum:4312}
	s, err := json.Marshal(johnAccount)
	fmt.Println("jsonFunc 1", string(s), err, johnAccount)

	// Unmarshal - takes the content of the marshal and converts to struct instance
	var newJohnAccount account

	if err := json.Unmarshal(s,&newJohnAccount); err != nil {
		// Fail fast - Abort - https://gobyexample.com/panic
        panic(err)
    }
    fmt.Println("jsonFunc 2 ",newJohnAccount.Name," ",newJohnAccount.AcNum)
}

func  strFunc() {
	// Hold reference to a function
	var p=fmt.Println
	p()
	p("String functions")
	p("================")
	p("Contains:  ", str.Contains("test", "es"))
    p("Count:     ", str.Count("test", "t"))
    p("HasPrefix: ", str.HasPrefix("test", "te"))
    p("HasSuffix: ", str.HasSuffix("test", "st"))
    p("Index:     ", str.Index("test", "e"))
    p("Join:      ", str.Join([]string{"a", "b"}, "-"))
    p("Repeat:    ", str.Repeat("a", 5))
    p("Replace:   ", str.Replace("foo", "o", "0", -1))
    p("Replace:   ", str.Replace("foo", "o", "0", 1))
    p("Split:     ", str.Split("a-b-c-d-e", "-"))
    p("ToLower:   ", str.ToLower("TEST"))
    p("ToUpper:   ", str.ToUpper("test"))
	p()
	p("Len: ", len("hello"))
    p("Char:", "hello"[1], string("hello"[1]))
}

// Convert bsic datatypes from/to string
func conversions() {
	var someStr string

	// convert string to number
	someStr = "32"
	i, _ := strconv.Atoi(someStr)
	
	// The parse functions return the widest type (float64, int64, and uint64), but if the size argument 
	// specifies a narrower width the result can be converted to that narrower type without data loss
	// multiple Parse?? funcs available
	i64, _ := strconv.ParseInt(someStr, 10, 32)
	i32 := int32(i64)
	// convert string to boolean
	someStr = "true"
	b, _ := strconv.ParseBool(someStr)

	fmt.Println("\nstrconv.Atoi=", i, " i32=", i32, " bool=", b)

	// Convert basic types to string 
	// multiple Format??? functions
	someStr = strconv.FormatInt(int64(i32), 10)
}

// JSON and G
// https://blog.golang.org/json-and-go
// JSON Tags https://godoc.org/encoding/json