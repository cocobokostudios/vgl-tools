package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/cocobokostudios/vgl-tools/fetchprice/lib"
)

type Game struct {
	GameId       string
	Title        string
	PlatformId   string
	RegionId     string
	Box          string
	Instructions string
	Media        string
	Sealed       string
	Graded       string
}

func main() {
	// read the CSV
	file, err := os.Open("./data/collection.csv")
	if err != nil {
		fmt.Printf("An error occured while opening the file. Error: %s", err.Error())
	}
	defer file.Close()

	// create new file
	newFile, err := os.Create("./data/collection.new.csv")
	if err != nil {
		fmt.Printf("An error occured while creating new file. Error: %s", err.Error())
	}
	defer file.Close()

	// read the values
	csvReader := csv.NewReader(file)
	games, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	// write updated file
	csvWriter := csv.NewWriter(newFile)
	for i, line := range games {
		var game [11]string
		if i > 0 {
			fmt.Println(line)
			for j, field := range line {
				if j == 0 {
					// do nothing
				} else if j == 1 {
					game[0] = lib.NormalizeString(field)
					game[1] = field
				} else if j == 2 {
					game[2] = field
				} else if j == 3 {
					game[3] = field
				} else if j == 4 {
					game[4] = field
				} else if j == 5 {
					game[5] = field
				} else if j == 6 {
					game[6] = field
				} else if j == 7 {
					game[7] = field
				} else if j == 8 {
					game[8] = field
				} else if j == 9 {
					game[9] = field
				} else if j == 10 {
					game[10] = field
				}
			}
			csvWriter.Write(game[:])
		} else { // header line
			csvWriter.Write(line)
		}

		csvWriter.Flush()
	}
}
