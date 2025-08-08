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
		fmt.Println("Policz se co:")
		eq := ""
		eq, err := in.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input")
			os.Exit(1)
		}

		if eq == "exit\n" || eq == "exit\r\n" {
			os.Exit(0)
		}

		if eq == "q\n" || eq == "q\r\n" {
			os.Exit(0)
		}

		eq_parsed, err := parser(eq)
		if err != nil {
			fmt.Println("Error parsing input")
			os.Exit(1)
		}

		// for _, k := range eq_parsed {

		// 	fmt.Println(k)
		// }

		fmt.Println(calculate(eq_parsed))

	}
}

// zero = 48, nine = 57
// + = 43
// - = 45
// * =42
// / = 47

func RemoveWhiteSpaces(s string) (string, error) {

	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\r", "")
	fmt.Println("Parsed: ", s)

	return s, nil
}

func parser(s string) ([]string, error) {

	var result []string

	eq, err := RemoveWhiteSpaces(s)

	if err != nil {
		return nil, errors.New("error while removing whitespaces sir")
	}

	if eq[0] < 48 || eq[0] > 57 {
		fmt.Println("Equation has to start with a number")
		return nil, errors.New("has to start with  a number")
	}

	if eq[len(eq)-1] < 48 || eq[len(eq)-1] > 57 {
		fmt.Println("Equation has to end with a number. Yours ended with:", eq[len(eq)-1])
		return nil, errors.New("has to end with  a number")
	}

	k := ""
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

	//TODO: add support for more than two factors to be calculated without checking

	for i, k := range eq {

		switch k {

		case "+":
			a, err := strconv.ParseFloat(eq[i-1], 64)
			if err != nil {

				fmt.Println("Error converting string to number")
			}

			b, err := strconv.ParseFloat(eq[i+1], 64)
			if err != nil {

				fmt.Println("Error converting string to number")
			}

			return a + b, nil

		case "-":
			a, err := strconv.ParseFloat(eq[i-1], 64)
			if err != nil {

				fmt.Println("Error converting string to number")
			}

			b, err := strconv.ParseFloat(eq[i+1], 64)
			if err != nil {

				fmt.Println("Error converting string to number")
			}

			return a - b, nil

		case "*":
			a, err := strconv.ParseFloat(eq[i-1], 64)
			if err != nil {

				fmt.Println("Error converting string to number")
			}

			b, err := strconv.ParseFloat(eq[i+1], 64)
			if err != nil {

				fmt.Println("Error converting string to number")
			}

			return a * b, nil

		case "/":
			a, err := strconv.ParseFloat(eq[i-1], 64)
			if err != nil {

				fmt.Println("Error converting string to number")
			}

			b, err := strconv.ParseFloat(eq[i+1], 64)
			if err != nil {

				fmt.Println("Error converting string to number")
			}

			return a / b, nil

		}
	}
	return 0.0, errors.New("function 'calculate' returned in a way it shouldn't (i think)")
}

func DetermineOperationsOrder([]string) ([]string, error) {

	//TODO write a function here that will sort the equation string into substring that are in the right order and can be calculated

	return nil, nil
}

// func printRunes(s string) {
// 	for _, l := range s {
// 		fmt.Print(l, " ")
// 	}

// }
