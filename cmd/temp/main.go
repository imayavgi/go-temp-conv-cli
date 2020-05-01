package main

//Run testing using go test ./cmd/temp -v -run M1 (M2)

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/imayavgi/go-temp-conv-cli/internal/pkg/conv"
	"github.com/imayavgi/go-temp-conv-cli/internal/pkg/util"
)

var originUnit string
var originValue float64

var shouldConvertAgain string

var err error

var errInvalidArguments = errors.New("Invalid arguments")
var errReadingInput = errors.New("Error reading input")

func main() {

	if len(os.Args) != 2 {
		util.PrintError(errInvalidArguments)
	}

	originUnit = strings.ToUpper(os.Args[1])

	for {
		fmt.Print("What is the current temperature in " + originUnit + " ? ")
		_, err = fmt.Scanln(&originValue)

		if err != nil {
			util.PrintError(errReadingInput)
		}

		if originUnit == "C" {
			conv.ConvertToFahrenheit(originValue)
		} else {
			conv.ConvertToCelsius(originValue)
		}

		fmt.Print("Would you like to convert another temperature ? (y/n) ")
		_, err = fmt.Scanln(&shouldConvertAgain)
		if err != nil {
			util.PrintError(errReadingInput)
		}

		if strings.ToUpper(strings.TrimSpace(shouldConvertAgain)) != "Y" {
			fmt.Println("Good bye!")
			break
		}
	}
}
