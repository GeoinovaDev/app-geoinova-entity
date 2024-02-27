package entity

type LicencaStatus string

const (
	LICENCA_STATUS_ATIVA   = LicencaStatus("A")
	LICENCA_STATUS_INATIVA = LicencaStatus("I")
)

func (u LicencaStatus) Equals(status LicencaStatus) bool {
	return string(u) == string(status)
}

func NewLicencaStatus(status string) LicencaStatus {
	return LicencaStatus(status)
}
