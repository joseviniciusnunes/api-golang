package assuntos

import (
	datavalidation "github.com/joseviniciusnunes/api-notificacao-golang/src/helpers/DataValidation"
)

type CriarAssuntoDtoRequest struct {
	Titulo       *string `json:"titulo"`
	Conteudo     *string `json:"conteudo"`
	Icone        *string `json:"icone"`
	ChaveAcaoApp *string `json:"chaveAcaoApp"`
	Opcional     *bool   `json:"opcional"`
	Topico       *bool   `json:"topico"`
}

func (dto CriarAssuntoDtoRequest) IsValid() []error {

	var errs []error
	if err := datavalidation.String("Título").NotNull(true).Min(5).Max(30).Validate(dto.Titulo); err != nil {
		errs = append(errs, err)
	}
	return errs
}

type AlterarAssuntoDtoRequest struct {
	Titulo       *string `json:"titulo"`
	Conteudo     *string `json:"conteudo"`
	Icone        *string `json:"icone"`
	ChaveAcaoApp *string `json:"chaveAcaoApp"`
	Opcional     *bool   `json:"opcional"`
	Topico       *bool   `json:"topico"`
}

func (dto AlterarAssuntoDtoRequest) IsValid() error {
	return nil
}
