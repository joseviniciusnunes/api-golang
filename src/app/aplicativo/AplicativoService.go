package aplicativo

import "fmt"

func CriarAplicativo(dto AplicativoDtoRequest) {
	fmt.Println(dto.Nome)
}
