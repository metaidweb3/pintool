package main

import (
	"self-tool/api"
	"self-tool/gui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())
	myWindow := myApp.NewWindow("MAN PIN CREATOR")
	myWindow.Resize(fyne.NewSize(640, 800))
	/*
		t := gui.CreateTop(&myWindow)
		t.Resize(fyne.NewSize(300, 400))
		top := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), t, layout.NewSpacer())
		content := container.New(
			layout.NewVBoxLayout(),
			top,
			gui.GetPinForm(),
		)
	*/
	utxoTxt := widget.NewLabel("UTXO:")
	utxoTxt.TextStyle.Bold = true

	selectedUtxo := binding.NewString()
	fn := binding.NewDataListener(func() {
		s, _ := selectedUtxo.Get()
		utxoTxt.SetText("UTXO: " + api.TxHashformat(s))
		//fmt.Println("1>>>", s)
	})
	selectedUtxo.AddListener(fn)
	top := gui.CreateTop(&myWindow, selectedUtxo)
	top.Resize(fyne.NewSize(600, 0))
	top.Move(fyne.NewPos(20, 20))

	utxoTxt.Resize(fyne.NewSize(600, 30))
	utxoTxt.Move(fyne.NewPos(20, 280))

	pin := gui.GetPinForm(selectedUtxo, &myWindow)
	pin.Resize(fyne.NewSize(600, 0))
	pin.Move(fyne.NewPos(20, 320))

	content := container.NewWithoutLayout(top, utxoTxt, pin)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
