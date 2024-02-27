package entity

type LicencaTipoBuilder struct {
	licencaTipo *LicencaTipo
}

func NewLicencaTipoBuilder(id uint) *LicencaTipoBuilder {
	return &LicencaTipoBuilder{
		licencaTipo: NewLicencaTipo(id),
	}
}

func (b *LicencaTipoBuilder) WithNome(nome string) *LicencaTipoBuilder {
	b.licencaTipo.Nome = nome
	return b
}

func (b *LicencaTipoBuilder) Build() *LicencaTipo {
	return b.licencaTipo
}
