package values

type User struct {
	Name string
	Age  int
}

// example
// str:="Hello"
// num:=42
// return Result{Str:str, Int:num}

type Result struct {
	Str     string             // str:="hello"
	Int     int                // Figure out yourself
	Float   float64            // Figure out yourself
	Bool    bool               // flag:=true
	Array   [3]int             // arr:=[3]int{6,7,8}
	Slice   []int              // Figure out yourself
	Map     map[string]int     // mp:=map[string]int{"apple":2,"banana":5}
	User    User               // user:=User{name:"Bob", age:20}
	Ptr     *int               // p:=&num
	AddFn   func(int, int) int // add:=func(a int,b int)int{return a-b}
	Any     interface{}        // var data interface{}="sample data"
	ZeroMap map[string]int
}

func Add(a, b int) int {
	return a + b
}

func BuildValues() Result {
	mp := map[string]int{"apple": 2, "banana": 5}
	x := 10
	var interfaceVal interface{} = 100
	var ZeroMap map[string]int
	res := Result{
		Str:     "go",
		Int:     42,
		Float:   3.14,
		Bool:    true,
		Array:   [3]int{1, 2, 3},
		Slice:   []int{4, 5, 6, 7},
		Map:     mp,
		User:    User{Name: "Alice", Age: 20},
		Ptr:     &x,
		AddFn:   Add,
		Any:     interfaceVal,
		ZeroMap: ZeroMap,
	}
	return res
}
