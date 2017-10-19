package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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
			exibirLogs()
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
	fmt.Print("Informe seu nome: ")
	var nome string
	fmt.Scanln(&nome)

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
		registrarLog(site, true)
	} else {
		fmt.Printf("Site %s esta com problemas. Status code: %d. \n", site, resp.StatusCode)
		registrarLog(site, false)
	}
}

func registrarLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("Erro ao registrar o log para o site %s. Mensagem: %s.", site, err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + " \n")

	arquivo.Close()
}

func exibirLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Printf("Erro ao exibir os logs. Mensagem: %s.", err)
	}

	fmt.Println("exibindo logs...")
	fmt.Println(string(arquivo))

}
