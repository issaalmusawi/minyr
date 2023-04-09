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
	if len(args) != 1 || args[0] != "./minyr" {
		fmt.Println("not build? no worries. Please 'go build', then enter './minyr' ")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter 'q' or 'exit' to quit")
	fmt.Println("Enter 'convert' to convert the file and get average temp for the file")
	fmt.Println("Enter 'avg period' to calculate the average temperature")
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
		case "avg period":
			/*fmt.Println("Enter Temperature unit (C or F):")
			scanner.Scan()
			unit := scanner.Text() // ny input som scanneren kan jobbe med

			fmt.Println("Enter start date (YYYYMMDD):")
			scanner.Scan()
			startDate := scanner.Text()

			fmt.Println("Enter endt date (YYYYMMDD):")
			scanner.Scan()
			endDate := scanner.Text()
			*/

			err := yr.AverageTemperature()
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
			return
		}
	}
}
