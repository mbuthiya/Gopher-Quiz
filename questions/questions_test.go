package questions

import(
	"testing"
	"reflect"
)

//  TestQuestionConstructor  is a function to test the question Constructor
func TestQuestionConstructor(t *testing.T){
	questionExample := Questions{"How Old are you","Ten"}
	got := reflect.TypeOf(NewQuestion("How Old are you","Ten"))
	expected := reflect.TypeOf(questionExample)

	if reflect.DeepEqual(got,expected){
		t.Errorf("TestQuestionConstructor got: '%v' but expected : '%v'",got,expected)
	}
}

// TestValidateAnswer is a function that test the ValidateAnswer method and returns a boolean
func TestValidateAnswer(t *testing.T){
	questionExample := Questions{"How Old are you","Ten"}
	got := questionExample.ValidateAnswer("Ten")
	expected :=true

	if got != expected{
		t.Errorf("TestValidateAnswer expected: '%v' but got '%v'",expected,got)
	}
	
}