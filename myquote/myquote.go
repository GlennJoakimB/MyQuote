package myquote

import (
	"fmt"

	"rsc.io/quote"
)

func MyQuote() {
	fmt.Println(quote.Hello())
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
}
