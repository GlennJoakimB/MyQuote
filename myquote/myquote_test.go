package myquote

import "testing"

//Eksempel på test for retur-verdie

func TestMyQuote(t *testing.T) {
	wanted := "Ahoy, world!"
	got := MyQuote()

	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}

}
