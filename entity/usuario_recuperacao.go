package entity

import "time"

type UsuarioRecuperacao struct {
	Id        uint
	Usuario   *Usuario
	Hash      string
	CreatedAt time.Time
}

func NewUsuarioRecuperacao(id uint, usuario *Usuario, hash string, createAt time.Time) *UsuarioRecuperacao {
	return &UsuarioRecuperacao{
		Id:        id,
		Usuario:   usuario,
		Hash:      hash,
		CreatedAt: createAt,
	}
}
