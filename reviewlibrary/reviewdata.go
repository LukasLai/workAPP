package reviewlibrary

import (
	"math/rand"
)

func ThankBtn() string {
	text := ""
	randomNumber := rand.Intn(2) + 1
	switch randomNumber {
	case 1:
		text = "謝謝您的評論~\n"
	case 2:
		text = "謝謝您的住宿回饋\n"
	}
	return text
}

func NoiceBtn() string {
	text := ""
	randomNumber := rand.Intn(2) + 1
	switch randomNumber {
	case 1:
		text = "遇到聲音上的問題時，我們櫃台旁都有提供免費的耳塞供住客使用喔\n"
	case 2:
		text = "如遇到吵雜的問題，都可以立即聯絡小管家，我們都很樂意為您協助\n"
	}
	return text
}

func BadsmellBtn() string {
	text := "遇到房內有異味，請馬上告知櫃台的服務人員我們都會馬上為您處理~\n"

	return text
}

func MoistureBtn() string {
	text := "建議冷氣可以調低一些, 濕氣會有所改善喔~\n"

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
	randomNumber := rand.Intn(2) + 1
	switch randomNumber {
	case 1:
		text = "(〒︿〒)\n"
	case 2:
		text = ":((\n"
	}
	return text
}

func InviteToPlayBtn() string {
	text := ""
	randomNumber := rand.Intn(2) + 1
	switch randomNumber {
	case 1:
		text = "歡迎下次再來玩喔~\n"
	case 2:
		text = "希望下次還可以來我們這邊玩喔~\n"
	}
	return text
}

func StaffComplimentBtn() string {
	text := ""
	randomNumber := rand.Intn(2) + 1
	switch randomNumber {
	case 1:
		text = "謝謝您對小管家們的稱讚\n"
	case 2:
		text = "小管家們會很高興有您的讚美\n"
	}
	return text
}

func TrafficComplimentBtn() string {
	text := ""
	randomNumber := rand.Intn(2) + 1
	switch randomNumber {
	case 1:
		text = "這邊的交通真的很方便呢，去哪都很近\n"
	case 2:
		text = "這邊不管搭公車還是搭火車都非常便利呢\n"
	}
	return text
}

func EnvironmentComplimentBtn() string {
	text := ""
	randomNumber := rand.Intn(2) + 1
	switch randomNumber {
	case 1:
		text = "交誼廳真的很大很舒適，一不小心就會太放鬆了呢\n"
	case 2:
		text = "交誼廳不管是要閱讀、看影片或是跟旅客們聊天都非常適合唷\n"
	}
	return text
}
