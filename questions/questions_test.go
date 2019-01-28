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