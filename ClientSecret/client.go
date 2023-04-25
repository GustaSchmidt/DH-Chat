package main

import (
	"fmt"
	"net"
	"os"
)

const (
	HOST = "127.0.0.1"
	PORT = "8080"
	TYPE = "tcp"
)

var conn *net.TCPConn

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

	registra_usuario() //Só para deixar claro que isso foi sacanagem

	// conn.Close()
}
