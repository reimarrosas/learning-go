package calculator

type Calculator struct {
    operands []float64
}

func NewCalculator(operands []float64) *Calculator {
    return &Calculator{operands}
}

func (c *Calculator) Add() float64 {
    sum := c.operands[0]
    rest := c.operands[1:]
    
    for _, n := range rest {
        sum += n
    }

    return sum
}

func (c *Calculator) Sub() float64 {
    diff := c.operands[0]
    rest := c.operands[1:]
    
    for _, n := range rest {
        diff -= n
    }

    return diff
}

func (c *Calculator) Mul() float64 {
    prod := c.operands[0]
    rest := c.operands[1:]
    
    for _, n := range rest {
        prod *= n
    }

    return prod
}

func (c *Calculator) Div() float64 {
    quot := c.operands[0]
    rest := c.operands[1:]
    
    for _, n := range rest {
        quot /= n
    }

    return quot
}
