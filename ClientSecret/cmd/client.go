package main

const ( // vai fica chumbado? talvez depende da preguiça
	HOST = "127.0.0.1"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	conecta_servidor()
	registra_usuario()
	chat_main()

}
