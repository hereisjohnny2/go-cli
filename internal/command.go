package internal

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
)

type Command interface {
	Name() string
	Example() string
	Help() string
	LongHelp() string
	Register(*flag.FlagSet)
	Run()
}

type CommandRoot struct {
	Name     string
	commands []Command
}

func CommandInit(name string) *CommandRoot {
	return &CommandRoot{
		Name: name,
	}
}

func (cr *CommandRoot) Start(commandsList []Command) error {
	if len(commandsList) == 0 {
		return errors.New("command line initialization requires one or more commands")
	}

	cr.commands = commandsList

	if len(os.Args) < 2 {
		cr.showHelp()
		return errors.New("please, pass some command")
	}

	arguments := os.Args[1:]

	userCommand := ArgumentsFilter(arguments)

	if userCommand.Command == "" {
		cr.showHelp()
		return errors.New("please, pass some valid command")
	}

	if userCommand.Command == "help" {
		cr.showHelp()
		return nil
	}

	for _, command := range cr.commands {
		if userCommand.Command == command.Name() {
			flagset := flag.NewFlagSet(command.Name(), flag.ContinueOnError)
			command.Register(flagset)
			flagset.Parse(os.Args[2:])
			command.Run()
			return nil
		}
	}

	cr.showHelp()

	return fmt.Errorf("%s is not a valid command", userCommand.Command)
}

func (cr *CommandRoot) showHelp() {
	fmt.Printf("Usage: %s [COMMAND] [OPTIONS]\n\n", cr.Name)
	tabwriter := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintf(tabwriter, "commands:\n\n")
	for _, command := range cr.commands {
		fmt.Fprintf(tabwriter, "\t- %s\t%s\n", command.Name(), command.Help())
	}

	tabwriter.Flush()

	fmt.Fprintf(tabwriter, "\nexamples:\n")
	for _, command := range cr.commands {
		fmt.Fprintf(tabwriter, "\t%s\n", command.Example())
	}
}
