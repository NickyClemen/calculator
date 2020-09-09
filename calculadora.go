package calculator

import (
	"bufio"
	"os"
	"strconv"
)

type Calc struct {}

func (calc) Operate(inputUno string, inputDos string, operator string) int {
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

func ParseInt(input string) int {
	operador, _ := strconv.Atoi(input)
	return operador
}

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}