package questions

type Questions struct{
	Question string
	Answer string
}


//NewQuestion Constructor  for the Question struct
func NewQuestion(question string, answer string) *Questions{

	return &Questions{question,answer}
}

//ValidateAnswer method that allows us to validate if the answer given is correct for the question.
func(q *Questions) ValidateAnswer(answer string) bool{
	if answer == q.Answer{
		return true
	}
	return false
}


