package main

import (
	"fmt"

	"vendors_sandbox/factory"
)

func main() {
	f := factory.NewFactory()

	callback := f.Callback("mancala")
	fmt.Println(callback.HandleRawData())

	callback = f.Callback("winfinity")
	fmt.Println(callback.HandleRawData())
}
