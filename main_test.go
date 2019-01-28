package main

import(
	"testing"
	"reflect"
	"github.com/mbuthiya/goProjects/quizApp/questions"
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

//TestOpenCSV is a unittest for opening and reading the csv file
func TestOpenCSV(t *testing.T){

	got := len(openCSV("got.csv"))
	

	if got <= 1{
		t.Errorf("TestOpenCSV got '%v' but was expecting '>1'",got)
	}
}

// TestGetQuestionsList is a unittest for making question instances from the slices
func TestGetQuestionsList(t *testing.T){
	readerSlices :=  openCSV("got.csv")
	got  := reflect.TypeOf(getQuestionsList(readerSlices))
	expected := reflect.TypeOf(make([]questions.Questions,1))

	if got != expected{
		t.Errorf("TestGetQuestionsList got '%v' but was expecting '%v'",got,expected)
	}
}