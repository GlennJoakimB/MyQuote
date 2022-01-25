package myquote

import (
	"fmt"

	"rsc.io/quote"
)

func MyQuote() string {
	// fmt.Println(quote.Hello())
	// fmt.Println(quote.Go())
	// fmt.Println(quote.Glass())
	// fmt.Println(quote.Opt())

	q1 := quote.Hello()
	q2 := quote.Go()
	q3 := quote.Glass()
	q4 := quote.Opt()
	fmt.Println(q1)
	fmt.Println(q2)
	fmt.Println(q3)
	fmt.Println(q4)

	return q1 //returverdi for myquote_test.go

}
