package entity

import "time"

type LicencaBuilder struct {
	licenca *Licenca
}

func NewLicencaBuilder(id uint) *LicencaBuilder {
	return &LicencaBuilder{
		licenca: NewLicenca(id),
	}
}

func (b *LicencaBuilder) WithNome(nome string) *LicencaBuilder {
	b.licenca.Nome = nome
	return b
}

func (b *LicencaBuilder) WithDescricao(descricao string) *LicencaBuilder {
	b.licenca.Descricao = descricao
	return b
}

func (b *LicencaBuilder) WithDataProtocolo(dataProtocolo time.Time) *LicencaBuilder {
	b.licenca.DataProtocolo = dataProtocolo
	return b
}

func (b *LicencaBuilder) WithDataVencimento(dataVencimento time.Time) *LicencaBuilder {
	b.licenca.DataVencimento = dataVencimento
	return b
}

func (b *LicencaBuilder) WithTipo(tipo *LicencaTipo) *LicencaBuilder {
	b.licenca.Tipo = tipo
	return b
}

func (b *LicencaBuilder) WithCamada(camada *Camada) *LicencaBuilder {
	b.licenca.Camada = camada
	return b
}

func (b *LicencaBuilder) WithAtivo(ativo *Ativo) *LicencaBuilder {
	b.licenca.Ativo = ativo
	return b
}

func (b *LicencaBuilder) WithStatus(status LicencaStatus) *LicencaBuilder {
	b.licenca.Status = status
	return b
}

func (b *LicencaBuilder) WithArquivo(arquivo *Arquivo) *LicencaBuilder {
	b.licenca.Arquivo = arquivo
	return b
}

func (b *LicencaBuilder) Build() *Licenca {
	return b.licenca
}
