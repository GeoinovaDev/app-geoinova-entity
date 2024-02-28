package entity

type LicencaTipo struct {
	Id        uint
	Nome      string
	Descricao string
}

func NewLicencaTipo(id uint) *LicencaTipo {
	return &LicencaTipo{Id: id}
}

func NewLicencaTipoWithNome(id uint, nome string) *LicencaTipo {
	return &LicencaTipo{Id: id, Nome: nome}
}
