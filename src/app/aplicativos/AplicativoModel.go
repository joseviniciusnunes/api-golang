package aplicativos

import (
	"github.com/joseviniciusnunes/api-notificacao-golang/src/helpers/Entity"
)

type AplicativoModel struct {
	Id                 *int    `gorm:"primary_key;not_null;auto_increment" json:"id"`
	Nome               *string `json:"nome"`
	CredencialFirebase *string `json:"-"`
	Entity.BaseModel
}
