package strategy

import "fmt"

func Example() {
	operator := NewOperator(&Add{})
	result := operator.Calculate(1, 2)
	fmt.Println("Do add:", result)

	operator.SetStrategy(&Reduce{})
	result = operator.Calculate(2, 1)
	fmt.Println("Do reduce:", result)
	// Output:
	// Do add: 3
	// Do reduce: 1
}
