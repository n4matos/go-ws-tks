package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/nathanzeras/go-ws-tks/app/models"
)

//Recebimento das autorizações geradas pelo WS da TKS
func Recebimento(w http.ResponseWriter, r *http.Request) {

	// slurp up the body.  p is a []byte
	p, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// handler error
	}

	var pat models.Pacientes

	if err := json.Unmarshal(p, &pat); err != nil {
		// handle error
	} else {
		models.CreatePatient(&pat)
	}

	/*var med models.Medicos
	if err := json.Unmarshal(p, &med); err != nil {
		// handle error
	} else {
		models.CreateMedico(&med)
	}

	var at models.Atendimentos
	if err := json.Unmarshal(p, &at); err != nil {
		// handle error
		fmt.Println(err)
	} else {
		//models.CreateAtendimentos(&at)

	}

	var ex models.Exames
	if err := json.Unmarshal(p, &ex); err != nil {
		// handle error
	} else {

	}*/
}

//Recebimento das autorizações geradas pelo WS da TKS
func RecebimentoGeral(w http.ResponseWriter, r *http.Request) {

	var retorno models.Retorno
	var cdExame int
	var cdPedido int
	var errRetorno string

	// slurp up the body.  p is a []byte
	p, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//log.Fatalf("Error reading request: %v", err)
		fmt.Printf("Error reading request: %v", err)
	}

	var pedidos models.Pedidos

	if err := json.Unmarshal(p, &pedidos); err != nil {
		//log.Fatalf("Error during unmarshall pedidos: %v", err)
		fmt.Printf("Error during unmarshall pedidos: %v", err)
		errRetorno = err.Error()
	} else {
		//Removing all occorruncies from caracter on string
		pedidos.Cpf = strings.Replace(pedidos.Cpf, ".", "", -1)
		pedidos.Cpf = strings.Replace(pedidos.Cpf, "-", "", -1)

		cdPedido = models.CreatePedidos(&pedidos)
	}

	var exames models.ProcedimentosSolicitados
	if err := json.Unmarshal(p, &exames); err != nil {
		//log.Fatalf("Error during unmarshall exames: %v", err)
		fmt.Printf("Error during unmarshall exames: %v", err)
		errRetorno = err.Error()
	} else {
		cdExame = pedidos.CreateExames(exames)
	}

	if cdPedido > 0 && cdExame > 0 {
		retorno.Status = true
		retorno.Descricao = "Autorizacao cadastrada com sucesso"
	} else {
		retorno.Status = false
		retorno.Descricao = errRetorno
	}

	output, err := json.Marshal(&retorno)
	if err != nil {
		fmt.Println("Error marshalling to Json:", err)
	}

	fmt.Fprint(w, string(output))
}
