package main

import (
	"github.com/oledakotajoe/codenvi-core/environment"
	"github.com/oledakotajoe/codenvi-core/terminal"
	"github.com/oledakotajoe/codenvi-core/types"
)

func main() {
	body := types.Closure{
		Args: nil,
		Mutator: func(mutator *types.Closure) {
			//terminal.EnviTermialV2()
			terminal.EnviTerminal()
		},
		ReturnValue: nil,
	}

	environment.WithEnv(map[string]string{"HELLO": "hello, world"}, &body)
}
