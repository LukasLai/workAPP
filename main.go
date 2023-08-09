package main

import (
	"fmt"
	"image/color"
	"os"
	"strings"
	"workAPP/inventoryparse"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	os.Setenv("FYNE_FONT", "TaipeiSansTCBeta-Regular.ttf")
	a := app.New()
	w := a.NewWindow("懶人APP")
	w.Resize(fyne.NewSize(700, 500))
	a.Settings().SetTheme(theme.DarkTheme())

	//output := widget.NewLabel("")
	//output.Wrapping = fyne.TextWrapWord
	//output.Alignment = fyne.TextAlignLeading

	// 初始化 reviewBtnVbox
	reviewBtnVbox := container.NewVBox(layout.NewSpacer())
	text := ""

	reviewData := binding.NewString()
	reviewLabel := widget.NewLabelWithData(reviewData)

	reviewBtn := widget.NewButton("評論", func() {
		// 清空 reviewBtnVbox 中的內容
		reviewBtnVbox.Objects = []fyne.CanvasObject{}

		// 生成新的按鈕並添加到 reviewBtnVbox 中
		button1 := widget.NewButton("感謝", func() {
			text += "謝謝您的評論~\n"
			reviewData.Set(text)

		})
		button2 := widget.NewButton("聲音", func() {
			text += "遇到聲音上的問題時，我們櫃台旁都有提供免費的耳塞供住客使用喔\n"
			reviewData.Set(text)
		})
		button3 := widget.NewButton("異味", func() {
			text += "如遇到房內有異味，請馬上告知櫃台的服務人員我們都會馬上為您處理~\n"
			reviewData.Set(text)
		})

		button4 := widget.NewButton("濕氣/霉味", func() {
			text += "建議冷氣可以調低一些, 濕氣會有所改善喔~\n"
			reviewData.Set(text)
		})

		clearBtn := widget.NewButton("重製", func() {
			text = ""
			reviewData.Set(text)
		})

		copyBtn := container.NewHBox(
			layout.NewSpacer(),
			widget.NewButtonWithIcon("Copy複製", theme.ContentCopyIcon(), func() {
				if content, err := reviewData.Get(); err == nil {
					w.Clipboard().SetContent(content)
				}
			}),
		)

		reviewBtnVbox.Add(button1)
		reviewBtnVbox.Add(button2)
		reviewBtnVbox.Add(button3)
		reviewBtnVbox.Add(button4)
		reviewBtnVbox.Add(layout.NewSpacer())
		reviewBtnVbox.Add(clearBtn)
		reviewBtnVbox.Add(copyBtn)

		// 刷新 reviewBtnVbox 使新的按鈕生效
		reviewBtnVbox.Refresh()
	})

	inventoryBtn := widget.NewButton("點貨", func() {
		reviewBtnVbox.Objects = []fyne.CanvasObject{}
		invenWindow(a, w)
		text = ""
		reviewData.Set(text)
	})
	other := canvas.NewText("更新中", color.White)
	other2 := canvas.NewText("更新中", color.White)
	sidebarVbox := container.NewGridWithRows(4, reviewBtn, inventoryBtn, other, other2)
	outputVbox := container.NewVBox(layout.NewSpacer(), reviewLabel, layout.NewSpacer(), layout.NewSpacer())
	content := container.NewHBox(
		sidebarVbox,
		reviewBtnVbox,
		outputVbox,
	)

	w.SetContent(content)
	w.ShowAndRun()
}

// 點貨視窗
func invenWindow(a fyne.App, w fyne.Window) {
	newWindow := a.NewWindow("點貨視窗")
	newWindow.Resize(fyne.NewSize(700, 500))

	/*topicText := widget.NewLabel("點貨")
	closeBtn := widget.NewButton("關閉", func() { newWindow.Close() })

	content := container.NewBorder(topicText, closeBtn, nil, nil, nil)
	*/
	//------------------------------------------------------------------------
	txtBoundV1 := binding.NewString()
	txtBoundV2 := binding.NewString()
	txtBoundV3 := binding.NewString()

	txtWidV1 := widget.NewLabelWithData(txtBoundV1)
	txtWidV2 := widget.NewLabelWithData(txtBoundV2)
	txtWidV3 := widget.NewLabelWithData(txtBoundV3)

	copyBtnV1 := container.NewHBox(
		layout.NewSpacer(),
		widget.NewButtonWithIcon("copy content", theme.ContentCopyIcon(), func() {
			if content, err := txtBoundV1.Get(); err == nil {
				newWindow.Clipboard().SetContent(content)
			}
		}),
	)

	copyBtnV2 := container.NewHBox(
		layout.NewSpacer(),
		widget.NewButtonWithIcon("copy content", theme.ContentCopyIcon(), func() {
			if content, err := txtBoundV2.Get(); err == nil {
				newWindow.Clipboard().SetContent(content)
			}
		}),
	)

	copyBtnV3 := container.NewHBox(
		layout.NewSpacer(),
		widget.NewButtonWithIcon("copy content", theme.ContentCopyIcon(), func() {
			if content, err := txtBoundV3.Get(); err == nil {
				newWindow.Clipboard().SetContent(content)
			}
		}),
	)
	entryMultiLine := widget.NewMultiLineEntry()
	entryMultiLine.SetPlaceHolder("請輸入內容...")

	resultText := canvas.NewText("叫貨清單:", color.Black)
	resultText.TextSize = 15

	sentText := canvas.NewText("", color.Black)
	sentBtn := widget.NewButton("Sent", func() {
		sentText.Text = "送出成功"
		fmt.Println("sent按鈕觸發")
		v1, v2, v3 := inventoryparse.ParseInvenroty(entryMultiLine.Text)

		txtBoundV1.Set(strings.Join(v1, ","))
		txtBoundV2.Set(strings.Join(v2, ","))
		txtBoundV3.Set(strings.Join(v3, ","))
	})
	vendor1 := canvas.NewText("升威:", color.White)
	vendor2 := canvas.NewText("安美潔:", color.White)
	vendor3 := canvas.NewText("遠隆:", color.White)

	v1Hbox := container.NewHBox(
		vendor1,
		txtWidV1,
		layout.NewSpacer(),
		copyBtnV1,
	)
	v2Hbox := container.NewHBox(
		vendor2,
		txtWidV2,
		layout.NewSpacer(),
		copyBtnV2,
	)

	v3Hbox := container.NewHBox(
		vendor3,
		txtWidV3,
		layout.NewSpacer(),
		copyBtnV3,
	)
	closeBtn := widget.NewButton("關閉", func() { newWindow.Close() })

	vbox := container.NewVBox(
		entryMultiLine,
		sentText,
		sentBtn,
		v1Hbox,
		v2Hbox,
		v3Hbox,
		layout.NewSpacer(),
		closeBtn,
	)
	newWindow.SetContent(vbox)
	newWindow.Show()
}
