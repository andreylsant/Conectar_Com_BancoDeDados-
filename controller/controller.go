package controller

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/andrylsant/model"
)

var temp = template.Must(template.ParseGlob("template/*.html"))

func Home(w http.ResponseWriter, r *http.Request) {
	produtos := model.BuscarProdutosDoBanco()
	temp.ExecuteTemplate(w, "index", produtos)
}

func New(w http.ResponseWriter, r *http.Request){
	temp.ExecuteTemplate(w, "New", nil)
}

func Inserir(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		name := r.FormValue("name")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil{
			log.Fatal("Error na converção Para Float")
		}

		quantidadeConvertidoParaInt, err := strconv.Atoi(quantidade)
		if err != nil{
			log.Fatal("Error na converção Para Float")
		}

		model.NewCreatedProduct(name, descricao, quantidadeConvertidoParaInt , precoConvertidoParaFloat)

	}
}

func Deletar(w http.ResponseWriter, r *http.Request){
	pegaID := r.URL.Query().Get("id")
	model.DeletarProdutos(pegaID)
}