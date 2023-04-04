package yr

import (
	"bufio"
	//"encoding/csv"
	//"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	//"time"
	//"io"
	//"go/scanner"
	conv "github.com/issaalmusawi/funtest/konv"
)

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

	// inputfile og lage outputfil
	inputFile, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create("2kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// buffere for input,output, samt writer,Reader
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	reader := bufio.NewReader(inputFile)

	// uendret første linje, read/write
	firstLine, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	_, err = writer.WriteString(firstLine)
	if err != nil {
		log.Fatal(err)
	}

	// Lese inputfil linje for linje
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

		// conv siste data på linje
		temperatureC, err := strconv.ParseFloat(parts[len(parts)-1], 64)
		if err != nil {
			continue // skip over non-float linje
		}
		temperatureF := conv.CelsiusToFahrenheit(temperatureC)

		// oppdatere siste data på linje før det skrives til ny fil
		parts[len(parts)-1] = fmt.Sprintln(math.Round(temperatureF*100) / 100)

		// samle parts slice tilbake til en string seperaret med filformatet
		newLine := strings.Join(parts, ";")

		// Skrive ny linje til fil
		_, err = writer.WriteString(newLine)
		if err != nil {
			log.Fatal(err)
		}
	}

	// lage egen linje som skal skrives som siste linje i outputfil
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

	// åpne fil på nytt
	inputFile, err = os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	reader = bufio.NewReader(inputFile)

	// lese fil linje for linje
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" { // ignore EOF error
				log.Fatal(err)
			}
			break
		}

		line = strings.TrimSpace(line)

		parts := strings.Split(line, ";")

		temperature, err := strconv.ParseFloat(parts[len(parts)-1], 64)
		if err != nil {
			continue // skip over non-float lines
		}

		if tempType == "F" {
			temperature = conv.CelsiusToFahrenheit(temperature)
		}

		// Add temperature to the sum and increment count
		sum += temperature
		count++
	}

	// kalkulere/skrive ut
	average := sum / float64(count)
	if tempType == "C" {
		fmt.Printf("The average temperature in Celsius is %.2f\n", average)
	} else if tempType == "F" {
		fmt.Printf("The average temperature in Fahrenheit is %.2f\n", average)
	} else {
		fmt.Println("Invalid temperature type entered.")
	}

	return nil
}

/*scanner := bufio.NewScanner(os.Stdin)
inputFile, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
if err != nil {
	log.Fatal(err)
}
defer inputFile.Close()


reader := csv.NewReader(inputFile)
reader.Comma = ';'
parts, err := reader.ReadAll()
if err != nil {
	log.Fatal(err)
}//antar at dette er riktig
*/
