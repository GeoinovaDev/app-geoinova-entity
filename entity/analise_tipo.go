package entity

type AnaliseTipo string

func (tipo AnaliseTipo) Equals(at AnaliseTipo) bool {
	return string(tipo) == string(at)
}
