package models

import (
	"fmt"
	"strconv"

	"github.com/nathanzeras/go-ws-tks/config"
)

type Medicos struct {
	IDMedico                   string
	NumeroConselhoProfissional string `json:"numeroConselhoProfissional"`
	NomeProfissional           string `json:"nomeProfissional"`
}

func CreateMedico(medico *Medicos) (*Medicos, int, error) {
	var CdMedico int

	db := config.ConnectDB()

	i := 0
	rows, err := db.Query(`select cd_medico as CdMedico
                              from medicos
							  where ds_crm_nr = $1 limit 1`, medico.NumeroConselhoProfissional)

	defer rows.Close()

	if err != nil {
		fmt.Println("Erro SQL")
		return medico, i, err
	} else {
		for rows.Next() {
			err = rows.Scan(&CdMedico)
		}
	}

	if CdMedico == 0 {
		//Cadastrar paciente
		err := db.QueryRow(`insert into medicos (ds_medico, ds_crm_nr, ds_crm_uf) values ($1, $2, $3) returning cd_medico`,
			medico.NomeProfissional, medico.NumeroConselhoProfissional, "GO").Scan(&CdMedico)
		if err != nil {
			//handle error
		}

	}

	medico.IDMedico = strconv.Itoa(CdMedico)

	return medico, i, err
}
