package models

type Retorno struct {
	Status    bool   `json:"sucesso"`
	Descricao string `json:"descricao"`
}
