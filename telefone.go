package entity

import "strings"

type Telefone struct {
	CodePais   string
	CodeRegiao string
	Numero     string
}

func NewTelefone() *Telefone {
	return &Telefone{}
}

func (t *Telefone) ToNumber() string {
	return t.CodeRegiao + t.Numero
}

func (t *Telefone) ToGlobalNumber() string {
	return t.CodePais + t.CodeRegiao + t.Numero
}

func ParseTelefone(numero string) *Telefone {
	_numero := strings.Replace(numero, "(", "", -1)
	_numero = strings.Replace(_numero, ")", "", -1)
	_numero = strings.Replace(_numero, "+", "", -1)
	_numero = strings.Replace(_numero, "-", "", -1)
	_numero = strings.Replace(_numero, "-", "", -1)
	_numero = strings.Replace(_numero, " ", "", -1)

	switch len(_numero) {
	case 8:
		// 39426655 - 8 digitos
		return &Telefone{
			CodePais:   "",
			CodeRegiao: "",
			Numero:     _numero,
		}
	case 9:
		// 982334440 - 9 digitos
		return &Telefone{
			CodePais:   "",
			CodeRegiao: "",
			Numero:     _numero,
		}
	case 10:
		// 6239426655 - 10 digitos
		return &Telefone{
			CodePais:   "",
			CodeRegiao: _numero[:2],
			Numero:     _numero[2:],
		}
	case 11:
		// 62982334440 - 11 digitos
		return &Telefone{
			CodePais:   "",
			CodeRegiao: _numero[:2],
			Numero:     _numero[2:],
		}
	case 12:
		// 556239426655 - 12 digitos
		return &Telefone{
			CodePais:   _numero[:2],
			CodeRegiao: _numero[2:4],
			Numero:     _numero[4:],
		}
	case 13:
		// 5562982334440 - 13 digitos
		return &Telefone{
			CodePais:   _numero[:2],
			CodeRegiao: _numero[2:4],
			Numero:     _numero[4:],
		}
	case 14:
		// 04162982334440 - 14 digitos
		return &Telefone{
			CodePais:   _numero[3:5],
			CodeRegiao: _numero[5:7],
			Numero:     _numero[7:],
		}
	case 15:
		// 041556239426655 - 15 digitos
		return &Telefone{
			CodePais:   _numero[3:5],
			CodeRegiao: _numero[5:7],
			Numero:     _numero[7:],
		}
	}

	return nil
}
