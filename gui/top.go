package gui

import (
	"image/color"
	"log"
	"self-tool/api"
	"self-tool/tool"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateTop(window *fyne.Window, selectedUtxo binding.String) *fyne.Container {
	//title := widget.NewLabel("MAN PIN CREATOR")
	title := canvas.NewText("MAN PIN CREATOR", color.White)
	title.TextStyle.Bold = true
	title.TextSize = 20
	blank := canvas.NewText("", color.White)
	blank.TextSize = 5
	priKeyInput := widget.NewEntry()
	addressInput := widget.NewEntry()
	apiHostInput := widget.NewEntry()
	netCombo := widget.NewSelect([]string{"Btc TestNet", "Btc MainNet", "Btc Regtest"}, func(value string) {
		log.Println("Select set to", value)
	})
	c := dialog.NewInformation("Error", "PriKey and Address and Net are required fields!", *window)
	saveBtn := widget.NewButton("Save", func() {
		if priKeyInput.Text == "" || addressInput.Text == "" || netCombo.Selected == "" {
			c.Show()
		} else {
			data := make(map[string]interface{})
			data["pk"] = priKeyInput.Text
			data["address"] = addressInput.Text
			data["net"] = netCombo.Selected
			data["api"] = apiHostInput.Text
			tool.WriteJSON("./setting.json", data)
		}
	})

	data := binding.BindStringList(&[]string{})
	var utxoList []string
	selectedId := binding.NewInt()
	selectedId.Set(-1)
	fn := binding.NewDataListener(func() {
		s, _ := selectedId.Get()
		if len(utxoList) <= 0 {
			return
		}
		selectedUtxo.Set(utxoList[s])
	})
	selectedId.AddListener(fn)

	popup := GetUtxoListUi(window, data, selectedId)
	popup.Resize(fyne.NewSize(400, 600))

	utxoBtn := widget.NewButton("Select Utxo", func() {
		//data = binding.BindStringList(&[]string{})
		data.Set([]string{})
		utxoList = api.GetUtxoList(addressInput.Text, netCombo.Selected)
		for _, item := range utxoList {
			data.Append(api.TxHashformat(item))
		}
		popup.Show()
	})

	btnContent := container.New(layout.NewGridLayout(4), saveBtn, utxoBtn)
	priKeyTxt := widget.NewLabel("PriKey:")
	priKeyTxt.TextStyle.Bold = true
	addressTxt := widget.NewLabel("Address:")
	addressTxt.TextStyle.Bold = true
	netTxt := widget.NewLabel("Net:")
	netTxt.TextStyle.Bold = true
	apiHostTxt := widget.NewLabel("ApiHost:")
	apiHostTxt.TextStyle.Bold = true

	setting, err := tool.ReadJSON("./setting.json")
	if err == nil {
		priKeyInput.SetText(setting["pk"].(string))
		addressInput.SetText(setting["address"].(string))
		netCombo.SetSelected(setting["net"].(string))
		apiHostInput.SetText(setting["api"].(string))
	}
	gridForm := container.New(layout.NewFormLayout(),
		priKeyTxt,
		priKeyInput,
		addressTxt,
		addressInput,
		netTxt,
		netCombo,
		apiHostTxt,
		apiHostInput,
		widget.NewLabel(""),
		btnContent,
	)
	grid := container.NewVBox(title, blank, gridForm, blank, widget.NewSeparator())

	return grid
}
