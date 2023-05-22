package main

import (
	"fmt"
	"net"
	"strconv"
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

	fmt.Fprintln(mensage_box, "[#1aab00]     VOCE -> "+msg)
	if msgSecret != "" {
		fmt.Fprintln(mensage_box, "[#cf1600]     VOCE EM SEGREDO -> "+msgSecret)
		conn.Write([]byte(msg + encodeString(msgSecret)))
	} else {
		conn.Write([]byte(msg))
	}

}

//Converte caracter para binario
func charToBinary(c byte) string {
	binary := fmt.Sprintf("%08b", c)
	return binary
}

//Esteganografia esta aqui (de rede o servidor não deu boa)
func encodeString(input string) string {
	encoded := ""

	// Converte cada caractere em sua representação binária
	for i := 0; i < len(input); i++ {
		binary := charToBinary(input[i])

		//Codifica cada bit da representação binária usando caracteres de largura zero
		for j := 0; j < len(binary); j++ {
			bit := binary[j]

			// Usa caracteres de largura zero para representar '0' ou '1'
			if bit == '0' {
				encoded += "\u200b" // U+200B ZERO WIDTH SPACE
			} else {
				encoded += "\u200d" // U+200D ZERO WIDTH JOINER
			}
		}
	}

	return encoded
}
func decodeString(encoded string) string {
	decoded := ""
	binary := ""

	// Iterar sobre cada caractere na string codificada
	for i := 0; i < len(encoded); i++ {
		char := encoded[i]

		// Verifique se o caractere é um espaço de largura zero ou um joiner de largura zero
		if string(char) == "\u200b" { // U+200B ZERO WIDTH SPACE
			binary += "0"
		} else if string(char) == "\u200d" { // U+200D ZERO WIDTH JOINER
			binary += "1"
		}

		// Verifique se coletamos 8 bits para formar um caractere completo
		if len(binary) == 8 {
			// Converter a string binária em um byte
			b, _ := strconv.ParseUint(binary, 2, 8)
			decoded += string(b)

			//Workaround: Redefina a string binária para o próximo caractere
			binary = ""
		}
	}

	return decoded
}
