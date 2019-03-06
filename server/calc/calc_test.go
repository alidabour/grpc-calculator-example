package calc

import "fmt"

func ExampleEvaluate() {
	fmt.Println(Evaluate("2*2/2"))
	// Output: 2 <nil>
}

func ExampleItoP() {
	fmt.Println(ItoP([]string{"1", "+", "2"}))
	// Output: [1 2 +]
}
