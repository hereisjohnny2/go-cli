package messages

import (
	"flag"
	"fmt"
)

type Message struct {
	Text  string
	Helpf bool
}

const helpText = "Prints out a simple text message"

const longHelpText = `
Prints out a simple text message

go-cli message --text [printText]
- printText: string
`
const exampleText = `
	go-cli message --text hello-mon
`

func (cmd *Message) Name() string {
	return "message"
}

func (cmd *Message) Help() string {
	return helpText
}

func (cmd *Message) LongHelp() string {
	return longHelpText
}

func (cmd *Message) Example() string {
	return exampleText
}

func (cmd *Message) Register(fs *flag.FlagSet) {
	fs.StringVar(&cmd.Text, "text", "", "printed text")
	fs.BoolVar(&cmd.Helpf, "help", false, "help command")
}

func (cmd *Message) Run() {
	if cmd.Helpf {
		fmt.Println(cmd.LongHelp())
		return
	}

	if cmd.Text == "" {
		fmt.Println("[--text] is required")
		return
	}

	fmt.Printf("[LOG] %s\n", cmd.Text)
}
