package server

import (
	"fmt"
	"github.com/detecc/detecctor/plugin"
	"github.com/detecc/detecctor/shared"
	"log"
	"strings"
)

func init() {
	example := &Example{}
	plugin.Register(example.GetCmdName(), example)
}

type Example struct {
	plugin.Handler
}

func (e Example) GetCmdName() string {
	return "/example"
}

func (e Example) Response(payload shared.Payload) shared.Reply {
	log.Println(payload)
	content := payload.Data.([]interface{})
	content2 := make([]string, len(content))
	for i, v := range content {
		content2[i] = fmt.Sprint(v)
	}

	return shared.Reply{
		ChatId:    904544573,
		ReplyType: shared.TypeMessage,
		Content:   strings.Join(content2, ", "),
	}
}

func (e Example) Execute(args ...string) ([]shared.Payload , error) {
	log.Println(args)
	return []shared.Payload{
		shared.Payload{
			Id:             "0",
			ServiceNodeKey: "a7309be127cf7g9127309",
			Data:           []string{"test", "test2"},
			Command:        e.GetCmdName(),
			Success:        true,
			Error:          "",
		},
	}, nil
}