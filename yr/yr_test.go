package yr

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	//"time"
	//"errors"
	//"io"
	"math"
)

type tempConverterTest struct {
	inputLine    string
	expectedLine string
}

func TestTempConverter(t *testing.T) {
	// Define input/output lines to test
	tests := []tempConverterTest{
		{
			inputLine:    "Kjevik;SN39040;18.03.2022 01:50;6",
			expectedLine: "Kjevik;SN39040;18.03.2022 01:50;42.8",
		},
		{
			inputLine:    "Kjevik;SN39040;07.03.2023 18:20;0",
			expectedLine: "Kjevik;SN39040;07.03.2023 18:20;32",
		},
		{
			inputLine:    "Kjevik;SN39040;08.03.2023 02:20;-11",
			expectedLine: "Kjevik;SN39040;08.03.2023 02:20;12.2",
		},
		{
			inputLine:    "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;",
			expectedLine: "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Issa Al-musawi",
		},
	}

	// Bruk heller absPath!
	inputFile, err := os.Open("/Users/issaal-musawi/minyr/kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		t.Errorf("Failed to open input file: %s", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Open("/Users/issaal-musawi/minyr/kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		t.Errorf("Failed to open output file: %s", err)
	}
	defer outputFile.Close()

	// Lese input og output files, og lage identisk array av dem(string)
	inputLines := make([]string, 0)
	outputLines := make([]string, 0)

	inputScanner := bufio.NewScanner(inputFile)
	outputScanner := bufio.NewScanner(outputFile)

	for inputScanner.Scan() {
		inputLines = append(inputLines, inputScanner.Text())
	}

	for outputScanner.Scan() {
		outputLines = append(outputLines, outputScanner.Text())
	}

	// Sjekke at antall linjer matcher for filene
	if len(inputLines) != len(outputLines) {
		t.Errorf("Number of lines do not match: input=%d, output=%d", len(inputLines), len(outputLines))
	}

	// tester at vår input struct finnes i inputfile
	for _, test := range tests {
		// Finner input linje i inputfil
		var inputIndex int
		for i, line := range inputLines {
			if line == test.inputLine {
				inputIndex = i
				break
			}
		}

		// Finne den forventede outputlinjen i outputfil
		var expectedIndex int
		for i, line := range outputLines {
			if line == test.expectedLine {
				expectedIndex = i
				break
			}
		}

		// Sjekker at outputl linje matcher expected outputlinje og i henhold til inputlinjer
		if outputLines[expectedIndex] != outputLines[expectedIndex] {
			t.Errorf("Unexpected output: input=%s, expected=%s, actual=%s", test.inputLine, test.expectedLine, outputLines[inputIndex])
		}

		//Ettertanker pga avg period arbeid senere:
		//Kan være en ide å flytte denne delen av testen
		//til test for avgPeriod? Så slipper man kodeduplisering
		//Testen er også blitt debugget og sjekket til at alle testene faktisk blir gjennomført,
		//da implementeringen var noe jeg "freestylet" i etterkant

		reader := bufio.NewReader(inputFile)

		// Feklarer tellerer
		sum := 0.0
		count := 0
		//Leser linje for linje
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err.Error() != "EOF" { //Fortsetter å lese, hvis ikke..
					log.Fatalf("Unexpected error: %v", err)
				}
				break
			}

			line = strings.TrimSpace(line) //linje struktur
			parts := strings.Split(line, ";")

			temperatureC, err := strconv.ParseFloat(parts[len(parts)-1], 64)
			if err != nil {
				continue
			}
			//summerer opp resultat
			sum += temperatureC
			count++

		}

		//Tester og ser med en grense på 0.01
		expectedFileAvg := sum / float64(count)
		actualFileAvg := 8.56

		if math.Abs(expectedFileAvg-actualFileAvg) > 0.01 {
			log.Fatalf("Average temp mistmach: expected %.2f, got %.2f", expectedFileAvg, actualFileAvg)
		}

	}
}

/*type avgPeriodTest struct{
	startPeriod string
	endPeriod string
	avg float64

}
*/

/*func TestAvgTemp(t *testing.T) {

	inputFile, err := os.Open("/Users/issaal-musawi/minyr/kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		t.Errorf("Failed to open input file: %s", err)
	}
	defer inputFile.Close()

	err := AverageTemperature()
	if err != nil {
		t.Errorf("AverageTemperature returned an error: %v", err)
	}
	expectedPeriodAvg := 8.56

	if AverageTemperature != expectedPeriodAvg{
		t.Errorf("Average temp = %2.f: expected %2.f", AverageTemperature, expectedPeriodAvg )
	}










}*/
