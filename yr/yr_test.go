package yr

import (
	"bufio"
	"os"
	"testing"
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
	}

	// Open input and output files
	inputFile, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		t.Errorf("Failed to open input file: %s", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Open("2kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		t.Errorf("Failed to open output file: %s", err)
	}
	defer outputFile.Close()

	// Read lines from input and output files
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

	// Check that number of lines match
	if len(inputLines) != len(outputLines) {
		t.Errorf("Number of lines do not match: input=%d, output=%d", len(inputLines), len(outputLines))
	}

	// Test each input/output pair
	for _, test := range tests {
		// Find the input line in input file
		var inputIndex int
		for i, line := range inputLines {
			if line == test.inputLine {
				inputIndex = i
				break
			}
		}

		// Find the expected output line in output file
		var expectedIndex int
		for i, line := range outputLines {
			if line == test.expectedLine {
				expectedIndex = i
				break
			}
		}

		// Check that the corresponding output line matches the expected output line
		if outputLines[inputIndex] != outputLines[expectedIndex] {
			t.Errorf("Unexpected output: input=%s, expected=%s, actual=%s", test.inputLine, test.expectedLine, outputLines[inputIndex])
		}
	}
}
