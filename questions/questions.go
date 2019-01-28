package questions

type Questions struct{
	question string
	answer string
}


//NewQuestion Constructor  for the Question struct
func NewQuestion(question string, answer string) *Questions{

	return &Questions{question,answer}
}


