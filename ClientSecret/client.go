package main

import (
	"fmt"
	"net"
	"os"

	"github.com/rivo/tview"
)

const (
	HOST = "127.0.0.1"
	PORT = "8080"
	TYPE = "tcp"
)

var conn *net.TCPConn
var app = tview.NewApplication()

func registra_usuario() { //Sacanagem essa implementação em Pedrão
	var usarname string

	received := make([]byte, 256)
	conn.Read(received)

	println(string(received))
	fmt.Scanf("%s", &usarname)
	conn.Write([]byte(usarname))

	confirm := make([]byte, 3)
	conn.Read(confirm)
	// println(string(confirm))

	if string(confirm) != "OK\x01" {
		conn.Write([]byte("exit\x01"))
	}

	online_users := make([]byte, 2)
	conn.Read(online_users)
	println(string("Usuarios conectados: " + string(online_users)))

}
func trata_usuario() {
	for { //n vou trata erro não to nem ai
		received := make([]byte, 2048)
		length, _ := conn.Read(received)
		fmt.Println(string(received[:length]))
		// var msg string
		// fmt.Scanf("%s", &msg)
		// conn.Write([]byte(msg))
	}

}
func GUI() {
	// Initialize application
	user_list := tview.NewBox().SetBorder(true).SetTitle("Usuarios")
	mensage_box := tview.NewBox().SetBorder(true).SetTitle("Chat Secreto")
	ininputs_msg := tview.NewBox().SetBorder(true).SetTitle("Input de mensagens")

	flex := tview.NewFlex().
		AddItem(user_list, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(mensage_box, 0, 3, false).
			AddItem(ininputs_msg, 5, 1, false), 0, 2, false)
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}

}
func main() {
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)

	if err != nil {
		println("Erro (Mano esse IP existe?):", err.Error())
		os.Exit(1)
	}

	conn, err = net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		println("Erro na conexão:", err.Error())
		os.Exit(1)
	}

	if err != nil {
		println("To conseguindo escrever no output não:", err.Error())
		os.Exit(1)
	}

	//Criando uma interface pro negocio
	GUI()

	//registra_usuario() //Só para deixar claro que isso foi sacanagem
	//trata_usuario()
	// conn.Close()
}
