package internals

//Imoveis referente as informações relavantes a serem enviadas ao kafka
type Imoveis struct {
	ValorImovel      string
	DescricaoImovel  string
	PublicacaoImovel string
	LinkDetalhes     string
	Endereco         string
	Hashdados        string
}
