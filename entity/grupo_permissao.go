package entity

type GrupoPermissao struct {
	Id        uint
	Nome      string
	Permissao []*Permissao
}

func NewGrupoPermissao(id uint) *GrupoPermissao {
	return &GrupoPermissao{
		Id:        id,
		Permissao: []*Permissao{},
	}
}

func (g *GrupoPermissao) Add(permissao *Permissao) *GrupoPermissao {
	g.Permissao = append(g.Permissao, permissao)
	return g
}

func (g *GrupoPermissao) Remove(permissao *Permissao) *GrupoPermissao {
	_p := []*Permissao{}

	for _, p := range g.Permissao {
		if p.Id == permissao.Id {
			continue
		}
		
		_p = append(_p, p)
	}
	g.Permissao = _p

	return g
}

func (g *GrupoPermissao) Exist(permissao *Permissao) bool {
	for _, p := range g.Permissao {
		if p.Id == permissao.Id {
			return true
		}
	}

	return false
}

func (g *GrupoPermissao) Get(nome string) *Permissao {
	for _, p := range g.Permissao {
		if p.Nome == nome {
			return p
		}
	}

	return nil
}

func (g *GrupoPermissao) Has(nome string) bool {
	for _, p := range g.Permissao {
		if p.Nome == nome {
			return true
		}
	}

	return false
}
