package entity

type PermissaoAction string

func (p PermissaoAction) Equals(action PermissaoAction) bool {
	return p == action
}

func (p PermissaoAction) ToString() string {
	return string(p)
}

const (
	PERMISSAO_INSERT_ATIVO = PermissaoAction("insert_ativo")
	PERMISSAO_VIEW_ATIVOS  = PermissaoAction("view_ativos")
	PERMISSAO_EDIT_ATIVO   = PermissaoAction("edit_ativo")
	PERMISSAO_DELETE_ATIVO = PermissaoAction("delete_ativo")

	PERMISSAO_INSERT_ANALISE = PermissaoAction("insert_analise")
	PERMISSAO_VIEW_ANALISES  = PermissaoAction("view_analises")
	PERMISSAO_EDIT_ANALISE   = PermissaoAction("edit_analise")
	PERMISSAO_DELETE_ANALISE = PermissaoAction("delete_analise")

	PERMISSAO_INSERT_IMAGEM = PermissaoAction("insert_imagem")
	PERMISSAO_VIEW_IMAGENS  = PermissaoAction("view_imagens")
	PERMISSAO_EDIT_IMAGEM   = PermissaoAction("edit_imagem")
	PERMISSAO_DELETE_IMAGEM = PermissaoAction("delete_imagem")

	PERMISSAO_INSERT_ALERTA = PermissaoAction("insert_alerta")
	PERMISSAO_VIEW_ALERTAS  = PermissaoAction("view_alertas")
	PERMISSAO_EDIT_ALERTA   = PermissaoAction("edit_alerta")
	PERMISSAO_DELETE_ALERTA = PermissaoAction("delete_alerta")

	PERMISSAO_INSERT_CAMADA = PermissaoAction("insert_camada")
	PERMISSAO_VIEW_CAMADAS  = PermissaoAction("view_camadas")
	PERMISSAO_EDIT_CAMADA   = PermissaoAction("edit_camada")
	PERMISSAO_DELETE_CAMADA = PermissaoAction("delete_camada")

	PERMISSAO_INSERT_LICENCA = PermissaoAction("insert_licenca")
	PERMISSAO_VIEW_LICENCAS  = PermissaoAction("view_licencas")
	PERMISSAO_EDIT_LICENCA   = PermissaoAction("edit_licenca")
	PERMISSAO_DELETE_LICENCA = PermissaoAction("delete_licenca")

	PERMISSAO_INSERT_USUARIO = PermissaoAction("insert_usuario")
	PERMISSAO_VIEW_USUARIOS  = PermissaoAction("view_usuarios")
	PERMISSAO_EDIT_USUARIO   = PermissaoAction("edit_usuario")
	PERMISSAO_DELETE_USUARIO = PermissaoAction("delete_usuario")
	PERMISSAO_CHANGE_SENHA   = PermissaoAction("change_senha")

	PERMISSAO_GRANT_PERMISSAO  = PermissaoAction("grant_permissao")
	PERMISSAO_REVOKE_PERMISSAO = PermissaoAction("revoke_permissao")
)

type Permissao struct {
	Id        uint
	Nome      string
	Action    PermissaoAction
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

func NewPermissaoFromAction(action PermissaoAction) *Permissao {
	return &Permissao{
		Action: action,
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

func (p *Permissoes) HasByNome(nome string) bool {
	for _, permissao := range *p {
		if permissao.Nome == nome {
			return true
		}
	}

	return false
}

func (p *Permissoes) Has(permissao *Permissao) bool {
	for _, permissao := range *p {
		if permissao.Action.Equals(permissao.Action) {
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
