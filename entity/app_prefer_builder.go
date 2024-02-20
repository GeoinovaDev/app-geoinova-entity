package entity

type AppPreferBuilder struct {
	appPrefer *AppPrefer
}

func NewAppPreferBuilder(id uint) *AppPreferBuilder {
	return &AppPreferBuilder{
		appPrefer: NewAppPrefer(id),
	}
}

func (b *AppPreferBuilder) WithColunasTabelaCamadas(cols []string) *AppPreferBuilder {
	b.appPrefer.ColunasTabelaCamadas = cols
	return b
}

func (b *AppPreferBuilder) AddColunaTabelaCamadas(col string) *AppPreferBuilder {
	b.appPrefer.ColunasTabelaCamadas = append(b.appPrefer.ColunasTabelaCamadas, col)
	return b
}

func (b *AppPreferBuilder) WithCliente(cliente *Cliente) *AppPreferBuilder {
	b.appPrefer.Cliente = cliente
	return b
}

func (b *AppPreferBuilder) Build() *AppPrefer {
	return b.appPrefer
}
