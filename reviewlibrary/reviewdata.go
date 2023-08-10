package reviewlibrary

import (
	"math/rand"
)

func ThankBtn() string {
	text := ""
	randomNumber := rand.Intn(3) + 1
	switch randomNumber {
	case 1:
		text = "謝謝您的評論~\n"
	case 2:
		text = "謝謝您的住宿回饋\n"
	case 3:
		text = "cc\n"
	}
	return text
}

func NoiceBtn() string {
	text := ""
	randomNumber := rand.Intn(3) + 1
	switch randomNumber {
	case 1:
		text = "遇到聲音上的問題時，我們櫃台旁都有提供免費的耳塞供住客使用喔\n"
	case 2:
		text = "bb\n"
	case 3:
		text = "cc\n"
	}
	return text
}

func BadsmellBtn() string {
	text := ""
	randomNumber := rand.Intn(3) + 1
	switch randomNumber {
	case 1:
		text = "如遇到房內有異味，請馬上告知櫃台的服務人員我們都會馬上為您處理~\n"
	case 2:
		text = "bb\n"
	case 3:
		text = "cc\n"
	}
	return text
}

func MoistureBtn() string {
	text := ""
	randomNumber := rand.Intn(3) + 1
	switch randomNumber {
	case 1:
		text = "建議冷氣可以調低一些, 濕氣會有所改善喔~\n"
	case 2:
		text = "bb\n"
	case 3:
		text = "cc\n"
	}
	return text
}

func HappyemojiBtn() string {
	text := ""
	randomNumber := rand.Intn(3) + 1
	switch randomNumber {
	case 1:
		text = "d(`･∀･)b\n"
	case 2:
		text = "(*´∀`)~♥\n"
	case 3:
		text = "(^u^)\n"
	}
	return text
}

func SademojiBtn() string {
	text := ""
	randomNumber := rand.Intn(3) + 1
	switch randomNumber {
	case 1:
		text = "(〒︿〒)\n"
	case 2:
		text = ":((\n"
	}
	return text
}
