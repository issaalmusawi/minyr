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

	"time"
	//"io"
	conv "github.com/issaalmusawi/funtest/konv"
)

func ConvertFile() error {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter output file name 'kjevik-temp-fahr-20220318-20230318.csv': ")
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

	outputFile, err := os.Create("kjevik-temp-fahr-20220318-20230318.csv")
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

		// Hvordan linje splittes
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

		//Legge temperatur til sum og inkrementere count
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

/*
funksjon leverer average temp for perioden, men perioden kan ikke være samme dag.
Ikke testet, annet enn at average for hele perioden gir ut samme result som convert
Mangler å evt lage test for å se til at resultat faktisk stemmer.
*/
//ny kommentar: gjort en "cheat" test, med table.csv, resultat av den er avg 4.92(regnet på kalk)
//mens funksjon returnerer 5, dermed er det noe som er riktig, men samtidig feil. Må sjekkes ut nøyere!

func AverageTemperature() error {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter start date (DD-MM-YYYY):")
	scanner.Scan()
	startDateString := scanner.Text()

	fmt.Println("Enter end date (DD-MM-YYYY):")
	scanner.Scan()
	endDateString := scanner.Text()

	startDate, err := time.Parse("02-01-2006", startDateString)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	endDate, err := time.Parse("02-01-2006", endDateString)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	if startDate.After(endDate) {
		log.Printf("Startdate must be before end date")
	} else if startDate.Equal(endDate) {
		log.Fatalf("This program cannot calculate the average temp for the same day")
	}

	inputFile, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	reader := bufio.NewReader(inputFile)

	sum := 0.0
	count := 0

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				log.Fatal(err)
			}
			break //break, og ikke continue som vil resultere i unreachable code
			//fordi if print statement er inne i loopen, og må være utenfor loop
			//fordi loopen vil fortsette å lese siden det er flere linjer i filen
		}

		line = strings.TrimSpace(line)
		parts := strings.Split(line, ";")

		date, err := time.Parse("02.01.2006 15:04", parts[2]) //Hvordan det er formatert i filen
		if err != nil {
			continue
		}
		if date.Before(startDate) || date.After(endDate) {
			continue
		}

		temperatureC, err := strconv.ParseFloat(parts[len(parts)-1], 64)
		if err != nil {
			continue
		}

		sum += temperatureC
		count++

	}

	if count > 0 {
		average := sum / float64(count)
		fmt.Printf("The average temperature for the period %s to %s in C is %.2f\n", startDate.Format("02-01-2006"), endDate.Format("02-01-2006"), average)
	} else {
		fmt.Println("No temperature data found for the period.")
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
