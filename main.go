package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/twpayne/go-polyline"
)

func main() {
	polyLine := EncodedPolyline()

	fmt.Printf("PolyLine ==> %s\n", polyLine)
}

func EncodedPolyline() []byte {
	coords := make([][]float64, 0)
	nameFile := "logs.txt"

	// Open the file logs.txt
	logsFile, err := os.Open(nameFile)

	if err != nil {
		log.Fatalf("EncodedPolyline - Error open file %s: %v\n", nameFile, err)
	}

	defer logsFile.Close()

	logsFileScanner := bufio.NewScanner(logsFile)

	var i int

	// Read the file logs.txt line by line
	for logsFileScanner.Scan() {
		log := logsFileScanner.Text()
		containedString := "arrived <+"

		if strings.Contains(log, containedString) {
			resultLog := strings.Split(log, containedString)
			coords = append(coords, make([]float64, 0))

			rlSplit := resultLog[1]
			latitude, _ := strconv.ParseFloat(rlSplit[0:11], 64)
			longitude, _ := strconv.ParseFloat(rlSplit[12:24], 64)

			coords[i] = append(coords[i], latitude)
			coords[i] = append(coords[i], longitude)
			i++
		}
	}

	if lfsErr := logsFileScanner.Err(); lfsErr != nil {
		log.Fatalf("EncodedPolyline - Error while reading file: %s", lfsErr)
	}

	pLine := polyline.EncodeCoords(coords)

	return pLine
}
