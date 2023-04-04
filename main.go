package main

import (
	"bufio"
	"fmt"

	"log"
	//"math"
	"os"
	//"strconv"
	//"strings"

	//"go/scanner"
	//conv "github.com/issaalmusawi/funtest/konv"
	"github.com/issaalmusawi/minyr/yr"
	//"github.com/issaalmusawi/minyr/yr"
)

func main() {

	args := os.Args
	if len(args) != 2 || args[1] != "minyr" {
		fmt.Println("not build? no worries, enter 'minyr' as second cmd line arg")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter 'q' or 'exit' to quit")
	fmt.Println("Enter 'convert' to convert the file and get average temp for the file")
	fmt.Println("Enter 'average year' to calculate the average temperature")
	fmt.Println("Enter text:")
	for scanner.Scan() {
		text := scanner.Text()
		switch text {
		case "q", "exit":
			os.Exit(0)
		case "convert":
			err := yr.ConvertFile()
			if err != nil {
				log.Fatal(err)
			}
			/*	case "average year":
				fmt.Println("Enter Temperature unit (C or F):")
				scanner.Scan()
				unit := scanner.Text() // ny input som scanneren kan jobbe med

				fmt.Println("Enter start date (YYYYMMDD):")
				scanner.Scan()
				startDate := scanner.Text()

				fmt.Println("Enter endt date (YYYYMMDD):")
				scanner.Scan()
				endDate := scanner.Text()

				err := yr.AverageTemperature(unit, startDate, endDate)
				if err != nil {
					fmt.Printf("Error: %v", err)
				}*/
			return
		}
	}
}

/*
if text == "q" || text == "exit" {
			os.Exit(0)
		} else if text == "convert" {
			err := ConvertFile()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} /* else if text == "average" {
			average, err := calculateAverageTemperature()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Printf("Average temperature: %.2f\n", average)
			}
		} else {
			fmt.Println("Invalid input. Try again.")
		}
	}
}

func ConvertFile() error {

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
		}

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

	// Calculate average temperature in Celsius or Fahrenheit
	fmt.Println("Which average temperature would you like to calculate? Enter 'C' for Celsius, 'F' for Fahrenheit:")
	scanner.Scan()
	tempType := scanner.Text()

	var sum float64
	var count int

	// Open the input file again to read the temperatures
	inputFile, err = os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	reader = bufio.NewReader(inputFile)

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

		// Get the temperature value from the last column
		temperature, err := strconv.ParseFloat(parts[len(parts)-1], 64)
		if err != nil {
			continue // skip over non-float lines
		}

		// Convert temperature to Celsius or Fahrenheit
		if tempType == "F" {
			temperature = conv.CelsiusToFahrenheit(temperature)
		}

		// Add temperature to the sum and increment count
		sum += temperature
		count++
	}

	// Calculate average temperature and print the result
	average := sum / float64(count)
	if tempType == "C" {
		fmt.Printf("The average temperature in Celsius is %.2f\n", average)
	} else if tempType == "F" {
		fmt.Printf("The average temperature in Fahrenheit is %.2f\n", average)
	} else {
		fmt.Println("Invalid temperature type entered.")
	}

	return nil
}*/
