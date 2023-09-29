package entity

const (
	ATIVO_STATUS_ATIVO   = "A"
	ATIVO_STATUS_INATIVO = "I"
)

type Ativo struct {
	Id        uint
	Nome      string
	Wkt       string
	Status    AtivoStatus
	Cliente   *Cliente
	Categoria *AtivoCategoria
}

func NewAtivo(id uint) *Ativo {
	return &Ativo{
		Id:     id,
		Status: ATIVO_STATUS_ATIVO,
	}
}
