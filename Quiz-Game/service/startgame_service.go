package service

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	q, a string
}

func LoadGame() {
	csvFileName := flag.String("csv", "../Quiz-Game/data/questions.csv", "questions file in form of csv")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Println(err, "error occured while processing file")
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	problems := ParseLines(lines)
	fs := StartGame(problems)
	fmt.Printf("your score is %d out of %d\n", fs, len(problems))
}

func ParseLines(lines [][]string) (pbs []problem) {
	for _, line := range lines {
		var p problem
		p.q, p.a = line[0], line[1]
		pbs = append(pbs, p)
	}
	return
}

func StartGame(pbs []problem) (finalScore int) {
	in := bufio.NewScanner(os.Stdin)
	for _, pb := range pbs {
		fmt.Printf("%s: ", pb.q)
		in.Scan()
		if in.Text() == pb.a {
			finalScore++
		}
	}
	return
}
