package models

import (
	"fmt"

	"github.com/nathanzeras/go-ws-tks/config"
)

type Pedidos struct {
	NomeBeneficiario           string `json:"nomeBeneficiario"`
	Cpf                        string `json:"cpf"`
	NumeroGuiaPrestador        string `json:"numeroGuiaPrestador"`
	CodigoAutorizacao          int    `json:"codigoAutorizacao"`
	SenhaAutorizacao           string `json:"senhaAutorizacao"`
	DataAutorizacao            string `json:"dataAutorizacao"`
	DescricaoGlosa             string `json:"descricaoGlosa"`
	CodigoGlosa                string `json:"codigoGlosa"`
	DataExpiracaoAutorizacao   string `json:"dataExpiracaoAutorizacao"`
	TipoGuia                   string `json:"tipoGuia"`
	NumeroConselhoProfissional string `json:"numeroConselhoProfissional"`
	NomeProfissional           string `json:"nomeProfissional"`
	NumeroCarteira             string `json:"numeroCarteira"`
}

//Função que verifica se já existe um paciente cadastrado, caso contrário, realiza a inserção no Clinux
func CreatePedidos(pedidos *Pedidos) int {
	var cdPedido int

	db := config.ConnectDB()

	if len(pedidos.NumeroGuiaPrestador) > 0 {
		//Inserir pedido
		err := db.QueryRow(`insert into ipasgo_pedidos (ds_beneficiario, ds_cpf, ds_guia_prestador,
			 cd_autorizacao, ds_senha_autorizacao, dt_autorizacao, ds_glosa, cd_glosa,
			 dt_expiracao_autorizacao, ds_tipo_guia, ds_crm, ds_medico, nr_carteira) values
			  ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) returning id`,
			pedidos.NomeBeneficiario, pedidos.Cpf, pedidos.NumeroGuiaPrestador, pedidos.CodigoAutorizacao, pedidos.SenhaAutorizacao,
			pedidos.DataAutorizacao, pedidos.DescricaoGlosa, pedidos.CodigoGlosa, pedidos.DataExpiracaoAutorizacao,
			pedidos.TipoGuia, pedidos.NumeroConselhoProfissional, pedidos.NomeProfissional, pedidos.NumeroCarteira).Scan(&cdPedido)
		if err != nil {
			//handle error
			fmt.Printf("Error during insert pedidos: %v", err)
			//log.Fatalf("Error during insert pedidos: %v", err)
		}
	}

	defer db.Close()

	return cdPedido
}
