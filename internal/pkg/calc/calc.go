package calc

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var operators = map[rune]int{
	'+': 1,
	'-': 1,
	'*': 2,
	'/': 2,
}

var leftAssociative = map[rune]bool{
	'+': true,
	'-': true,
	'*': true,
	'/': true,
}

func precedence(op rune) int {
	return operators[op]
}

func isOperator(c rune) bool {
	_, exists := operators[c]
	return exists
}

func toPostfix(expression string) (string, error) {
	var output []string
	var stack []rune
	expression = correctExpression(spaceStringsBuilder(expression))
	for i := 0; i < len(expression); i++ {
		c := rune(expression[i])

		if unicode.IsDigit(c) {
			num := strings.Builder{}
			num.WriteRune(c)
			for i+1 < len(expression) && (unicode.IsDigit(rune(expression[i+1])) || rune(expression[i+1]) == '.') {
				i++
				num.WriteRune(rune(expression[i]))
			}
			output = append(output, num.String())
		} else if c == '(' {
			stack = append(stack, c)
		} else if c == ')' {
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				output = append(output, string(stack[len(stack)-1]))
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return "", errors.New("mismatched parentheses")
			}
			stack = stack[:len(stack)-1]
		} else if isOperator(c) {
			for len(stack) > 0 && isOperator(stack[len(stack)-1]) &&
				((leftAssociative[c] && precedence(c) <= precedence(stack[len(stack)-1])) ||
					(!leftAssociative[c] && precedence(c) < precedence(stack[len(stack)-1]))) {
				output = append(output, string(stack[len(stack)-1]))
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, c)
		} else if !unicode.IsSpace(c) {
			return "", errors.New("invalid character in expression")
		}
	}

	for len(stack) > 0 {
		if stack[len(stack)-1] == '(' || stack[len(stack)-1] == ')' {
			return "", errors.New("mismatched parentheses")
		}
		output = append(output, string(stack[len(stack)-1]))
		stack = stack[:len(stack)-1]
	}

	return strings.Join(output, " "), nil
}

func evalPostfix(postfix string) (float64, error) {
	var stack []float64
	tokens := strings.Fields(postfix)

	for _, token := range tokens {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, num)
		} else if len(token) == 1 && isOperator(rune(token[0])) {
			if len(stack) < 2 {
				return 0, errors.New("insufficient operands")
			}
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			var result float64
			switch token[0] {
			case '+':
				result = a + b
			case '-':
				result = a - b
			case '*':
				result = a * b
			case '/':
				if b == 0 {
					return 0, errors.New("division by zero")
				}
				result = a / b
			default:
				return 0, errors.New("invalid operator")
			}
			stack = append(stack, result)
		} else {
			return 0, errors.New("invalid token in expression")
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}

	return stack[0], nil
}

func spaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func correctExpression(expression string) string {
	if len(expression) > 0 && expression[0] == '-' {
		expression = "0" + expression
	}
	expression = strings.ReplaceAll(expression, "(-", "(0-")
	return expression
}

func Calc(expression string) (float64, error) {
	postfix, err := toPostfix(expression)
	if err != nil {
		return 0, err
	}

	result, err := evalPostfix(postfix)
	if err != nil {
		return 0, err
	}

	return result, nil
}
