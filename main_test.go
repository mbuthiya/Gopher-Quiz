package main

import(
	"testing"
)

//TestValidateFlag is a unittest for the ValidateFlag function
func TestValidateFlag(t *testing.T){

	flagName := "got"
	got := ValidateFlag(flagName)
	expect := true

	if got!= expect{
		t.Errorf("TestValidateFlag got '%v' but was expecting '%v'",got,expect)
	}
}