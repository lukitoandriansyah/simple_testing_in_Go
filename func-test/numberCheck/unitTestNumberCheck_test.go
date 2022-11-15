package numberCheck

import (
	"fmt"
)

func ExampleCheckNumber_WithInputLessThanZero() {
	fmt.Println(CheckNumber(-1))
	//Output:
	//small number
}
func ExampleCheckNumber_WithInputGreaterThanZeroAndBelowOneHundred() {
	fmt.Println(CheckNumber(91))
	//Output:
	//medium number
}
func ExampleCheckNumber_WithInputGreaterThanOneHundred() {
	fmt.Println(CheckNumber(100))
	//Output:
	//big number
}
