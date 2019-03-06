package calc

import (
	"fmt"
	"regexp"
	"server/stack"
	"strconv"
)

//Evaluate calcuate math task
func Evaluate(in string) (float64, error) {
	//check if string only contain numbers and operators ( + - / * )
	ok := regexp.MustCompile(`^[0-9,\+,\-,\*,/]+$`).MatchString(in)
	if !ok {
		return 0, fmt.Errorf("Can't parse %v, make sure to send number 0-9 , and + - * / ", in)
	}
	//split string on operator and keep it
	//ex : 1+2*3 => [ 1, +, 2, *, 3]
	matches := regexp.MustCompile(`[0-9']+|[\+\*-/]`).FindAllString(in, -1)
	postfix := ItoP(matches)
	f, err := evalPostfix(postfix)
	if err != nil {
		return 0, err
	}
	return f, nil
}

//ItoP convert Infix Expression to Post-fix Expression
func ItoP(input []string) (postfix []string) {
	s := stack.New()
	for _, i := range input {
		if isOperand(i) {
			postfix = append(postfix, i)
		} else if isOperator(i) {
			for !s.IsEmpty() && hasHigherPrec(i, s.Top()) {
				postfix = append(postfix, s.Top())
				s, _ = s.Pop()
			}
			s = s.Push(i)
		}
	}
	for !s.IsEmpty() {
		postfix = append(postfix, s.Top())
		s, _ = s.Pop()
	}
	return
}

// evalPostfix Evaluate postfix expression
func evalPostfix(postfix []string) (float64, error) {
	s := stack.New()
	for _, i := range postfix {
		if isOperand(i) {
			s = s.Push(i)
		} else if isOperator(i) {
			o1 := s.Top()
			s, _ = s.Pop()
			o2 := s.Top()
			s, _ = s.Pop()
			r, err := calculate(o2, i, o1)
			if err != nil {
				return 0, err
			}
			s = s.Push(strconv.FormatFloat(r, 'G', -1, 64))
		}
	}
	f, err := strconv.ParseFloat(s.Top(), 64)
	return f, err
}
func isOperand(i string) bool {
	return regexp.MustCompile(`^[0-9]+$`).MatchString(i)
}

func isOperator(i string) bool {
	switch i {
	case "+":
		return true
	case "-":
		return true
	case "/":
		return true
	case "*":
		return true
	}
	return false
}

func hasHigherPrec(top, i string) bool {
	if top == "" {
		return false
	}
	if i == top {
		return true
	}
	if top == "+" || top == "-" {
		if i == "*" || i == "/" {
			return true
		}
	}
	if top == "*" && i == "/" {
		return true
	} else if top == "/" && i == "*" {
		return false
	}
	return false
}

func calculate(o1, op, o2 string) (float64, error) {
	r := 0.0
	p1, err := strconv.ParseFloat(o1, 64)
	if err != nil {
		return r, err
	}
	p2, err := strconv.ParseFloat(o2, 64)
	if err != nil {
		return r, err
	}
	switch op {
	case "+":
		r = p1 + p2
	case "-":
		r = p1 - p2
	case "*":
		r = p1 * p2
	case "/":
		r = p1 / p2
	}
	return r, nil
}
