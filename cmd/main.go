package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"log"
	"meusprojetos/crawler-olx-project/internals"
	"strings"

	"github.com/gocolly/colly"
	_ "github.com/lib/pq"
)

func reiceveKafka() {

	connReceive := internals.ConnectKafka("my-topic-final")

	r := internals.ReaderKafka(connReceive, "my-topic")

	for {
		m, err := r.FetchMessage(context.Background())
		if err != nil {
			break
		}
		//fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
		dadosimovel := internals.Imoveis{}
		error := json.Unmarshal(m.Value, &dadosimovel)
		if error != nil {
			fmt.Println(error)
		}
		c2 := colly.NewCollector()

		c2.OnHTML(`div[class="h3us20-5 jHoWDW"]`, func(e *colly.HTMLElement) {

			var firstHash hash.Hash

			firstHash = sha256.New()
			dados := dadosimovel.LinkDetalhes
			firstHash.Write([]byte(dados))

			hashstr := hex.EncodeToString(firstHash.Sum(nil))
			dadosimovel.Endereco = e.ChildText(`dd[class="sc-ifAKCX sc-1f2ug0x-1 kFBcla"]`)
			dadosimovel.Hashdados = hashstr
			ret, _ := json.Marshal(dadosimovel)

			internals.WriteMessageKafka(connReceive, ret)

		})

		c2.Visit(dadosimovel.LinkDetalhes)
		r.CommitMessages(context.Background(), m)
	}

	r.Close()
}

func main() {
	conn := internals.ConnectKafka("my-topic")

	c := colly.NewCollector()

	log.Println("Starting the crawller....")

	c.OnHTML("ul[id=ad-list]", func(e *colly.HTMLElement) {
		go reiceveKafka()
		//go internals.ListMsgTopics("my-topic-final")

		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			linkdetalhes := el.ChildAttr("a[data-lurker-detail=list_id]", "href")

			valorstr := el.ChildText(`p[class="fnmrjs-16 jqSHIm"]`)
			valorstr = strings.Replace(valorstr, "R$", "", -1)
			valorstr = strings.Replace(valorstr, " ", "", -1)
			descricaostr := el.ChildText(`p[class="jm5s8b-0 jDoirm"]`)
			publicacaostr := el.ChildText(`p[class="fnmrjs-19 eJIIxH"]`)

			if valorstr != "" {

				dadosimovel := internals.Imoveis{
					ValorImovel:      valorstr,
					DescricaoImovel:  descricaostr,
					PublicacaoImovel: publicacaostr,
					LinkDetalhes:     linkdetalhes,
				}
				log.Println(dadosimovel)
				ret, _ := json.Marshal(dadosimovel)

				internals.WriteMessageKafka(conn, ret)

			}

		})
	})

	c.OnHTML(`ul[class="sc-1m4ygug-4 cXxSMf"]`, func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			el.Request.Visit(el.ChildAttr("a", "href"))
		})
	})

	c.Visit("https://pb.olx.com.br/paraiba/joao-pessoa/imoveis/venda")
}
