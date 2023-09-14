package bindings

import (
	"encoding/json"
	"log"

	ofctx "github.com/OpenFunction/functions-framework-go/context"
)

type Message struct {
	Numbers []int `json:"numbers"`
}

func HandleCronInput(ctx ofctx.Context, in []byte) (ofctx.Out, error) {
	var greeting []byte
	if in != nil {
		log.Printf("binding - Data: %s", in)
		n := 10
		var numbers []int
		for i := 1; i <= n; i++ {
			numbers = append(numbers, i)
		}
		message := Message{Numbers: numbers}
		greeting, err := json.Marshal(message)
		if err != nil {
			log.Printf("Error: %v\n", err)
			return ctx.ReturnOnInternalError(), err
		}
	} else {
		log.Print("binding - Data: no input provided")
	}

	_, err := ctx.Send("kafka-server", greeting)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return ctx.ReturnOnInternalError(), err
	}

	return ctx.ReturnOnSuccess(), nil
}
