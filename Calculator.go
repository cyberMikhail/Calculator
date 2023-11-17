package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNum = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func calculate(expression string) (string, error) {
	tokens := strings.Split(expression, " ")
	if len(tokens) != 3 {
		return "", errors.New("Ошибка: неверный формат операции")
	}

	num1 := tokens[0]
	operator := tokens[1]
	num2 := tokens[2]

	isArabic1 := isArabic(num1)
	isArabic2 := isArabic(num2)

	if isArabic1 && isArabic2 {
		a, _ := strconv.Atoi(num1)
		b, _ := strconv.Atoi(num2)
		result, err := startArabicOperation(a, b, operator)
		if err != nil {
			return "", err
		}
		return strconv.Itoa(result), nil
	} else if !isArabic1 && !isArabic2 {
		a, err := checkRoman(num1)
		if err != nil {
			return "", err
		}
		b, err := checkRoman(num2)
		if err != nil {
			return "", err
		}
		result, err := startRomanOperation(a, b, operator)
		if err != nil {
			return "", err
		}
		return result, nil
	} else {
		return "", errors.New("Ошибка: используются разные системы счисления")
	}
}

func isArabic(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

func checkRoman(input string) (int, error) {
	value, ok := romanNum[input]
	if !ok {
		return 0, errors.New("Ошибка: неверный ввод римского числа")
	}
	return value, nil
}

func startArabicOperation(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b != 0 {
			return a / b, nil
		} else {
			return 0, errors.New("Ошибка: деление на ноль")
		}
	default:
		return 0, errors.New("Ошибка: неверная операция")
	}
}

func startRomanOperation(a, b int, operator string) (string, error) {
	result, err := startArabicOperation(a, b, operator)
	if err != nil {
		return "", err
	}
	if result <= 0 {
		return "", errors.New("Ошибка: результат римской операции не может быть меньше единицы")
	}
	return convertToRoman(result), nil
}

func convertToRoman(num int) string {
	values := []int{50, 40, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	numerals := []string{"L", "XL", "X", "IX", "VII", "VII", "VI", "V", "IV", "III", "II", "I"}

	result := ""
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			num -= values[i]
			result += numerals[i]
		}
	}

	return result
}

func main() {
	fmt.Println("Введите математическую операцию:")
	fmt.Println("Обратите внимание, что математическая операция [-], [+], [/], [*]  должна быть отделена от чисел пробелами, иначе будет ошибка")
	fmt.Println("Пример IV + V или 2 + 2")
	reader := bufio.NewReader(os.Stdin)
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	result, err := calculate(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Результат:", result)
}
