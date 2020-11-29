package models

import (
	"fmt"
	"strconv"

	"github.com/nathanzeras/go-ws-tks/config"
)

type Pacientes struct {
	IDPaciente       string `gorm:"column:cd_paciente"`
	nomeBeneficiario string `json:"nomeBeneficiario" gorm:"column:ds_paciente"`
	Cpf              string `json:"cpf" gorm:"column:ds_cpf"`
}

//Função que verifica se já existe um paciente cadastrado, caso contrário, realiza a inserção no Clinux
func CreatePatient(patient *Pacientes) (*Pacientes, int, error) {
	var CdPaciente int

	db := config.ConnectDB()

	i := 0
	rows, err := db.Query(`select cd_paciente as CdPaciente
                              from pacientes
							  where ds_cpf = $1 limit 1`, patient.Cpf)

	defer rows.Close()

	fmt.Println(CdPaciente)

	if err != nil {
		fmt.Println("Erro SQL")
		return patient, i, err
	} else {
		for rows.Next() {
			err = rows.Scan(&CdPaciente)
		}
	}
	//i, err = dbr.Load(rows, &patient)

	if CdPaciente == 0 {
		//Cadastrar paciente
		err := db.QueryRow(`insert into pacientes (ds_cpf, ds_paciente) values ($1, $2) returning cd_paciente`,
			patient.Cpf, patient.nomeBeneficiario).Scan(&CdPaciente)
		if err != nil {
			//handle error
		}

	}

	patient.IDPaciente = strconv.Itoa(CdPaciente)

	/*if err != nil {
	} else if i == 0 {
		//Cadastrar paciente
		err := db.QueryRow(`insert into pacientes (ds_cpf, ds_paciente) values ($1, $2) returning cd_paciente`,
			patient.CpfPaciente, patient.NomePaciente).Scan(&CdPaciente)
		if err != nil {
			//handle error
		}

		fmt.Println(CdPaciente)

	}*/

	return patient, i, err

	/* GORM SQL
	row := db.Where(&Pacientes{CpfPaciente: patient.CpfPaciente}).Find(&patient).Row()


	var result Pacientes
	db.Raw("select cd_paciente from pacientes where ds_cpf = ?", patient.CpfPaciente).Scan(&patient)

	fmt.Println(result.IDPaciente)


	if row != nil {
		err := db.Save(patient).Error

		fmt.Println(patient)

		if err != nil {
			log.Fatalf("Error during obj creating: %v", err)
		}
	}*/

	//return cdPaciente
}
