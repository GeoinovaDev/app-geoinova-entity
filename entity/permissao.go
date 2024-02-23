package entity

type Permissao struct {
	Id        uint
	Nome      string
	Action    string
	Categoria *PermissaoCategoria
	Leitura   bool
	Escrita   bool
	Exclusao  bool
}

type Permissoes []*Permissao

func NewPermissao(id uint) *Permissao {
	return &Permissao{
		Id: id,
	}
}

func (p *Permissoes) Add(permissao *Permissao) *Permissoes {
	*p = append(*p, permissao)
	return p
}

func (p *Permissoes) Remove(permissao *Permissao) *Permissoes {
	_p := *p
	for i, p := range _p {
		if p.Id == permissao.Id {
			_p = append(_p[:i], _p[i+1:]...)
			break
		}
	}

	return p
}

func (p *Permissoes) Exist(permissao *Permissao) bool {
	for _, p := range *p {
		if p.Id == permissao.Id {
			return true
		}
	}

	return false
}

func (p *Permissoes) Has(nome string) bool {
	for _, permissao := range *p {
		if permissao.Nome == nome {
			return true
		}
	}

	return false
}

func (p *Permissoes) Get(nome string) *Permissao {
	for _, permissao := range *p {
		if permissao.Nome == nome {
			return permissao
		}
	}

	return nil
}
