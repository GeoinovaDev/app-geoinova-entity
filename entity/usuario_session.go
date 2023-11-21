package entity

import (
	"strings"
	"time"
)

type UsuarioSessionDecoder interface {
	GetClaimFloat64(string) (float64, bool)
	GetClaimBoolean(string) (bool, bool)
	GetClaimString(string) (string, bool)
}

type UsuarioSession struct {
	CreateAt        int64
	ExpiredAt       int64
	Usuario         *Usuario
	Dominios        []*UsuarioDominio
	ManterConectado bool
}

func NewUsuarioSession(usuario *Usuario, dominios []*UsuarioDominio, manterConectado bool) *UsuarioSession {
	u := &UsuarioSession{
		Usuario:  usuario,
		Dominios: dominios,
	}

	return u
}

func NewUsuarioSessionFromDecoder(decoder UsuarioSessionDecoder) *UsuarioSession {
	u := &UsuarioSession{}

	userId, exist := decoder.GetClaimFloat64("user_id")
	if exist {
		u.Usuario = NewUsuario(uint(userId))
	}

	initAt, exist := decoder.GetClaimFloat64("init_at")
	if exist {
		u.CreateAt = int64(initAt)
	}

	expiredAt, exist := decoder.GetClaimFloat64("exp_at")
	if exist {
		u.ExpiredAt = int64(expiredAt)
	}

	clienteId, exist := decoder.GetClaimFloat64("cliente_id")
	if exist {
		u.Usuario.Cliente = NewCliente(uint(clienteId))
	}

	manterConectado, exist := decoder.GetClaimBoolean("keep_con")
	if exist {
		u.ManterConectado = manterConectado
	}

	dominios, exist := decoder.GetClaimString("dominios")
	if exist {
		_dominios := strings.Split(dominios, ",")
		for _, dominio := range _dominios {
			u.Dominios = append(u.Dominios, NewUsuarioDominioBuilder(0).WithDominio(dominio).Build())
		}
	}

	return u
}

func (u *UsuarioSession) Claims() map[string]interface{} {
	timeStart := time.Now().Unix()
	claims := map[string]interface{}{}

	claims["init_at"] = timeStart
	claims["exp_at"] = timeStart + expireTime(u.ManterConectado)
	claims["keep_con"] = u.ManterConectado

	claims["user_id"] = u.Usuario.Id
	claims["operador_id"] = u.Usuario.Id
	claims["cliente_id"] = u.Usuario.Cliente.Id

	_dominios := []string{}
	for _, dominio := range u.Dominios {
		_dominios = append(_dominios, dominio.Dominio)
	}

	claims["dominios"] = strings.Join(_dominios, ",")

	return claims
}

func (u *UsuarioSession) IsExpired() bool {
	return time.Now().Unix() > int64(u.GetExpiredAt())
}

func (u *UsuarioSession) ExistDominios(dominios ...string) bool {
	for _, d := range u.Dominios {
		for _, dominio := range dominios {
			if d.Dominio == dominio {
				return true
			}
		}
	}

	return false
}

func (u *UsuarioSession) ExistDominio(dominio string) bool {
	dominios := u.Dominios

	for _, d := range dominios {
		if d.Dominio == dominio {
			return true
		}
	}

	return false
}

func (u *UsuarioSession) GetCreatedAt() int64 {
	return u.CreateAt
}

func (u *UsuarioSession) GetExpiredAt() int64 {
	return u.ExpiredAt
}

func (u *UsuarioSession) GetUserId() uint {
	return u.Usuario.Id
}

func expireTime(manterConectado bool) int64 {
	if manterConectado {
		return 3 * 24 * 60 * 60
	}

	return 86400
}
