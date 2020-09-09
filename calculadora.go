package calculator

import (
	"bufio"
	"os"
	"strconv"
)

type calc struct {}

func (calc) operate(inputUno string, inputDos string, operator string) int {
	operadorUno := parseInt(inputUno)
	operadorDos := parseInt(inputDos)

	switch(operator) {
		case "+":
			return operadorUno + operadorDos
		case "-":
			return operadorUno - operadorDos
		case "*":
			return operadorUno * operadorDos
		case "/":
			return operadorUno / operadorDos
		default:
			return 0
	}
}

func parseInt(input string) int {
	operador, _ := strconv.Atoi(input)
	return operador
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}