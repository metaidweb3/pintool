package man_tutorials

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}

func welcomeScreen(_ fyne.Window) fyne.CanvasObject {
	//logo := canvas.NewImageFromResource(data.FyneLogoTransparent)
	//logo.FillMode = canvas.ImageFillContain
	//if fyne.CurrentDevice().IsMobile() {
	//	logo.SetMinSize(fyne.NewSize(192, 192))
	//} else {
	//	logo.SetMinSize(fyne.NewSize(256, 256))
	//}

	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("Welcome to the MySelf-Tool app", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabel(""), // balance the header on the tutorial screen we leave blank on this content
	))
}
