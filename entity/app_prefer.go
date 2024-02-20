package entity

type AppPrefer struct {
	ColunasTabelaCamadas []string
	Cliente              *Cliente
}

func NewAppPrefer(cliente *Cliente) *AppPrefer {
	return &AppPrefer{
		Cliente: cliente,
	}
}
