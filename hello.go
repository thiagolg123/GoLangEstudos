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
			monitoraSites(montaArraySites())
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

func monitoraSites(sites []string) {
	fmt.Println("Monitorando sites...")

	for _, site := range sites {
		resp, err := http.Get(site)

		if err != nil {
			fmt.Println(err)
		}

		if resp.StatusCode == 200 {
			fmt.Println("Site:", site, "working...")
		} else {
			fmt.Println("Site:", site, "com problemas, codigo de resposta:", resp.StatusCode)
		}
	}

	// for tradicional
	// for i := 0; i < len(sites); i++ {
	// 	site := sites[i]
	// }
}

func montaArraySites() []string {
	sites := []string{"https://www.google.com", "https://www.alura.com.br"}
	//atencao a = nesse ponto para pop/modify variaveis
	sites = append(sites, "https://aws.amazon.com/tp")

	//capacidade do slice
	fmt.Println("Capacidade do slice:", cap(sites))
	return sites
}
