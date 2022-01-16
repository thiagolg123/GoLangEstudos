package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {

	for {
		exibeIntro()

		opcao := leCmd("Thiago")

		switch opcao {
		case 1:
			fmt.Println("endereco da memoria", &opcao)
			fmt.Println("Monitar site")
			resp := monitoraSites()
			fmt.Println(resp.StatusCode)
		case 2:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("opcao nao existe")
			os.Exit(-1)
		}
	}
}

func exibeIntro() {
	nome := "Thiago"
	var versao float32 = 1.1
	tipoInt := 1

	fmt.Println("Hello", nome, "tipo int value:", tipoInt)
	fmt.Println("Version", versao, "tipo int", reflect.TypeOf(versao))

	fmt.Println("1 - monitorar site")
	fmt.Println("2 - Sair")
}

func leCmd(nome string) int16 {
	fmt.Println("parametro recebido", nome)
	var opcao int16
	fmt.Scan(&opcao)
	return opcao
}

func monitoraSites() *http.Response {
	fmt.Println("Monitorando...")
	resp, err := http.Get("http://google.com.br")
	if err != nil {
		fmt.Println(err)
	}
	return resp
}
