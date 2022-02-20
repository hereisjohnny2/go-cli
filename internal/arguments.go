package internal

import "regexp"

type UserCommand struct {
	Command   string
	Arguments []string
}

func ArgumentsFilter(commandList []string) UserCommand {
	regexValidator := regexp.MustCompile("(?m)-")
	commandSet := false

	var userCommand UserCommand

	for _, a := range commandList {
		isAMatch := regexValidator.MatchString(a)
		if !isAMatch && !commandSet {
			userCommand.Command = a
			commandSet = true
		} else if isAMatch {
			userCommand.Arguments = append(userCommand.Arguments, a)
		}
	}

	return userCommand
}
