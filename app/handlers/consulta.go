package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nathanzeras/go-ws-tks/app/models"
)

func Search(w http.ResponseWriter, r *http.Request) {

	//fmt.Println(mux.Vars(r))

	//fmt.Println(mux.Vars(r)["cpf"])
	//fmt.Println(mux.Vars(r)["procedimento"])

	output := models.SearchAutorizacao(mux.Vars(r))
	//fmt.Println(r.URL.Query())

	fmt.Fprint(w, string(output))
}
