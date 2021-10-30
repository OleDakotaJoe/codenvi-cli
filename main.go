package main

import (
	core "github.com/oledakotajoe/codenvi-core"
	"github.com/oledakotajoe/codenvi-core/environment"
	"github.com/oledakotajoe/codenvi-core/terminal"
	"github.com/oledakotajoe/codenvi-core/types"
	"os"
)

func main() {
	core.Config()
	body := types.Closure{
		Args:        nil,
		Mutator: func(mutator *types.Closure) {
			print(os.Getenv("HELLO"))
			print("ran")
			terminal.EnviTerminal()
		},
		ReturnValue: nil,
	}

	environment.WithEnv(map[string]string{"HELLO": "hello, world"}, &body)
}