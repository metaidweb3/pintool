package man_tutorials

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func metaidInscriptionScreen(_ fyne.Window) fyne.CanvasObject {

	formStruct := struct {
		MetaIDFlag  string
		Operation   string
		Path        string
		Content     string
		Encryption  string
		Version     string
		ContentType string
	}{}

	formData := binding.BindStruct(&formStruct)
	form := newFormWithData(formData)
	form.OnSubmit = func() {
		fmt.Println("Struct:\n", formStruct)
	}

	listPanel := container.NewBorder(nil, nil, nil, nil)
	return container.NewBorder(nil, nil, nil, nil, container.NewGridWithColumns(2, listPanel, form))

	return container.NewCenter(container.NewVBox())
}

func newFormWithData(data binding.DataMap) *widget.Form {
	keys := data.Keys()
	items := make([]*widget.FormItem, len(keys))
	for i, k := range keys {
		data, err := data.GetItem(k)
		if err != nil {
			items[i] = widget.NewFormItem(k, widget.NewLabel(err.Error()))
		}
		items[i] = widget.NewFormItem(k, createBoundItem(data))
	}

	return widget.NewForm(items...)
}

func createBoundItem(v binding.DataItem) fyne.CanvasObject {
	switch val := v.(type) {
	case binding.Bool:
		return widget.NewCheckWithData("", val)
	case binding.Float:
		s := widget.NewSliderWithData(0, 1, val)
		s.Step = 0.01
		return s
	case binding.Int:
		return widget.NewEntryWithData(binding.IntToString(val))
	case binding.String:
		return widget.NewEntryWithData(val)
	default:
		return widget.NewLabel("")
	}
}
