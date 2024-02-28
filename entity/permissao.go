package entity

type PermissaoAction string

func (p PermissaoAction) Equals(permissao *Permissao) bool {
	if permissao == nil {
		return false
	}

	if permissao.Action == p {
		return true
	}

	return false
}

const (
	PERMISSAO_VIEW_ALERTAS                = PermissaoAction("view_alertas")
	PERMISSAO_INSERT_ALERTA_JUSTIFICATIVA = PermissaoAction("insert_alerta_justificativa")
	PERMISSAO_EDIT_ALERTA_JUSTIFICATIVA   = PermissaoAction("insert_edit_justificativa")
	PERMISSAO_INSERT_CAMADA               = PermissaoAction("insert_camada")
	PERMISSAO_EDIT_CAMADA                 = PermissaoAction("edit_camada")
	PERMISSAO_DELETE_CAMADA               = PermissaoAction("delete_camada")
	PERMISSAO_INSERT_LICENCA              = PermissaoAction("insert_licenca")
	PERMISSAO_EDIT_LICENCA                = PermissaoAction("edit_licenca")
	PERMISSAO_DELETE_LICENCA              = PermissaoAction("delete_licenca")
	PERMISSAO_VIEW_LICENCA                = PermissaoAction("view_licenca")
	PERMISSAO_INSERT_USUARIO              = PermissaoAction("insert_usuario")
	PERMISSAO_EDIT_USUARIO                = PermissaoAction("edit_usuario")
	PERMISSAO_DELETE_USUARIO              = PermissaoAction("delete_usuario")
	PERMISSAO_VIEW_USUARIOS               = PermissaoAction("view_usuarios")
	PERMISSAO_GRANT_PERMISSAO             = PermissaoAction("grant _permissao")
	PERMISSAO_REVOKE_PERMISSAO            = PermissaoAction("revoke_permissao")
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
