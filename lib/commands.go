package lib

import (
	"regexp"
	"strings"
)

var lists []ICommand

func NewCommands(cmd *ICommand) {
	lists = append(lists, *cmd)
}

func GetList() []ICommand {
	return lists
}

func Get(c *Event, m *IMessage) {
	var prefix string
	pattern := regexp.MustCompile(`[?!.#]`)
	for _, f := range pattern.FindAllString(m.Command, -1) {
		prefix = f
	}
	for _, cmd := range lists {
		if cmd.After != nil {
			cmd.After(c, m)
		}
		re := regexp.MustCompile(`^` + cmd.Name + `$`)
		if valid := len(re.FindAllString(strings.ReplaceAll(m.Command, prefix, ""), -1)) > 0; valid {
			var cmdWithPref bool
			var cmdWithoutPref bool
			if cmd.IsPrefix && (prefix != "" && strings.HasPrefix(m.Command, prefix)) {
				cmdWithPref = true
			} else {
				cmdWithPref = false
			}

			if !cmd.IsPrefix {
				cmdWithoutPref = true
			} else {
				cmdWithoutPref = false
			}

			if !cmdWithPref && !cmdWithoutPref {
				continue
			}

			//Checking
			if cmd.IsOwner && !m.IsOwner {
				continue
			}

			if cmd.IsMedia && m.Media == nil {
				m.Reply("Please use command with added media!")
				continue
			}

			if cmd.IsQuerry && m.Querry == "" {
				m.Reply("Please use command with added query!")
				continue
			}

			if cmd.IsGroup && !m.IsGroup {
				m.Reply("This command is only for groups!")
				continue
			}

			if cmd.IsPrivate && m.IsGroup {
				m.Reply("This command is only for private!")
				continue
			}

			if (m.IsGroup && cmd.IsAdmin) && !m.IsAdmin {
				m.Reply("This command is for group admins only!")
				continue
			}

			if (m.IsGroup && cmd.IsBotAdmin) && !m.IsBotAdmin {
				m.Reply("Before using this command, please make the bot an admin!")
				continue
			}

			if cmd.IsWaitt {
				m.Reply("The request is being processed!")
			}

			cmd.Exec(c, m)
		}
	}
}
