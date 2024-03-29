package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main()  {

	fmt.Print("Enter a grade : ")
	reader := bufio.NewReader(os.Stdin)
	input,err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade,err := strconv.ParseFloat(input,64)
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passed"
	}else {
		status = "failed"
	}

	fmt.Println("A grade of",grade,"is",status)
}