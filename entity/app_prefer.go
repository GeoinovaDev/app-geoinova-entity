package entity

type AppPrefer struct {
	Id                   uint
	ColunasTabelaCamadas []string
	Cliente              *Cliente
}

func NewAppPrefer(id uint) *AppPrefer {
	return &AppPrefer{
		Id: id,
	}
}
