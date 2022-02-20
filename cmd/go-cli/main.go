package main

import (
	"fmt"

	"github.com/hereisjohnny2/go-cli/commands/messages"
	"github.com/hereisjohnny2/go-cli/internal"
)

var commandList = []internal.Command{
	new(messages.Message),
}

func main() {
	if err := internal.CommandInit("go-cli").Start(commandList); err != nil {
		fmt.Println(err.Error())
	}
}
