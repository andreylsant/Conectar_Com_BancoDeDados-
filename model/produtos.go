package model

import (
	"log"

	"github.com/andrylsant/db"
)

type Produtos struct{
	ID int 
	Name string
	Descricao string
	Quantidade int
	Preco float64
}

func BuscarProdutosDoBanco() []Produtos {
	db := db.ConectaComBanco()

	selecionaDadosDoBanco, err := db.Query("select * from Produtos")
	if err != nil{
		log.Fatal(err)
	}

	p := Produtos{}
	produtos := []Produtos{}

	for selecionaDadosDoBanco.Next(){
		var id, quantidade int
		var name, descricao string
		var preco float64

		selecionaDadosDoBanco.Scan(&id, &name, &descricao, &preco, &quantidade)

		p.ID = id
		p.Name = name
		p.Preco = preco
		p.Descricao = descricao
		p.Quantidade = quantidade

		produtos = append(produtos, p)	
	}

	defer db.Close()
	return produtos
}

func NewCreatedProduct(name, descricao string, quantidade int, preco float64){
	db := db.ConectaComBanco()

	inserirProdutosNoBanco, err := db.Prepare("insert infor(name, descricao, quantidade, preco) value($1, $2, $3, $4)")
	if err != nil{
		log.Fatal(err.Error(), "Error ao inserir dados no banco")
	}

	inserirProdutosNoBanco.Exec(name, descricao, quantidade, preco)

	defer db.Close()
}

func DeletarProdutos(id string){
	db := db.ConectaComBanco()

	deletarProdutos, err := db.Prepare("")
	if err != nil{
		log.Fatal(err.Error())
	}

	deletarProdutos.Exec(id)

	defer db.Close()
}