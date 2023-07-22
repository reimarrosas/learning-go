package main

import (
	"beginner/simple-calculator/calculator"
	"flag"
	"log"
	"strconv"
)

func parseFloatArgs(args []string) (*[]float64, error) {
	parsedArgs := make([]float64, len(args))

	for i, s := range args {
		n, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, err
		}
		parsedArgs[i] = n
	}

	return &parsedArgs, nil
}

func main() {
	var op string
	flag.StringVar(&op, "op", "", "The operation to be used: +,-,*,/")
	flag.Parse()
    
    args := flag.Args()
    if len(args) < 2 {
        log.Fatalf("Operands must be more than 2\n")
    }
	operands, err := parseFloatArgs(flag.Args())
	if err != nil {
		log.Fatalf("Cannot parse arguments, Err: %q", err)
	}

	calc := calculator.NewCalculator(*operands)

	res := 0.0
	switch op {
	case "+":
		res = calc.Add()
	case "-":
		res = calc.Sub()
	case "*":
		res = calc.Mul()
	case "/":
		res = calc.Div()
	default:
		log.Fatalf("Invalid operation: %q\n", op)
	}

	log.Printf("Result: %v\n", res)
}
