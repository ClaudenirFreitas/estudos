package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramento = 3
const delay = 2

func main() {

	exibeIntroducao()

	for {

		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando desconhecido")
			os.Exit(-1)
		}

	}

}

func exibeIntroducao() {
	nome := "Freitas"
	versao := 1.2
	fmt.Println("Seja bem vindo Sr.", nome)
	fmt.Println("Este programa esta na versao ", versao)
}

func exibeMenu() {
	fmt.Print("\n\n")
	fmt.Println("1 - iniciar o monitoramento")
	fmt.Println("2 - exibir os logs")
	fmt.Println("0 - sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("Comando selecionado: ", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := buscarSites()

	for i := 0; i < monitoramento; i++ {
		for _, site := range sites {
			monitorar(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Print("\n")
	}

}

func buscarSites() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Erro ao ler o arquivo ", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func monitorar(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Printf("Erro ao monitor o site %s. Mensagem: %s", site, err)
	}

	if resp.StatusCode == 200 {
		fmt.Printf("Site %s carregado com sucesso. \n", site)
	} else {
		fmt.Printf("Site %s esta com problemas. Status code: %d. \n", site, resp.StatusCode)
	}
}
