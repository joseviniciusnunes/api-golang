package assuntos

import (
	"github.com/joseviniciusnunes/api-notificacao-golang/src/helpers/Entity"
)

type AssuntoModel struct {
	Id           *int    `gorm:"primary_key;not_null;auto_increment" json:"id"`
	Titulo       *string `json:"titulo"`
	Conteudo     *string `json:"conteudo"`
	Icone        *string `json:"icone"`
	ChaveAcaoApp *string `json:"chaveAcaoApp"`
	Opcional     *bool   `json:"opcional"`
	Topico       *bool   `json:"topico"`
	Entity.BaseModel
}
