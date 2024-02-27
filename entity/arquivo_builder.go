package entity

type ArquivoBuilder struct {
	arquivo *Arquivo
}

func NewArquivoBuilder(arquivo *Arquivo) *ArquivoBuilder {
	return &ArquivoBuilder{
		arquivo: arquivo,
	}
}

func (b *ArquivoBuilder) WithNome(nome string) *ArquivoBuilder {
	b.arquivo.Nome = nome
	return b
}

func (b *ArquivoBuilder) WithPath(path string) *ArquivoBuilder {
	b.arquivo.Path = path
	return b
}

func (b *ArquivoBuilder) WithTipo(tipo string) *ArquivoBuilder {
	b.arquivo.Tipo = tipo
	return b
}

func (b *ArquivoBuilder) WithSize(size uint) *ArquivoBuilder {
	b.arquivo.Size = size
	return b
}

func (b *ArquivoBuilder) WithSource(source string) *ArquivoBuilder {
	b.arquivo.Source = source
	return b
}

func (b *ArquivoBuilder) WithBucket(bucket string) *ArquivoBuilder {
	b.arquivo.Bucket = bucket
	return b
}

func (b *ArquivoBuilder) Build() *Arquivo {
	return b.arquivo
}
