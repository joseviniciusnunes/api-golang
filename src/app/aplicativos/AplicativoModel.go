package aplicativos

import (
	"time"
)

type AplicativoModel struct {
	Id                   *int       `gorm:"primary_key;not_null;auto_increment" json:"id"`
	Nome                 *string    `json:"nome"`
	ExcluidoEm           *time.Time `json:"excluidoEm"`
	CriadoPorTipo        *int       `json:"criadoPorTipo"`
	CriadoPorMatricula   *int       `json:"criadoPorMatricula"`
	CriadoEm             *time.Time `gorm:"autoCreateTime" json:"criadoEm"`
	AlteradoPorTipo      *int       `json:"alteradoPorTipo"`
	AlteradoPorMatricula *int       `json:"alteradoPorMatricula"`
	AlteradoEm           *time.Time `gorm:"autoUpdateTime" json:"alteradoEm"`
	CredencialFirebase   *string    `json:"credencialFirebase"`
}
