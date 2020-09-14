package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	csvFile := flag.String("csv","problem.csv", "csv file in format'quelstion,answer'")
	timeLimit := flag.Int("limit",30,"time limit in seconds")
	flag.Parse()


	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Can't open the CSV file: %s", *csvFile))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Couldn't read lines")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	//fmt.Println(problems)
	correct :=0
	for i,p := range problems{
		select{
		case <-timer.C:
			fmt.Printf("You scored %d from %d\n", correct, len(problems))
			return

		default:
			fmt.Printf("Problem #%d: %s=\n", i+1, p.q)
			var answer string
			fmt.Scanf("%s\n",&answer)
			if answer == p.a {
				correct++
		}

		}
	}
	fmt.Printf("You scored %d from %d\n", correct, len(problems))
}

func parseLines  (lines [][] string) []question{
	ret := make([]question, len(lines))
	for i,line := range lines{
		ret[i] = question{
			q:line[0],
			a:strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type question struct{
	q string
	a string
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}
