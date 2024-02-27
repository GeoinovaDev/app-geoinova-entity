package entity

import "time"

type Licenca struct {
	Id             uint
	Nome           string
	Descricao      string
	DataProtocolo  *time.Time
	DataVencimento *time.Time
	Status         LicencaStatus
	Tipo           *LicencaTipo
	Camada         *Camada
	Ativo          *Ativo
	Arquivo        *Arquivo
}

func NewLicenca(id uint) *Licenca {
	return &Licenca{
		Id: id,
	}
}
