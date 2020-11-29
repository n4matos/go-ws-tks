package models

import (
	"encoding/json"
	"fmt"

	"github.com/nathanzeras/go-ws-tks/config"
)

type Autorizacao struct {
	NumeroguiaPrestador        string `json:"numeroGuiaPrestador"`
	CodigoAutorizacao          int    `json:"codigoAutorizacao"`
	SenhaAutorizacao           string `json:"senhaAutorizacao"`
	DataAutorizacao            string `json:"dataAutorizacao"`
	DescricaoGlosa             string `json:"descricaoGlosa"`
	CodigoGlosa                string `json:"codigoGlosa"`
	DataExpiracaoAutorizacao   string `json:"dataExpiracaoAutorizacao"`
	TipoGuia                   string `json:"tipoGuia"`
	CodigoProcedimento         int    `json:"codigoProcedimento"`
	DescricaoProcedimento      string `json:"descricaoProcedimento"`
	QuantidadeSolicitada       int    `json:"quantidadeSolicitada"`
	QuantidadeAutorizada       int    `json:"quantidadeAutorizada"`
	CodigoStatus               string `json:"codigoStatus"`
	NumeroConselhoProfissional string `json:"numeroConselhoProfissional"`
	NomeProfissional           string `json:"nomeProfissional"`
	NomeBeneficiario           string `json:"nomeBeneficiario"`
	Cpf                        string `json:"cpf"`
	NumeroCarteira             string `json:"numeroCarteira"`
}

type RetornoAutorizacao struct {
	Status    bool   `json:"Status"`
	Descricao string `json:"Descricao"`
}

func SearchAutorizacao(p map[string]string) []byte {

	db := config.ConnectDB()
	defer db.Close()
	var retornoAutorizacao RetornoAutorizacao

	//i := 0
	rows, err := db.Query(`select
	ip.ds_guia_prestador        as numeroGuiaPrestador,
	ip.cd_autorizacao           as codigoAutorizacao,
	ip.ds_senha_autorizacao     as senhaAutorizacao,
	ip.dt_autorizacao           as dataAutorizacao,
	ip.ds_glosa                 as descricaoGlosa,
	ip.cd_glosa                 as codigoGlosa,
	ip.dt_expiracao_autorizacao as dataExpiracaoAutorizacao,
	ie.cd_procedimento          as codigoProcedimento,
	ie.ds_procedimento          as descricaoProcedimento,
	ie.qtd_solicitada           as quantidadeSolicitada,
	ie.qtd_autorizada           as quantidadeAutorizada,
	ie.cd_status                as codigoStatus,
	ip.ds_tipo_guia             as tipoGuia,
	ip.ds_crm                   as numeroConselhoProfissional,
	ip.ds_medico                as nomeProfissional,
	ip.ds_beneficiario          as nomeBeneficiario,
	ip.ds_cpf                   as cpf,
	ip.nr_carteira              as numeroCarteira
from ipasgo_pedidos ip
	  join ipasgo_exames ie on ip.ds_guia_prestador = ie.ds_guia_prestador
 and ie.created_at >= now()::date -1 and ip.ds_cpf = $1 and ie.cd_procedimento = $2
 order by ie.created_at desc limit 1`, p["cpf"], p["procedimento"])

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	//fmt.Println(CdPaciente)

	autorizacao := Autorizacao{}

	for rows.Next() {
		if err := rows.Scan(&autorizacao.NumeroguiaPrestador, &autorizacao.CodigoAutorizacao, &autorizacao.SenhaAutorizacao,
			&autorizacao.DataAutorizacao, &autorizacao.DescricaoGlosa, &autorizacao.CodigoGlosa, &autorizacao.DataExpiracaoAutorizacao,
			&autorizacao.CodigoProcedimento, &autorizacao.DescricaoProcedimento, &autorizacao.QuantidadeSolicitada, &autorizacao.QuantidadeAutorizada,
			&autorizacao.CodigoStatus, &autorizacao.TipoGuia, &autorizacao.NumeroConselhoProfissional, &autorizacao.NomeProfissional,
			&autorizacao.NomeBeneficiario, &autorizacao.Cpf, &autorizacao.NumeroCarteira); err != nil {
			fmt.Println(err)
		}
	}

	if autorizacao.QuantidadeSolicitada > 0 {
		output, err := json.Marshal(autorizacao)
		if err != nil {
			fmt.Println("Error marshalling to json:", err)
		}
		fmt.Print(string(output))
		return output
	} else {
		retornoAutorizacao.Status = false
		retornoAutorizacao.Descricao = "Autorização não encontrada."

		output, err := json.Marshal(retornoAutorizacao)
		if err != nil {
			fmt.Println("Error marshalling to json:", err)
		}
		fmt.Print(string(output))
		return output
	}

	/*if err != nil {
		fmt.Println(err)
		//return patient, i, err
	} else {
		for rows.Next() {
			err = rows.Scan(&autorizacao)
		}*/

	/*_, err = dbr.Load(rows, &autorizacao)

	output, err := json.Marshal(autorizacao)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
	//}*/
	//return output
}
