# Go CLI Project Template

This repository is a template to create a new project in go with an Command Line Interface. It uses the Flags feature in the GoLang and it's supposed to make easier the creation of a complete CLI with simple commands. 

## Commands

New commands should be add at the `commands` folder, on project root, inside it's own package. The `message.go` is a base code to help creating new commands.

Commands should follow the `Command` interface: 

```go
type Command interface {
	Name() string
	Example() string
	Help() string
	LongHelp() string
	Register(*flag.FlagSet)
	Run()
}
```

## Build and Run

```bash
# Build the project
$ make

# Run the binary
$ build/bin/go-cli [COMMAND] [OPTIONS]

# To run the tests
$ make test
```