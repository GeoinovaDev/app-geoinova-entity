package entity

import (
	"time"
)

type Incendio struct {
	Id          uint
	Wkt         string
	AreaHa      float64
	Camada      *Camada
	DateTime    *time.Time
	Titulo      string
	Criticidade string
	Confianca   string
	Descricao   string
	Temperatura int
	TotalFocos  uint
	Satelites   []string
}

func NewIncendio(id uint) *Incendio {
	return &Incendio{
		Id: id,
	}
}
