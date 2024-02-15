package entity

const (
	FORMULARIO_PROPRIEDADE_TEXT        = "TX"
	FORMULARIO_PROPRIEDADE_SELECT      = "SL"
	FORMULARIO_PROPRIEDADE_MULTISELECT = "MS"
	FORMULARIO_PROPRIEDADE_RADIO       = "RD"
	FORMULARIO_PROPRIEDADE_FLOAT       = "FL"
	FORMULARIO_PROPRIEDADE_INTEGER     = "IN"
	FORMULARIO_PROPRIEDADE_DATE        = "DT"
	FORMULARIO_PROPRIEDADE_DATETIME    = "DE"
	FORMULARIO_PROPRIEDADE_BOOLEAN     = "BO"
)

type FormularioPropriedadeTipo string

func (t FormularioPropriedadeTipo) Equals(tipo FormularioPropriedadeTipo) bool {
	return string(t) == string(tipo)
}

type FormularioPropriedade struct {
	Id    uint
	Nome  string
	Field string
	Tipo  FormularioPropriedadeTipo
}

func NewFormularioPropriedade(id uint) *FormularioPropriedade {
	return &FormularioPropriedade{
		Id:   id,
		Tipo: FORMULARIO_PROPRIEDADE_TEXT,
	}
}
