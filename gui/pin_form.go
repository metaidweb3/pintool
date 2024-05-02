package gui

import (
	"log"
	"strconv"

	service "self-tool/service/inscription_service"

	"self-tool/api"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetPinForm(selectedUtxo binding.String, window *fyne.Window) *fyne.Container {
	operation := widget.NewSelect([]string{"init", "create", "modify", "revoke"}, func(value string) {
		log.Println("Select set to", value)
	})
	shovel := widget.NewEntry()
	shovel.SetPlaceHolder("For use only during MRC20 minting, input PIN ID.")
	feeInput := widget.NewEntry()
	path := widget.NewEntry()
	pinContent := widget.NewMultiLineEntry()
	pinContent.SetMinRowsVisible(7)
	contentType := widget.NewEntry()
	contentType.SetText("application/jason")
	encryption := widget.NewSelect([]string{"0", "1"}, func(value string) {
	})
	shovelTxt := widget.NewLabel("Shovel(mint):")
	shovelTxt.TextStyle.Bold = true
	feeTxt := widget.NewLabel("FeeRate:")
	feeTxt.TextStyle.Bold = true
	feeInput.SetText("2")
	version := widget.NewEntry()
	version.SetText("1")
	encryption.SetSelectedIndex(0)
	opTxt := widget.NewLabel("Operation:")
	opTxt.TextStyle.Bold = true
	ptTxt := widget.NewLabel("Path:")
	ptTxt.TextStyle.Bold = true
	contentTxt := widget.NewLabel("Content:")
	contentTxt.TextStyle.Bold = true
	contentTypeTxt := widget.NewLabel("ContentType:")
	contentTypeTxt.TextStyle.Bold = true
	encryptionTxt := widget.NewLabel("Encryption:")
	encryptionTxt.TextStyle.Bold = true
	versionTxt := widget.NewLabel("Version:")
	versionTxt.TextStyle.Bold = true

	subBtn := widget.NewButton("Send", func() {
		metaIdData := service.InscriptionMetaIdData{
			MetaIDFlag:  "",
			Operation:   operation.Selected,
			Path:        path.Text,
			Content:     []byte(pinContent.Text),
			Encryption:  encryption.Selected,
			Version:     version.Text,
			ContentType: contentType.Text,
			Destination: "",
		}
		s, _ := selectedUtxo.Get()
		//fmt.Println("send", s, metaIdData)
		f, _ := strconv.ParseInt(feeInput.Text, 10, 64)
		err := api.SendPin(metaIdData, s, "", f)
		if err != nil {
			c := dialog.NewInformation("Error", err.Error(), *window)
			c.Show()
		} else {
			c := dialog.NewInformation("Susccess", "Susccess", *window)
			c.Show()
		}
	})
	btnContent := container.New(layout.NewGridLayout(4), subBtn)

	gridForm := container.New(layout.NewFormLayout(),
		shovelTxt, shovel,
		feeTxt, feeInput,
		opTxt,
		operation,
		ptTxt,
		path,
		contentTxt,
		pinContent,
		contentTypeTxt,
		contentType,
		encryptionTxt,
		encryption,
		versionTxt,
		version,
		widget.NewLabel(""),
		btnContent,
	)

	content := container.NewVBox(gridForm)
	return content
}
