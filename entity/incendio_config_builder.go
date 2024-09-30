package entity

type IncendioConfigBuilder struct {
	config *IncendioConfig
}

func NewIncendioConfigBuilder(id uint) *IncendioConfigBuilder {
	return &IncendioConfigBuilder{
		config: NewIncendioConfig(id),
	}
}

func (b *IncendioConfigBuilder) WithNome(nome string) *IncendioConfigBuilder {
	b.config.Nome = nome
	return b
}

func (b *IncendioConfigBuilder) WithRaio(raio float64) *IncendioConfigBuilder {
	b.config.Raio = raio
	return b
}

func (b *IncendioConfigBuilder) WithCliente(cliente *Cliente) *IncendioConfigBuilder {
	b.config.Cliente = cliente
	return b
}

func (b *IncendioConfigBuilder) WithCamada(camada ...*Camada) *IncendioConfigBuilder {
	b.config.Camadas = append(b.config.Camadas, camada...)
	return b
}

func (b *IncendioConfigBuilder) WithContato(contatos *IncendioContato) *IncendioConfigBuilder {
	b.config.Contatos = contatos
	return b
}

func (b *IncendioConfigBuilder) WithUsuario(usuario ...*Usuario) *IncendioConfigBuilder {
	b.config.Usuarios = append(b.config.Usuarios, usuario...)
	return b
}

func (b *IncendioConfigBuilder) Build() *IncendioConfig {
	return b.config
}
