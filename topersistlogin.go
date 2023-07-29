package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/go-vgo/robotgo"
)

var option = 999

func clearCLI() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func toOnMsg() {
	fmt.Println("No momento o Robo está Desligado...")
	fmt.Println("[1] - Para LIGAR o robo!")
	fmt.Println("[0] - Para FECHAR o programa!")
}

func toOffMsg() {
	fmt.Println("No momento o Robo está Ligado...")
	fmt.Println("[2] - Para DESLIGAR o robo!")
	fmt.Println("[0] - Para FECHAR o programa!")
}

func runRobotCore(browserName string) {
	fmt.Println("Ligando...")
	for {
		time.Sleep(7 * time.Second)
		if option == 1 {
			robotgo.ActiveName(browserName)
			robotgo.KeyPress("f5")
			clearCLI()
			toOffMsg()
		} else {
			break
		}
	}
}

func main() {
	var browserName = "chrome"

	fmt.Println("Qual navegador será utilizado? [Padrão: chrome]")
	fmt.Println("[1] - Chrome")
	fmt.Println("[2] - Edge")
	fmt.Println("[3] - Mozilla")

	var browserOpt int
	fmt.Scan(&browserOpt)

	clearCLI()

	if browserOpt == 2 {
		browserName = "edge"
	} else if browserOpt == 3 {
		browserName = "firefox"
	}

	for {
		toOnMsg()

		fmt.Scan(&option)

		clearCLI()
		if option == 1 {
			go runRobotCore(browserName)
			clearCLI()
			toOffMsg()
			fmt.Scan(&option)
		}

		clearCLI()
		fmt.Println("Desligando...")

		if option == 0 {
			clearCLI()
			break
		}
	}

	fmt.Println("Robo foi ENCERRADO...")
	fmt.Println("Não esqueça de FECHAR essa JANELA!!")
}
