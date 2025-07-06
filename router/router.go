package router

import (
	"net/http"

	"github.com/andrylsant/controller"
)

func HandelRouter() {
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/new", controller.New)
	http.HandleFunc("/inserir", controller.Inserir)
	http.HandleFunc("/deletar", controller.Deletar)
}
