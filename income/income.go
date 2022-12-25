package income

import (
	"vendors_sandbox/domain"
)

type Callback interface {
	HandleRawData() (CallbackParams, error)
	SendWin(params CallbackParams) error
}

type JackpotCallback interface {
	SendJackpot(params CallbackParams) error
}

type CallbackParams interface {
	Operation() domain.Operation
}
