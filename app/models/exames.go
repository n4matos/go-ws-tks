package models

import (
	"fmt"

	"github.com/nathanzeras/go-ws-tks/config"
)

type ProcedimentosSolicitados struct {
	Exames []Exames `json:"procedimentosSolicitados"`
}

type Exames struct {
	NumeroGuiaPrestador   string `json:"numeroGuiaPrestador"`
	CodigoProcedimento    string `json:"codigoProcedimento"`
	DescricaoProcedimento string `json:"descricaoProcedimento"`
	QuantidadeSolicitada  int    `json:"quantidadeSolicitada"`
	QuantidadeAutorizada  int    `json:"quantidadeAutorizada"`
	CodigoStatus          string `json:"codigoStatus"`
	DescricaoGlosa        string `json:"descricaoGlosa"`
	CodigoGlosa           string `json:"codigoGlosa"`
}

//Função que verifica se já existe um paciente cadastrado, caso contrário, realiza a inserção no Clinux
func (pedido *Pedidos) CreateExames(exames ProcedimentosSolicitados) int {
	var cdExame int
	//var err error

	db := config.ConnectDB()

	for _, exame := range exames.Exames {

		err := db.QueryRow(`insert into ipasgo_exames (ds_guia_prestador, cd_procedimento, ds_procedimento, qtd_solicitada,
			qtd_autorizada, cd_status, ds_glosa, cd_glosa) values ($1, $2, $3, $4,
				$5, $6, $7, $8) returning id`,
			pedido.NumeroGuiaPrestador, exame.CodigoProcedimento, exame.DescricaoProcedimento, exame.QuantidadeSolicitada,
			exame.QuantidadeAutorizada, exame.CodigoStatus, exame.DescricaoGlosa, exame.CodigoGlosa).Scan(&cdExame)
		if err != nil {
			//handle error
			fmt.Printf("Error during insert exames: %v", err)
			//log.Fatalf("Error during insert exames: %v", err)
		}
	}

	defer db.Close()

	return cdExame
}
