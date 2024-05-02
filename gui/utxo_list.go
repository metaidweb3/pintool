package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func GetUtxoListUi(window *fyne.Window, data binding.ExternalStringList, selectedId binding.Int) *widget.PopUp {
	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("List Item")
		},
		func(i binding.DataItem, obj fyne.CanvasObject) {
			obj.(*widget.Label).Bind(i.(binding.String))
		},
	)
	list.Resize(fyne.NewSize(400, 500))

	selectdItem := -1
	list.OnSelected = func(id widget.ListItemID) {
		selectdItem = id
	}
	var popup *widget.PopUp
	closBtn := widget.NewButton("Close", func() {
		popup.Hide()
	})
	subBtn := widget.NewButton("Confirm", func() {
		//d, _ := data.Get()
		selectedId.Set(selectdItem)
		//fmt.Println("Selected item:", d[selectdItem])
		popup.Hide()
	})
	txt := widget.NewLabel("Select a Utxo")
	txt.TextStyle.Bold = true
	content := container.NewBorder(
		txt,
		container.NewHBox(subBtn, closBtn),
		nil,
		nil,
		list,
	)
	popup = widget.NewModalPopUp(content, (*window).Canvas())

	return popup
}
