package main

import(
	"flag"
	"os"
	"log"
	"fmt"
	"bufio"
	"encoding/csv"
	"github.com/mbuthiya/goProjects/quizApp/questions"
)

//ValidateFlag allows us to validate if the user has inputed a flag Item
func ValidateFlag(flagName string)  bool{
	if flagName == ""{
		return false
	}
	return true;
}


// Read csv file
func openCSV(csvfile string) [][]string{
	game,err := os.Open(csvfile)

		// Error handling
		if err!= nil{
			fmt.Println("I was unable to open the file")
			log.Fatal("Could not open the CSV File")
			os.Exit(2)
		}
		// Opening CSV
		reader := csv.NewReader(bufio.NewReader(game))
		readerSlices,err := reader.ReadAll()

		// Error Handling
		if err != nil{
			fmt.Println("I was unable to read the file")
			log.Fatal("Could not Read the CSV File")
			os.Exit(2)
		}
		return readerSlices
}

//getQuestionsList takes nested string slice and returns a slice of Questions
 func getQuestionsList(readerSlices [][]string) []questions.Questions{

	questionSlices := make([]questions.Questions,1)
	for _,questionSlice:= range readerSlices{

		newQuestion := questions.NewQuestion(questionSlice[0],questionSlice[1])
		questionSlices = append(questionSlices,*newQuestion)
	}
	return questionSlices
}



func main(){

	//var count int
	
	// Creating flags
	var csvString string
	flag.StringVar(&csvString,"csv","","Add the csv link you want to connect")
	flag.Parse()
	

	if ValidateFlag(csvString){	
		readQuestions := openCSV(csvString)
		// Remove the header text
		readQuestions = readQuestions[1:] 

		for _,quizQuestion := range getQuestionsList(readQuestions){

			fmt.Println(quizQuestion.Question)
		}
	}


}