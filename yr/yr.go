package yr

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	//"go/scanner"
	conv "github.com/issaalmusawi/funtest/konv"
)

func convertFile() error {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter output file name: ")
	scanner.Scan()
	outputFileName := scanner.Text()

	_, err := os.Stat(outputFileName)
	if err == nil {
		fmt.Print("File already exists. Do you want to overwrite it? (y/n): ")
		scanner.Scan()
		answer := scanner.Text()
		if answer != "y" {
			return nil
		}
	}
	/*
		fmt.Print("Do you want to convert temperatures to Fahrenheit? (y/n): ")
		scanner.Scan()
		answer := scanner.Text()
		var convertToFahrenheit bool
		if answer == "y" {
			convertToFahrenheit = true
		} else {
			convertToFahrenheit = false
		}

		fmt.Print("Do you want to calculate average temperature? (y/n): ")
		scanner.Scan()
		answer = scanner.Text()
		var calculateAverage bool
		if answer == "y" {
			calculateAverage = true
		} else {
			calculateAverage = false
		}*/

	// Open input file
	inputFile, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	// Create output file
	outputFile, err := os.Create("2kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// Create a new buffered writer for output file
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	// Create a new buffered reader for input file
	reader := bufio.NewReader(inputFile)

	// Read and write the first line
	firstLine, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	_, err = writer.WriteString(firstLine)
	if err != nil {
		log.Fatal(err)
	}

	// Read input file line by line
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" { // ignore EOF error
				log.Fatal(err)
			}
			break
		}

		// Trim newline character from line
		line = strings.TrimSpace(line)

		// Split line by semicolon separator
		parts := strings.Split(line, ";")

		// Convert last data from Celsius to Fahrenheit
		temperatureC, err := strconv.ParseFloat(parts[len(parts)-1], 64)
		if err != nil {
			continue // skip over non-float lines
		}
		temperatureF := conv.CelsiusToFahrenheit(temperatureC)

		// Update last data in parts slice
		parts[len(parts)-1] = fmt.Sprintln(math.Round(temperatureF*100) / 100)

		// Join parts slice back into a single string
		newLine := strings.Join(parts, ";")

		// Write new line to output file
		_, err = writer.WriteString(newLine)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Write the new last line to the output file
	lastLine := "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Issa Al-musawi \n"
	_, err = writer.WriteString(lastLine)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
