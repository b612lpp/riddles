package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//Riddle структура содержит в себе загадку и ответы на нее
type Riddle struct {
	question, answer1, answer2 string
}

func main() {

	var SlicaOfRiddles []Riddle
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/riddles")
	if err != nil {
		fmt.Println(err)
	}

	GetRiddleDB, err := db.Query("SELECT question, answer1, answer2 FROM qada")
	for GetRiddleDB.Next() {
		var question, answer1, answer2 string

		err = GetRiddleDB.Scan(&question, &answer1, &answer2)

		s := Riddle{question, answer1, answer2}
		SlicaOfRiddles = append(SlicaOfRiddles, s)

	}

	PrintQuestion(SlicaOfRiddles)

}

//PrintQuestion draws the questions
func PrintQuestion(SlicaOfRiddles []Riddle) {
	var UserAnswer string
	var ScoreCounter int

	RiddlesLen := len(SlicaOfRiddles)
	fmt.Println("Длинна массива составляет : ", RiddlesLen, " загадок")
	if RiddlesLen > 0 {
		rand.Seed(time.Now().UnixNano())
		riddleID := rand.Intn(RiddlesLen)
		fmt.Println(SlicaOfRiddles[riddleID].question)
		//fmt.Println(SlicaOfRiddles[riddleID])
		fmt.Println("Ваш ответ?")
		fmt.Scan(&UserAnswer)
		if UserAnswer == SlicaOfRiddles[riddleID].answer1 || UserAnswer == SlicaOfRiddles[riddleID].answer2 {
			fmt.Println("win")
			ScoreCounter = ScoreCounter + 1
			SlicaOfRiddles = append(SlicaOfRiddles[:riddleID], SlicaOfRiddles[riddleID+1:]...)
			//fmt.Println("Длинна массива составляет: ", RiddlesLen)
			PrintQuestion(SlicaOfRiddles)
		} else {
			fmt.Println("Ответ неверный")
			SlicaOfRiddles = append(SlicaOfRiddles[:riddleID], SlicaOfRiddles[riddleID+1:]...)
			PrintQuestion(SlicaOfRiddles)
		}
	} else {
		fmt.Println("Игра законченаю Ваш Счет: ", ScoreCounter)
	}

}
