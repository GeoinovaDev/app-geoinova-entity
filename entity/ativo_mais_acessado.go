package entity

type AtivoMaisAcessado struct {
	Id          uint
	Ativo       *Ativo
	Usuario     *Usuario
	TotalAcesso uint
}

func NewAtivoMaisAcessado(id uint, ativo *Ativo, usuario *Usuario, totalAcesso uint) *AtivoMaisAcessado {
	return &AtivoMaisAcessado{id, ativo, usuario, totalAcesso}
}
