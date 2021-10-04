package Entity

import (
	"time"
)

type BaseModel struct {
	CriadoEm             *time.Time `gorm:"autoCreateTime" json:"criadoEm"`
	CriadoPorTipo        *int       `json:"criadoPorTipo"`
	CriadoPorMatricula   *int       `json:"criadoPorMatricula"`
	AlteradoEm           *time.Time `gorm:"autoUpdateTime" json:"alteradoEm"`
	AlteradoPorTipo      *int       `json:"alteradoPorTipo"`
	AlteradoPorMatricula *int       `json:"alteradoPorMatricula"`
	ExcluidoEm           *time.Time `json:"excluidoEm" sql:"index"`
}
