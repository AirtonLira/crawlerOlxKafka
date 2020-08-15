# Crawler - OLX - Kafka Connect - GoLang - Colly

##### Projeto desenvolvido com o objetivo de estudar a velocidade de web scrapping/crawler da biblioteca Colly em GoLang. Efetuado captura de informações relevantes dos imoveis de JoãoPessoa - PB e enviado ao Kafka para utilização de qual quer parte da aplicação e por fim utilizado o Kafka Connect como sink para um datasource de escolha propria. Com isso a partir do data source destino dos dados é possivel aplicar uma analise de dados sobre as informações de imoveis de JP ou até mesmo aplicação de modelos de ML. O objetivo desse projeto foi aprimorar meus conhecimentos em GoLang bem como entender a biblioteca colly.

![Desenho do fluxo](https://i.ibb.co/4FksF3K/desenho.jpg)

> Exemplo da mensagem final do topico de sink do Kafka connect:
<code> {"ValorImovel":"225.000","DescricaoImovel":"2 quartos | 61m² | 1 vaga","PublicacaoImovel":"Hoje11:14","LinkDetalhes":"https://pb.olx.com.br/paraiba/imoveis/>excelente-empreendimento-02-e-03-quartos-na-torre-718100990","Endereco":"58040295","Hashdados":"229b4e58a1af4b31f3cf4c6abb8323410395da07169230b52e6c8c0b4e0a2fbf"} </code>


- Estrutura do projeto: 
- cmd 
- main.go 

- internals
- configuratio.go
- db.go      
- kafka.go
- monitor.go

- resources
- application.yml

##### Exemplo de request no Kafka connect para geração do arquivo json a partir do topico final:
![GitHub Logo](https://i.ibb.co/H7s8LJm/postman.png)