package internals

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// ListMsgTopics - Read mensagens all time from topics received params
func ListMsgTopics(topic string) {

	for {
		connKafka := ConnectKafka(topic)
		r := ReaderKafka(connKafka, topic)

		for {
			m, err := r.FetchMessage(context.Background())
			if err != nil {
				break
			}

			dadosimovel := Imoveis{}
			error := json.Unmarshal(m.Value, &dadosimovel)
			if error != nil {
				fmt.Println(error)
			}

			log.Println("[MONITOR] Data complete: ", dadosimovel)
		}
		connKafka.Close()
	}
}
