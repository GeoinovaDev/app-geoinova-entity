package entity

type Usuario struct {
	Id                uint
	Nome              string
	Email             string
	Senha             string
	Descricao         string
	Status            UsuarioStatus
	Cliente           *Cliente
	PermissaoAcesso   []*Ativo
	Papel             UsuarioPapel
	GrupoPermissao    *GrupoPermissao
	Telefone          *Telefone
	RecebeNotificacao bool
}

func NewUsuario(id uint) *Usuario {
	return &Usuario{
		Id:              id,
		PermissaoAcesso: []*Ativo{},
	}
}

func (u *Usuario) SetTelefoneNumero(numero *string) *Usuario {
	if numero == nil {
		return u
	}

	u.Telefone = ParseTelefone(*numero)
	return u
}

func (u *Usuario) GetTelefoneNumero() *string {
	if u.Telefone == nil {
		return nil
	}

	numero := u.Telefone.ToNumber()
	return &numero
}
