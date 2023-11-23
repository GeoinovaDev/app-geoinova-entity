package entity

import (
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
	Version         string
	Usuario         *Usuario
	ManterConectado bool
}

func NewUsuarioSession(version string, usuario *Usuario, manterConectado bool) *UsuarioSession {
	u := &UsuarioSession{
		Usuario:         usuario,
		Version:         version,
		ManterConectado: manterConectado,
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

	version, exist := decoder.GetClaimString("version")
	if exist {
		u.Version = version
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

	claims["version"] = u.Version

	return claims
}

func (u *UsuarioSession) IsExpired() bool {
	return time.Now().Unix() > int64(u.GetExpiredAt())
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

func (u *UsuarioSession) GetVersion() string {
	return u.Version
}

func (u *UsuarioSession) GetClienteId() uint {
	if u.Usuario == nil {
		return 0
	}

	if u.Usuario.Cliente == nil {
		return 0
	}

	return u.Usuario.Cliente.Id
}

func (u *UsuarioSession) IsManterConectado() bool {
	return u.ManterConectado
}

func expireTime(manterConectado bool) int64 {
	if manterConectado {
		return 3 * 24 * 60 * 60
	}

	return 86400
}
