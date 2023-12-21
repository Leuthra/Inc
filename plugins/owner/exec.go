package owner

import (
	"fmt"
	"inc/lib"
	"os/exec"
)

func init() {
	lib.NewCommands(&lib.ICommand{
		Name:     `\$`,
		As:       []string{"$"},
		Tags:     "owner",
		IsPrefix: false,
		IsOwner:  true,
		Exec: func(client *lib.Event, m *lib.IMessage) {
			out, err := exec.Command("bash", "-c", m.Querry).Output()
			if err != nil {
				m.Reply(fmt.Sprintf("%v", err))
				return
			}
			m.Reply(string(out))
		},
	})
}
