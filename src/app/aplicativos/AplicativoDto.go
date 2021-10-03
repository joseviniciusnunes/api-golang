package aplicativos

type CriarAplicativoDtoRequest struct {
	Nome               *string `json:"nome"`
	CredencialFirebase *string `json:"credencialFirebase"`
}

func (dto CriarAplicativoDtoRequest) IsValid() error {
	return nil
}

type AlterarAplicativoDtoRequest struct {
	Nome               *string `json:"nome"`
	CredencialFirebase *string `json:"credencialFirebase"`
}

func (dto AlterarAplicativoDtoRequest) IsValid() error {
	return nil
}
