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
		log.Fatalf("EncodedPolyline - Error open file %v: %v\n", nameFile, err)
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

			locationValueOne := resultLog[1][0:11]
			locationValueTwo := resultLog[1][12:24]

			latitude, _ := strconv.ParseFloat(locationValueOne, 64)
			longitude, _ := strconv.ParseFloat(locationValueTwo, 64)

			coords[i] = append(coords[i], latitude, longitude)
			i++
		}
	}

	if lfsErr := logsFileScanner.Err(); lfsErr != nil {
		log.Fatalf("EncodedPolyline - Error while reading file: %s", lfsErr)
	}

	pLine := polyline.EncodeCoords(coords)

	return pLine
}
