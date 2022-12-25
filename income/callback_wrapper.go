package income

import (
	"fmt"

	"vendors_sandbox/domain"
)

type CallbackWrapper struct {
	cb Callback
}

func NewCallbackWrapper(cb Callback) *CallbackWrapper {
	return &CallbackWrapper{cb: cb}
}

func (c *CallbackWrapper) HandleRawData() error {
	// Common code
	callbackParams, err := c.cb.HandleRawData()
	if err != nil {
		return err
	}
	// Common code
	switch op := callbackParams.Operation(); op {
	case domain.OperationWin:
		// Common code
		return c.cb.SendWin(callbackParams)
	case domain.OperationJackpot:
		jackpotCb, ok := c.cb.(JackpotCallback)
		if !ok {
			return fmt.Errorf("JackpotCallback is unimplemented in %T", c.cb)
		}
		// Common code
		return jackpotCb.SendJackpot(callbackParams)
	default:
		return fmt.Errorf("unknown operation: %d", op)
	}
}
