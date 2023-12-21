package owner

import (
	"encoding/json"
	"inc/lib"

	"github.com/robertkrimen/otto"
)

func init() {
	lib.NewCommands(&lib.ICommand{
		Name:     `>`,
		As:       []string{">"},
		Tags:     "owner",
		IsPrefix: false,
		Exec: func(client *lib.Event, m *lib.IMessage) {
			vm := otto.New()
			vm.Set("M", m)

			h, err := vm.Run(m.Querry)
			if err != nil {
				m.Reply(err.Error())
				return
			}

			if h.IsObject() {
				var data interface{}
				h, _ := vm.Run("JSON.stringify(" + m.Querry + ")")
				json.Unmarshal([]byte(h.String()), &data)
				pe, _ := json.MarshalIndent(data, "", "  ")
				m.Reply(string(pe))
			} else {
				m.Reply(h.String())
			}
		},
	})
}
