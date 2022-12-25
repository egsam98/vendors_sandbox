package factory

import (
	"vendors_sandbox/income"
	"vendors_sandbox/vendors/mancala"
	"vendors_sandbox/vendors/winfinity"
)

type Factory struct {
	incomes map[string]income.Callback
}

func NewFactory() *Factory {
	return &Factory{
		incomes: map[string]income.Callback{
			"winfinity": winfinity.NewCallback(),
			"mancala":   mancala.NewCallback(),
		},
	}
}

func (f *Factory) Callback(vendorType string) *income.CallbackWrapper {
	return income.NewCallbackWrapper(f.incomes[vendorType])
}
