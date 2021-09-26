package aplicativo

import "time"

type AplicativoModel struct {
	Id                   *string
	nome                 *string
	excluidoEm           *time.Time
	criadoPorTipo        *int
	criadoPorMatricula   *int
	criadoEm             *time.Time
	alteradoPorTipo      *int
	alteradoPorMatricula *int
	alteradoEm           *time.Time
	credencialFirebase   *string
}
