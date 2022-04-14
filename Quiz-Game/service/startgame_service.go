package service

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func StartGame() {
	csvFileName := flag.String("csv", "../Quiz-Game/data/questions.csv", "questions file in form of csv")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Println(err, "error occured while processing file")
		os.Exit(1)
	}
	r := csv.NewReader(file)
	fmt.Println(r)
}
