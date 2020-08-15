# Crawler - OLX - Kafka Connect - GoLang - Colly

##### Projeto desenvolvido com o objetivo de estudar a velocidade de web scrapping/crawler da biblioteca Colly em GoLang. Efetuado captura de informações relevantes dos imoveis de JoãoPessoa - PB e enviado ao Kafka para utilização de qual quer parte da aplicação e por fim utilizado o Kafka Connect como sink para um datasource de escolha propria. Com isso a partir do data source destino dos dados é possivel aplicar uma analise de dados sobre as informações de imoveis de JP ou até mesmo aplicação de modelos de ML.


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