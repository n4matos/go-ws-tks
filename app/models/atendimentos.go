package models

type Atendimentos struct {
	NumeroGuiaPrestador      string `json:"numeroGuiaPrestador"`
	CodigoAutorizacao        int    `json:"codigoAutorizacao"`
	SenhaAutorizacao         string `json:"senhaAutorizacao"`
	DataAutorizacao          string `json:"dataAutorizacao"`
	DescricaoGlosa           string `json:"descricaoGlosa"`
	CodigoGlosa              string `json:"codigoGlosa"`
	DataExpiracaoAutorizacao string `json:"dataExpiracaoAutorizacao"`
	TipoGuia                 string `json:"tipoGuia"`
}

//Função que realiza a inserção de um novo atendimento no Clinux
func CreateAtendimentos(atendimentos *Atendimentos) (*Atendimentos, int, error) {

	i := 0
	var err error

	//db := config.ConnectDB()

	//err := db.QueryRow(`insert into atendimentos (cd_sala, cd_medico, cd_paciente, dt_data, `)

	return atendimentos, i, err
}
