package main

import (
	"fmt"

	"github.com/rivo/tview"
)

var app = tview.NewApplication()
var user_list = tview.NewList().ShowSecondaryText(false)
var mensage_box = tview.NewTextView().
	SetDynamicColors(true).
	SetRegions(true).
	SetChangedFunc(func() {
		app.Draw()
	})

func chat_main() {
	//inputs de msg
	var input_msg = tview.NewInputField().
		SetLabel("Mensagem O: ")
	var input_secret = tview.NewInputField().
		SetLabel("Mensagem S: ")
	var send_bt = tview.NewButton("Enviar").
		SetSelectedFunc(func() {
			//Enviar msg
			envia_msg(input_msg.GetText(), input_secret.GetText())
			input_msg.SetText("")
			input_secret.SetText("")
			app.SetFocus(input_msg)
			app.Sync()
		})

	//TODO: receber mensagem de forma descente
	fmt.Fprintf(mensage_box, "%s ", "word")

	//Adicona o propio usario a lista de usuarios
	user_list.AddItem(username, "O propio", '-', nil)
	user_list.SetBorder(true)

	//layout do chat
	flex := tview.NewFlex().
		AddItem(user_list, 0, 25, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(mensage_box, 0, 3, false).
			AddItem(input_msg, 2, 1, false).
			AddItem(input_secret, 2, 1, false).
			AddItem(send_bt, 1, 1, false), 0, 75, false)

	//INIT INTERFACE
	app.SetRoot(flex, true).EnableMouse(true).SetFocus(input_msg).Run()
}