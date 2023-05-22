package main

import (
	"fmt"
	"net"
)

var conn *net.TCPConn
var username string

func conecta_servidor() string {
	fmt.Printf("Connectando ao servidor")
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)

	if err != nil {
		return fmt.Sprint("Erro (Mano esse IP existe?):", err.Error())
	}

	conn, err = net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		return fmt.Sprint("Erro na conexão:" + err.Error())

	}

	fmt.Println("Connectado ao servidor")
	return "Connectado ao servidor"
}
func desconecta_servidor() {
	conn.Close()
	fmt.Println("Até mas:" + username)
}
func registra_usuario() { //Sacanagem essa implementação em Pedrão

	received := make([]byte, 256)
	conn.Read(received)

	println(string(received))
	fmt.Scanf("%s", &username)
	conn.Write([]byte(username))

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

func envia_msg(msg string, msgSecret string) {
	//estegranografar aqui a msg
	print(msg)
	conn.Write([]byte(msg))
	fmt.Fprintln(mensage_box, "[#1aab00] VOCE -> "+msg)
	if msgSecret != "" {
		fmt.Fprintln(mensage_box, "[#cf1600] VOCE EM SEGREDO -> "+msgSecret)
	}

}
