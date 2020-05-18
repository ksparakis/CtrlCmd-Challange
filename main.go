package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type void struct{}

var emptyValue void

func readCSV(filePath string) (int, int, int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)

	uploadCounter := 0
	jeffCounter := 0
	UserMap := make(map[string]void) //  creating a map with an empty struct which consumes 0 memory
	isFirstRow := true
	for {
		// Read each record from csv
		record, err := csvReader.Read()

		//If we have reached end of  the  file  break out of loop
		if err == io.EOF {
			break
		}

		if isFirstRow {
			isFirstRow = false
			continue
		}

		//  if reading line from csv fails return
		if err != nil {
			return 0, 0, 0, fmt.Errorf("Failed to read Line in csv: %v", err)
		}

		//Turn  data into variables for ease of use
		timestamp, username, operation := record[0], record[1], record[2]

		// Turn size into an int
		size, err := strconv.Atoi(record[3])
		if err != nil {
			return 0, 0, 0, fmt.Errorf("Incorrect Row Size Data, will not convert to int: %v", err)
		}

		//1. How many users accessed the server?
		//Use a map to keep  track of users
		UserMap[username] = emptyValue

		//2. How many uploads were larger than `50kB`?
		if operation == "upload" && size > 50 {
			uploadCounter += 1
		}

		//3. How many times did `jeff22` upload to the server on April 15th, 2020?
		if strings.Contains(timestamp, "Apr 15") &&
			strings.Contains(timestamp, "2020") &&
			username == "jeff22" &&
			operation == "upload" {

			jeffCounter += 1
		}

	}

	return len(UserMap), uploadCounter, jeffCounter, nil
}

func main() {
	solution1, solution2, solution3, err := readCSV("server_log.csv")
	if err != nil {
		fmt.Printf("Fatal Error %v", err)
		os.Exit(2)
	}

	fmt.Printf("%v users accesed the server\n", solution1)
	fmt.Printf("%v files where larger than 50kb\n", solution2)
	fmt.Printf("%v files where uploaded by jeff22 on April 15 2020\n", solution3)
}
