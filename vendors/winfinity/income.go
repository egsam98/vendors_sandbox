package winfinity

import (
	"fmt"

	"vendors_sandbox/domain"
	"vendors_sandbox/income"
)

type Callback struct{}

func NewCallback() *Callback {
	return &Callback{}
}

type callbackParams struct {
	Op     domain.Operation
	Param1 int
	Param2 string
}

func (c *callbackParams) Operation() domain.Operation {
	return c.Op
}

func (c *Callback) HandleRawData() (income.CallbackParams, error) {
	return &callbackParams{Op: domain.OperationJackpot}, nil
}

func (c Callback) SendWin(p income.CallbackParams) error {
	params, ok := p.(*callbackParams)
	if !ok {
		panic(fmt.Errorf("invalid type of callback params: %T", p))
	}
	fmt.Printf("%+v\n", *params)
	return nil
}

func (c Callback) SendJackpot(p income.CallbackParams) error {
	params, ok := p.(*callbackParams)
	if !ok {
		panic(fmt.Errorf("invalid type of callback params: %T", p))
	}
	fmt.Printf("%+v\n", *params)
	return nil
}
