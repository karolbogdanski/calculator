package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Witaj w kalkulatorze!!!")

	in := bufio.NewReader(os.Stdin)

	for {
		eq := ""
		eq, err := in.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input")
			os.Exit(1)
		}

		if eq == "exit\n" {
			os.Exit(0)
		}

		eq_parsed, err := parser(eq)
		if err != nil {
			fmt.Println("Error parsing input")
			os.Exit(1)
		}

		for _, k := range eq_parsed {

			fmt.Println(k)
		}

		calculate(eq_parsed)

	}
}

// zero = 48, nine = 57
// + = 43
// - = 45
// * =42
// / = 47

func parser(eq string) ([]string, error) {

	var result []string

	eq = strings.Replace(eq, " ", "", -1)
	eq = strings.Replace(eq, "\n", "", -1)
	fmt.Println("Parsed: ", eq)

	if eq[0] < 48 || eq[0] > 57 {
		fmt.Println("Equation has to start with a number")
		return nil, errors.New("has to start with  a number")
	}

	if eq[len(eq)-1] < 48 || eq[len(eq)-1] > 57 {
		fmt.Println("Equation has to end with a number")
		return nil, errors.New("has to end with  a number")
	}

	k := ""
	//var op string
	for _, ch := range eq {

		if ch >= 48 && ch <= 57 {
			k = k + string(ch)
		} else if ch == 43 || ch == 45 || ch == 42 || ch == 47 {
			result = append(result, k, string(ch))
			k = ""
		} else if ch == 32 || ch == 10 {
			continue
		} else {
			message := fmt.Sprintf("Program detected an unsupported character: \"%s\"!  Your execution is scheduled for tomorrow 8:45.", string(ch))
			fmt.Println(message)
			return nil, errors.New("detected unsupported character")
		}

	}
	result = append(result, k)

	return result, nil
}

func calculate(eq []string) (float64, error) {

	result := 0.0

	for i, k := range eq {

		if k == "+" {
			// result = float64(i-1) + float64(eq[i+1]) //seems you can't convert strting to a float that easly
			a, err := strconv.ParseFloat(eq[i-1], 64)
			if err != nil {

				fmt.Println("Error converting string to number")
			}

			b, err := strconv.ParseFloat(eq[i+1], 64)
			if err != nil {

				fmt.Println("Error converting string to number")
			}

			result = a + b
		} else if k == "-" {
			// result = float64(i-1) + float64(eq[i+1]) //seems you can't convert strting to a float that easly
			a, err := strconv.ParseFloat(eq[i-1], 64)
			if err != nil {

				fmt.Println("Error converting string to number")
			}

			b, err := strconv.ParseFloat(eq[i+1], 64)
			if err != nil {

				fmt.Println("Error converting string to number")
			}

			result = a - b
		} else if k == "*" {
			// result = float64(i-1) + float64(eq[i+1]) //seems you can't convert strting to a float that easly
			a, err := strconv.ParseFloat(eq[i-1], 64)
			if err != nil {

				fmt.Println("Error converting string to number")
			}

			b, err := strconv.ParseFloat(eq[i+1], 64)
			if err != nil {

				fmt.Println("Error converting string to number")
			}

			result = a * b
		} else if k == "/" {
			// result = float64(i-1) + float64(eq[i+1]) //seems you can't convert strting to a float that easly
			a, err := strconv.ParseFloat(eq[i-1], 64)
			if err != nil {

				fmt.Println("Error converting string to number")
			}

			b, err := strconv.ParseFloat(eq[i+1], 64)
			if err != nil {

				fmt.Println("Error converting string to number")
			}

			result = a / b
		}

	}

	fmt.Println("Result: ", result)
	return result, nil
}
