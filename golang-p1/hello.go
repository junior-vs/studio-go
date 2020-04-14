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

const monitoramentos = 3
const delay = 5

func main() {

	exibiIntroducao()
	for {
		exibiMenu()
		comando := leComando()
		fmt.Println("Comando escolhido é ", comando)

		switch comando {
		case 1:
			monitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Opcao não reconhecida, finalizando programa")
			os.Exit(-1)
		}
	}

}

func exibiIntroducao() {
	nome := "Junior"
	versao := 1.1
	ano := 2020

	fmt.Println("Ola", nome)
	fmt.Println("Versão do programa:", versao)
	fmt.Println("Ano de inicio de desevolvimento:", ano)
}

func exibiMenu() {

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}

func monitoramento() {
	fmt.Println("Iniciando monitoramento...")
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {

		for i, site := range sites {

			resp, err := http.Get(sites[i])
			if resp.StatusCode == 200 {
				fmt.Println(i, ": Site: ", site, " foi carregado com sucesso")
				registraLog(site, true)
			} else {
				fmt.Println(i, ": Site:", site, "está com problemas. Status Code:", resp.StatusCode)
				registraLog(site, false)
			}
			if err != nil {
				fmt.Println("erro encontrado:", err)
			}
		}
		time.Sleep(delay * time.Second)
	}

}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
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

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	arquivo.WriteString("[" + time.Now().Format("02/01/2006 15:04:05:Z07") + "] " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()

}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}
