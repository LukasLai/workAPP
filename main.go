package main

import (
	"fmt"
	"image/color"
	"os"
	"strings"
	"workAPP/inventoryparse"
	"workAPP/reviewlibrary"

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

	// 初始化 functionalBtnVbox
	functionalBtnVbox := container.NewVBox(layout.NewSpacer())
	text := ""

	outputData := binding.NewString()
	outputLabel := widget.NewLabelWithData(outputData)

	//評論區按鈕
	reviewBtn := widget.NewButton("評論", func() {
		// 清空 functionalBtnVbox 中的內容
		functionalBtnVbox.Objects = []fyne.CanvasObject{}
		// 清空text內容
		text = ""
		outputData.Set(text)

		// 生成新的按鈕並添加到 functionalBtnVbox 中
		button1 := widget.NewButton("感謝", func() {
			text += reviewlibrary.ThankBtn()
			outputData.Set(text)
		})

		button2 := widget.NewButton("聲音", func() {
			text += reviewlibrary.NoiceBtn()
			outputData.Set(text)
		})

		button3 := widget.NewButton("異味", func() {
			text += reviewlibrary.BadsmellBtn()
			outputData.Set(text)
		})

		button4 := widget.NewButton("濕氣/霉味", func() {
			text += reviewlibrary.MoistureBtn()
			outputData.Set(text)
		})

		button5 := widget.NewButton("開心Emoji", func() {
			text += reviewlibrary.HappyemojiBtn()
			outputData.Set(text)
		})

		button6 := widget.NewButton("傷心Emoji", func() {
			text += reviewlibrary.SademojiBtn()
			outputData.Set(text)
		})

		button7 := widget.NewButton("邀請來玩", func() {
			text += reviewlibrary.InviteToPlayBtn()
			outputData.Set(text)
		})

		button8 := widget.NewButton("人員稱讚", func() {
			text += reviewlibrary.StaffComplimentBtn()
			outputData.Set(text)
		})

		button9 := widget.NewButton("交通稱讚", func() {
			text += reviewlibrary.TrafficComplimentBtn()
			outputData.Set(text)
		})

		button10 := widget.NewButton("環境稱讚", func() {
			text += reviewlibrary.EnvironmentComplimentBtn()
			outputData.Set(text)
		})

		clearBtn := widget.NewButton("重製", func() {
			text = ""
			outputData.Set(text)
		})

		copyBtn := container.NewHBox(
			layout.NewSpacer(),
			widget.NewButtonWithIcon("Copy複製", theme.ContentCopyIcon(), func() {
				if content, err := outputData.Get(); err == nil {
					w.Clipboard().SetContent(content)
				}
			}),
		)

		functionalBtnVbox.Add(button1)
		functionalBtnVbox.Add(button2)
		functionalBtnVbox.Add(button3)
		functionalBtnVbox.Add(button4)
		functionalBtnVbox.Add(button5)
		functionalBtnVbox.Add(button6)
		functionalBtnVbox.Add(button7)
		functionalBtnVbox.Add(button8)
		functionalBtnVbox.Add(button9)
		functionalBtnVbox.Add(button10)
		functionalBtnVbox.Add(layout.NewSpacer())
		functionalBtnVbox.Add(clearBtn)
		functionalBtnVbox.Add(copyBtn)

		// 刷新 functionalBtnVbox 使新的按鈕生效
		functionalBtnVbox.Refresh()
	})

	inventoryBtn := widget.NewButton("點貨", func() {
		functionalBtnVbox.Objects = []fyne.CanvasObject{}
		invenWindow(a, w)
		text = ""
		outputData.Set(text)
	})

	//日常回覆按鈕
	replyBtn := widget.NewButton("日常回覆", func() {
		// 清空 functionalBtnVbox 中的內容
		functionalBtnVbox.Objects = []fyne.CanvasObject{}
		// 清空text內容
		text = ""
		outputData.Set(text)

		// 生成新的按鈕並添加到 functionalBtnVbox 中
		checkinTimeBtn := widget.NewButton("詢問入住(平)", func() {
			text = reviewlibrary.CheckinTimebtn()
			outputData.Set(text)
		})

		checkinTimeBtn2 := widget.NewButton("詢問入住(假)", func() {
			text = reviewlibrary.CheckinTimebtn2()
			outputData.Set(text)
		})

		uploadidBtn := widget.NewButton("上傳證件", func() {
			text = reviewlibrary.Uploadidbtn()
			outputData.Set(text)
		})

		breakfastBtn := widget.NewButton("早餐", func() {
			text = reviewlibrary.Breakfastbtn()
			outputData.Set(text)
		})

		clearBtn := widget.NewButton("重製", func() {
			text = ""
			outputData.Set(text)
		})

		copyBtn := container.NewHBox(
			layout.NewSpacer(),
			widget.NewButtonWithIcon("Copy複製", theme.ContentCopyIcon(), func() {
				if content, err := outputData.Get(); err == nil {
					w.Clipboard().SetContent(content)
				}
			}),
		)

		functionalBtnVbox.Add(checkinTimeBtn)
		functionalBtnVbox.Add(checkinTimeBtn2)
		functionalBtnVbox.Add(uploadidBtn)
		functionalBtnVbox.Add(breakfastBtn)
		functionalBtnVbox.Add(layout.NewSpacer())
		functionalBtnVbox.Add(clearBtn)
		functionalBtnVbox.Add(copyBtn)

		// 刷新 functionalBtnVbox 使新的按鈕生效
		functionalBtnVbox.Refresh()
	})

	other2 := canvas.NewText("更新中", color.White)
	//最左邊的垂直欄位
	sidebarVbox := container.NewGridWithRows(4, reviewBtn, inventoryBtn, replyBtn, other2)
	//整體UI的水平欄位
	outputVbox := container.NewVBox(layout.NewSpacer(), outputLabel, layout.NewSpacer(), layout.NewSpacer())

	content := container.NewHBox(
		sidebarVbox,
		functionalBtnVbox,
		outputVbox,
	)

	w.SetContent(content)
	w.ShowAndRun()
}

// 點貨視窗
func invenWindow(a fyne.App, w fyne.Window) {
	newWindow := a.NewWindow("點貨視窗")
	newWindow.Resize(fyne.NewSize(700, 500))

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

		txtBoundV1.Set(strings.Join(v1, " "))
		txtBoundV2.Set(strings.Join(v2, " "))
		txtBoundV3.Set(strings.Join(v3, " "))
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
