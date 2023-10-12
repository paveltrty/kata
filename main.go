package main

import (
	"awesomeProject1/roman"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Type int

const (
	arabianType Type = iota
	romanType
	unknown
)

func determineNumberType(s string) (Type, error) {
	isArabian := true
	isRoman := true
	for i := 0; i < len(s); i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			isArabian = false
		}
	}
	for i := 0; i < len(s); i++ {
		if !(s[i] == 'I' || s[i] == 'V' || s[i] == 'X') {
			isRoman = false
		}
	}
	if isArabian {
		return arabianType, nil
	} else if isRoman {
		return romanType, nil
	}
	return unknown, fmt.Errorf("wrong type")
}

func decompose(s string) (string, string, string, error) {
	elems := strings.Split(s, " ")
	if len(elems) != 3 {
		return "", "", "", fmt.Errorf("wrong format of expression")
	}
	return elems[0], elems[1], elems[2], nil
}

func arabianCalculation(a string, b string, operator string) int {
	aD, _ := strconv.Atoi(a)
	bD, _ := strconv.Atoi(b)

	result := 0
	switch operator {
	case "+":
		result = aD + bD
		break
	case "-":
		result = aD - bD
		break
	case "*":
		result = aD * bD
		break
	case "/":
		result = aD / bD
		break
	}

	return result
}

func romanCalculation(a string, b string, operator string) string {
	aD := roman.ToNumber(a)
	bD := roman.ToNumber(b)
	result := 0

	switch operator {
	case "+":
		result = aD + bD
		break
	case "-":
		result = aD - bD
		break
	case "*":
		result = aD * bD
		break
	case "/":
		result = aD / bD
		break
	}

	romanResult := roman.NumberToRoman(result)

	return romanResult
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		expression, _ := reader.ReadString('\n')
		expression = expression[:len(expression)-1]
		a, operator, b, err := decompose(expression)
		if err != nil {
			log.Fatal(err)
		}
		aT, err := determineNumberType(a)
		if err != nil {
			log.Fatal(err)
		}
		bT, err := determineNumberType(b)
		if err != nil {
			log.Fatal(err)
		}
		if aT != bT {
			log.Fatal(fmt.Errorf("wrong number format"))
		}
		if aT == arabianType {
			result := arabianCalculation(a, b, operator)
			fmt.Printf("%s = %d\n", expression, result)
		} else if aT == romanType {
			result := romanCalculation(a, b, operator)
			fmt.Printf("%s = %s\n", expression, result)
		} else {
			log.Fatal(fmt.Errorf("wrong number format"))
		}
	}
}
