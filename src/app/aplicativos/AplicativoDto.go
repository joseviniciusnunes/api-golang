package aplicativos

import "time"

type CriarAplicativoDtoRequest struct {
	Id                   *string    `json:"id"`
	Nome                 *string    `json:"nome"`
	ExcluidoEm           *time.Time `json:"excluidoEm"`
	CriadoPorTipo        *int       `json:"criadoPorTipo"`
	CriadoPorMatricula   *int       `json:"criadoPorMatricula"`
	CriadoEm             *time.Time `json:"criadoEm"`
	AlteradoPorTipo      *int       `json:"alteradoPorTipo"`
	AlteradoPorMatricula *int       `json:"alteradoPorMatricula"`
	AlteradoEm           *time.Time `json:"alteradoEm"`
	CredencialFirebase   *string    `json:"credencialFirebase"`
}

func (dto CriarAplicativoDtoRequest) IsValid() error {
	return nil
}
