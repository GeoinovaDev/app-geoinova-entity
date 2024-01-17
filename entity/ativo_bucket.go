package entity

type AtivoBucket struct {
	Id        uint
	Ativo     *Ativo
	Bucket    string
	IsDefault bool
}

func NewAtivoBucket(id uint) *AtivoBucket {
	return &AtivoBucket{
		Id: id,
	}
}
